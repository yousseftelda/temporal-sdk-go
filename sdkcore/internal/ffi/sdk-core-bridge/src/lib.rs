#![allow(non_camel_case_types)]

extern crate libc;

use prost::Message;

#[no_mangle]
pub extern "C" fn hello_rust() {
    println!("Hello, Rust!");
}

pub struct tmprl_error_t {
    message: String,
    // TODO(cretz): Enums?
    code: i32,
}

impl tmprl_error_t {
    fn new(message: String, code: i32) -> tmprl_error_t {
        tmprl_error_t { message, code }
    }
}

#[no_mangle]
pub extern "C" fn tmprl_error_free(error: *mut tmprl_error_t) {
    if !error.is_null() {
        unsafe {
            Box::from_raw(error);
        }
    }
}

/// Not null terminated (use tmprl_error_message_len()) and remains owned by
/// error
#[no_mangle]
pub extern "C" fn tmprl_error_message(error: *const tmprl_error_t) -> *const libc::c_char {
    if error.is_null() {
        return std::ptr::null_mut();
    }
    unsafe { (*error).message.as_ptr() as *const libc::c_char }
}

#[no_mangle]
pub extern "C" fn tmprl_error_message_len(error: *const tmprl_error_t) -> libc::size_t {
    if error.is_null() {
        return 0;
    }
    unsafe { (*error).message.len() as libc::size_t }
}

// This is zero unless the error is known to have a code
#[no_mangle]
pub extern "C" fn tmprl_error_message_code(error: *const tmprl_error_t) -> libc::c_int {
    if error.is_null() {
        return -1;
    }
    unsafe { (*error).code as libc::c_int }
}

pub struct tmprl_runtime_t {
    tokio_runtime: std::sync::Arc<tokio::runtime::Runtime>,
}

#[no_mangle]
pub extern "C" fn tmprl_runtime_new() -> *mut tmprl_runtime_t {
    Box::into_raw(Box::new(tmprl_runtime_t {
        // TODO(cretz): Options to configure thread pool?
        tokio_runtime: std::sync::Arc::new(
            tokio::runtime::Builder::new_multi_thread()
                .enable_all()
                .build()
                .unwrap(),
        ),
    }))
}

#[no_mangle]
pub extern "C" fn tmprl_runtime_free(runtime: *mut tmprl_runtime_t) {
    if !runtime.is_null() {
        unsafe {
            Box::from_raw(runtime);
        }
    }
}

pub struct tmprl_core_t {
    tokio_runtime: std::sync::Arc<tokio::runtime::Runtime>,
    // TODO(cretz): Any concerns with the dynamic dispatch overhead here?
    // TODO(cretz): Should we ask Rust SDK to give us back a struct instead?
    // TODO(cretz): We could use generics, but impl traits don't cross the FFI
    // boundary properly
    core: std::sync::Arc<dyn temporal_sdk_core::Core>,
}

impl tmprl_core_t {
    fn borrow_buf(&mut self) -> Vec<u8> {
        // TODO(cretz): Implement real thread-safe pool?
        Vec::new()
    }

    fn return_buf(&mut self, vec: Vec<u8>) {
        // TODO(cretz): Implement real thread-safe pool?
    }

    fn encode_proto(&mut self, proto: &impl prost::Message) -> tmprl_bytes_or_error_t {
        let mut buf = self.borrow_buf();
        buf.clear();
        // Increase buf capacity if needed
        buf.reserve(proto.encoded_len());
        match proto.encode(&mut buf) {
            Ok(_) => tmprl_bytes_or_error_t {
                bytes: tmprl_bytes::from_vec(buf),
                error: std::ptr::null_mut(),
            },
            Err(err) => {
                self.return_buf(buf);
                tmprl_bytes_or_error_t {
                    bytes: std::ptr::null_mut(),
                    error: Box::into_raw(Box::new(tmprl_error_t::new(
                        format!("failed encoding proto: {}", err),
                        0,
                    ))),
                }
            }
        }
    }
}

/// One and only one field is non-null, caller must free whichever it is
#[repr(C)]
pub struct tmprl_core_or_error_t {
    core: *mut tmprl_core_t,
    error: *mut tmprl_error_t,
}

#[repr(C)]
pub struct tmprl_core_new_options_t {
    /// UTF-8, not null terminated, not owned
    target_url: *const libc::c_char,
    target_url_len: libc::size_t,
    // TODO(cretz): Other stuff
}

fn utf8_string_ref<'a>(data: *const libc::c_char, len: libc::size_t) -> &'a str {
    if data.is_null() || len == 0 {
        return "";
    }
    unsafe {
        let bytes = std::slice::from_raw_parts(data as *const u8, len);
        std::str::from_utf8_unchecked(bytes)
    }
}

#[no_mangle]
pub extern "C" fn tmprl_core_new(
    runtime: *mut tmprl_runtime_t,
    options: *const tmprl_core_new_options_t,
) -> tmprl_core_or_error_t {
    match core_new(runtime, options) {
        Ok(core) => tmprl_core_or_error_t {
            core: Box::into_raw(Box::new(core)),
            error: std::ptr::null_mut(),
        },
        Err(err) => tmprl_core_or_error_t {
            core: std::ptr::null_mut(),
            error: Box::into_raw(Box::new(err)),
        },
    }
}

fn core_new(
    runtime: *mut tmprl_runtime_t,
    options: *const tmprl_core_new_options_t,
) -> Result<tmprl_core_t, tmprl_error_t> {
    if runtime.is_null() || options.is_null() {
        return Err(tmprl_error_t::new(
            "missing runtime or options".to_string(),
            0,
        ));
    }
    let runtime = unsafe { &*runtime };
    let options = unsafe { &*options };
    // Build options
    // TODO(cretz): Rest of the options
    let gateway_opts = temporal_sdk_core::ServerGatewayOptionsBuilder::default()
        .target_url(
            temporal_sdk_core::Url::parse(utf8_string_ref(
                options.target_url,
                options.target_url_len,
            ))
            .map_err(|err| tmprl_error_t::new(format!("invalid URL: {}", err), 0))?,
        )
        .namespace("mynamespace".to_string())
        .build()
        .map_err(|err| tmprl_error_t::new(format!("invalid gateway options: {}", err), 0))?;
    let telemetry_opts = temporal_sdk_core::TelemetryOptionsBuilder::default()
        .build()
        .map_err(|err| tmprl_error_t::new(format!("invalid telemetry options: {}", err), 0))?;
    let core_opts = temporal_sdk_core::CoreInitOptionsBuilder::default()
        .gateway_opts(gateway_opts)
        .telemetry_opts(telemetry_opts)
        .build()
        .map_err(|err| tmprl_error_t::new(format!("invalid core options: {}", err), 0))?;
    // Create core and return it
    Ok(tmprl_core_t {
        tokio_runtime: runtime.tokio_runtime.clone(),
        core: std::sync::Arc::new(
            runtime
                .tokio_runtime
                .block_on(temporal_sdk_core::init(core_opts))
                .map_err(|err| tmprl_error_t::new(format!("failed creating core: {}", err), 0))?,
        ),
    })
}

#[no_mangle]
pub extern "C" fn tmprl_core_free(core: *mut tmprl_core_t) {
    if !core.is_null() {
        unsafe {
            Box::from_raw(core);
        }
    }
}

#[repr(C)]
pub struct tmprl_bytes {
    bytes: *const u8,
    len: libc::size_t,
    cap: libc::size_t,
}

impl tmprl_bytes {
    fn from_vec(vec: Vec<u8>) -> *mut tmprl_bytes {
        // Mimics Vec::into_raw_parts that's only available in nightly
        let mut vec = std::mem::ManuallyDrop::new(vec);
        return Box::into_raw(Box::new(tmprl_bytes {
            bytes: vec.as_mut_ptr(),
            len: vec.len(),
            cap: vec.capacity(),
        }));
    }
}

/// One and only one field is non-null, caller must free whichever it is
#[repr(C)]
pub struct tmprl_bytes_or_error_t {
    bytes: *const tmprl_bytes,
    error: *mut tmprl_error_t,
}

#[no_mangle]
pub extern "C" fn tmprl_bytes_free(core: *mut tmprl_core_t, bytes: *const tmprl_bytes) {
    let core = unsafe { &mut *core };
    // Return vec back to core before dropping bytes
    let bytes = bytes as *mut tmprl_bytes;
    let vec = unsafe { Vec::from_raw_parts((*bytes).bytes as *mut u8, (*bytes).len, (*bytes).cap) };
    core.return_buf(vec);
    if !bytes.is_null() {
        unsafe {
            Box::from_raw(bytes);
        }
    }
}

#[repr(C)]
pub struct tmprl_worker_config_t {
    /// UTF-8, not null terminated, not owned
    task_queue: *const libc::c_char,
    task_queue_len: libc::size_t,
    // TODO(cretz): Other stuff
}

#[no_mangle]
pub extern "C" fn tmprl_register_worker(
    core: *mut tmprl_core_t,
    config: *const tmprl_worker_config_t,
) -> *mut tmprl_error_t {
    match register_worker(core, config) {
        Ok(_) => std::ptr::null_mut(),
        Err(err) => Box::into_raw(Box::new(err)),
    }
}

fn register_worker(
    core: *mut tmprl_core_t,
    config: *const tmprl_worker_config_t,
) -> Result<(), tmprl_error_t> {
    if core.is_null() {
        return Err(tmprl_error_t::new("missing core".to_string(), 0));
    }
    let core = unsafe { &*core };
    let config = unsafe { &*config };
    // Build config
    // TODO(cretz): More config
    let worker_config = temporal_sdk_core::WorkerConfigBuilder::default()
        .task_queue(utf8_string_ref(config.task_queue, config.task_queue_len))
        .build()
        .map_err(|err| tmprl_error_t::new(format!("failed creating worker config: {}", err), 0))?;
    // Block on call to register worker
    core.tokio_runtime
        .block_on(core.core.register_worker(worker_config))
        .map_err(|err| tmprl_error_t::new(format!("failed registering worker: {}", err), 0))
}

pub struct tmprl_wf_activation_t {
    activation: temporal_sdk_core_protos::coresdk::workflow_activation::WfActivation,
}

#[no_mangle]
pub extern "C" fn tmprl_wf_activation_free(activation: *mut tmprl_wf_activation_t) {
    if !activation.is_null() {
        unsafe {
            Box::from_raw(activation);
        }
    }
}

#[no_mangle]
pub extern "C" fn tmprl_wf_activation_to_proto(
    core: *mut tmprl_core_t,
    activation: *const tmprl_wf_activation_t,
) -> tmprl_bytes_or_error_t {
    let core = unsafe { &mut *core };
    let activation = unsafe { &*activation };
    core.encode_proto(&activation.activation)
}

/// One and only one field is non-null, caller must free whichever it is
#[repr(C)]
pub struct tmprl_wf_activation_or_error_t {
    activation: *mut tmprl_wf_activation_t,
    error: *mut tmprl_error_t,
}

/// String is UTF-8, not null terminated, and caller still owns
#[no_mangle]
pub extern "C" fn tmprl_poll_workflow_activation(
    core: *mut tmprl_core_t,
    task_queue: *const libc::c_char,
    task_queue_len: libc::size_t,
) -> tmprl_wf_activation_or_error_t {
    match poll_workflow_activation(core, task_queue, task_queue_len) {
        Ok(v) => tmprl_wf_activation_or_error_t {
            activation: Box::into_raw(Box::new(v)),
            error: std::ptr::null_mut(),
        },
        Err(err) => tmprl_wf_activation_or_error_t {
            activation: std::ptr::null_mut(),
            error: Box::into_raw(Box::new(err)),
        },
    }
}

fn poll_workflow_activation(
    core: *mut tmprl_core_t,
    task_queue: *const libc::c_char,
    task_queue_len: libc::size_t,
) -> Result<tmprl_wf_activation_t, tmprl_error_t> {
    let core = unsafe { &*core };
    match core.tokio_runtime.block_on(
        core.core
            .poll_workflow_activation(utf8_string_ref(task_queue, task_queue_len)),
    ) {
        Ok(v) => Ok(tmprl_wf_activation_t { activation: v }),
        Err(err) => Err(tmprl_error_t::new(
            format!("failed polling workflow activation: {}", err),
            0,
        )),
    }
}

pub struct tmprl_activity_task_t {
    task: temporal_sdk_core_protos::coresdk::activity_task::ActivityTask,
}

#[no_mangle]
pub extern "C" fn tmprl_activity_task_free(task: *mut tmprl_activity_task_t) {
    if !task.is_null() {
        unsafe {
            Box::from_raw(task);
        }
    }
}

#[no_mangle]
pub extern "C" fn tmprl_activity_task_to_proto(
    core: *mut tmprl_core_t,
    task: *const tmprl_activity_task_t,
) -> tmprl_bytes_or_error_t {
    let core = unsafe { &mut *core };
    let task = unsafe { &*task };
    core.encode_proto(&task.task)
}

/// One and only one field is non-null, caller must free whichever it is
#[repr(C)]
pub struct tmprl_activity_task_or_error_t {
    task: *mut tmprl_activity_task_t,
    error: *mut tmprl_error_t,
}

/// String is UTF-8, not null terminated, and caller still owns
#[no_mangle]
pub extern "C" fn tmprl_poll_activity_task(
    core: *mut tmprl_core_t,
    task_queue: *const libc::c_char,
    task_queue_len: libc::size_t,
) -> tmprl_activity_task_or_error_t {
    match poll_activity_task(core, task_queue, task_queue_len) {
        Ok(v) => tmprl_activity_task_or_error_t {
            task: Box::into_raw(Box::new(v)),
            error: std::ptr::null_mut(),
        },
        Err(err) => tmprl_activity_task_or_error_t {
            task: std::ptr::null_mut(),
            error: Box::into_raw(Box::new(err)),
        },
    }
}

fn poll_activity_task(
    core: *mut tmprl_core_t,
    task_queue: *const libc::c_char,
    task_queue_len: libc::size_t,
) -> Result<tmprl_activity_task_t, tmprl_error_t> {
    let core = unsafe { &*core };
    match core.tokio_runtime.block_on(
        core.core
            .poll_activity_task(utf8_string_ref(task_queue, task_queue_len)),
    ) {
        Ok(v) => Ok(tmprl_activity_task_t { task: v }),
        Err(err) => Err(tmprl_error_t::new(
            format!("failed polling activity task: {}", err),
            0,
        )),
    }
}

pub struct tmprl_wf_activation_completion_t {
    completion: temporal_sdk_core_protos::coresdk::workflow_completion::WfActivationCompletion,
}

/// One and only one field is non-null, caller must free error if it's error or
/// pass completion to tmprl_complete_workflow_activation
#[repr(C)]
pub struct tmprl_wf_activation_completion_or_error_t {
    completion: *mut tmprl_wf_activation_completion_t,
    error: *mut tmprl_error_t,
}

#[no_mangle]
pub extern "C" fn tmprl_wf_activation_completion_from_proto(
    bytes: *const u8,
    bytes_len: libc::size_t,
) -> tmprl_wf_activation_completion_or_error_t {
    let bytes = unsafe { std::slice::from_raw_parts(bytes, bytes_len) };
    match temporal_sdk_core_protos::coresdk::workflow_completion::WfActivationCompletion::decode(
        bytes,
    ) {
        Ok(v) => tmprl_wf_activation_completion_or_error_t {
            completion: Box::into_raw(Box::new(tmprl_wf_activation_completion_t { completion: v })),
            error: std::ptr::null_mut(),
        },
        Err(err) => tmprl_wf_activation_completion_or_error_t {
            completion: std::ptr::null_mut(),
            error: Box::into_raw(Box::new(tmprl_error_t::new(
                format!("failed decoding proto: {}", err),
                0,
            ))),
        },
    }
}

/// Takes ownership of the completion and frees internally
#[no_mangle]
pub extern "C" fn tmprl_complete_workflow_activation(
    core: *mut tmprl_core_t,
    completion: *mut tmprl_wf_activation_completion_t,
) -> *mut tmprl_error_t {
    match complete_workflow_activation(core, completion) {
        Ok(_) => std::ptr::null_mut(),
        Err(err) => Box::into_raw(Box::new(err)),
    }
}

fn complete_workflow_activation(
    core: *mut tmprl_core_t,
    completion: *mut tmprl_wf_activation_completion_t,
) -> Result<(), tmprl_error_t> {
    let core = unsafe { &*core };
    let completion = unsafe { Box::from_raw(completion) };
    core.tokio_runtime
        .block_on(
            core.core
                .complete_workflow_activation(completion.completion),
        )
        .map_err(|err| {
            tmprl_error_t::new(format!("failed completing workflow activation: {}", err), 0)
        })
}

pub struct tmprl_activity_task_completion_t {
    completion: temporal_sdk_core_protos::coresdk::ActivityTaskCompletion,
}

/// One and only one field is non-null, caller must free error if it's error or
/// pass completion to tmprl_complete_activity_task
#[repr(C)]
pub struct tmprl_activity_task_completion_or_error_t {
    completion: *mut tmprl_activity_task_completion_t,
    error: *mut tmprl_error_t,
}

#[no_mangle]
pub extern "C" fn tmprl_activity_task_completion_from_proto(
    bytes: *const u8,
    bytes_len: libc::size_t,
) -> tmprl_activity_task_completion_or_error_t {
    let bytes = unsafe { std::slice::from_raw_parts(bytes, bytes_len) };
    match temporal_sdk_core_protos::coresdk::ActivityTaskCompletion::decode(bytes) {
        Ok(v) => tmprl_activity_task_completion_or_error_t {
            completion: Box::into_raw(Box::new(tmprl_activity_task_completion_t { completion: v })),
            error: std::ptr::null_mut(),
        },
        Err(err) => tmprl_activity_task_completion_or_error_t {
            completion: std::ptr::null_mut(),
            error: Box::into_raw(Box::new(tmprl_error_t::new(
                format!("failed decoding proto: {}", err),
                0,
            ))),
        },
    }
}

/// Takes ownership of the completion and frees internally
#[no_mangle]
pub extern "C" fn tmprl_complete_activity_task(
    core: *mut tmprl_core_t,
    completion: *mut tmprl_activity_task_completion_t,
) -> *mut tmprl_error_t {
    match complete_activity_task(core, completion) {
        Ok(_) => std::ptr::null_mut(),
        Err(err) => Box::into_raw(Box::new(err)),
    }
}

fn complete_activity_task(
    core: *mut tmprl_core_t,
    completion: *mut tmprl_activity_task_completion_t,
) -> Result<(), tmprl_error_t> {
    let core = unsafe { &*core };
    let completion = unsafe { Box::from_raw(completion) };
    core.tokio_runtime
        .block_on(core.core.complete_activity_task(completion.completion))
        .map_err(|err| tmprl_error_t::new(format!("failed completing activity task: {}", err), 0))
}

pub struct tmprl_activity_heartbeat_t {
    heartbeat: temporal_sdk_core_protos::coresdk::ActivityHeartbeat,
}

/// One and only one field is non-null, caller must free error if it's error or
/// pass completion to tmprl_record_activity_heartbeat
#[repr(C)]
pub struct tmprl_activity_heartbeat_or_error_t {
    heartbeat: *mut tmprl_activity_heartbeat_t,
    error: *mut tmprl_error_t,
}

#[no_mangle]
pub extern "C" fn tmprl_activity_heartbeat_from_proto(
    bytes: *const u8,
    bytes_len: libc::size_t,
) -> tmprl_activity_heartbeat_or_error_t {
    let bytes = unsafe { std::slice::from_raw_parts(bytes, bytes_len) };
    match temporal_sdk_core_protos::coresdk::ActivityHeartbeat::decode(bytes) {
        Ok(v) => tmprl_activity_heartbeat_or_error_t {
            heartbeat: Box::into_raw(Box::new(tmprl_activity_heartbeat_t { heartbeat: v })),
            error: std::ptr::null_mut(),
        },
        Err(err) => tmprl_activity_heartbeat_or_error_t {
            heartbeat: std::ptr::null_mut(),
            error: Box::into_raw(Box::new(tmprl_error_t::new(
                format!("failed decoding proto: {}", err),
                0,
            ))),
        },
    }
}

/// Takes ownership of the heartbeat and frees internally
#[no_mangle]
pub extern "C" fn tmprl_record_activity_heartbeat(
    core: *mut tmprl_core_t,
    heartbeat: *mut tmprl_activity_heartbeat_t,
) {
    let core = unsafe { &*core };
    let heartbeat = unsafe { Box::from_raw(heartbeat) };
    core.core.record_activity_heartbeat(heartbeat.heartbeat)
}

/// Strings are UTF-8, not null terminated, and caller still owns
#[no_mangle]
pub extern "C" fn tmprl_request_workflow_eviction(
    core: *mut tmprl_core_t,
    task_queue: *const libc::c_char,
    task_queue_len: libc::size_t,
    run_id: *const libc::c_char,
    run_id_len: libc::size_t,
) {
    let core = unsafe { &*core };
    core.core.request_workflow_eviction(
        utf8_string_ref(task_queue, task_queue_len),
        utf8_string_ref(run_id, run_id_len),
    )
}

#[no_mangle]
pub extern "C" fn tmprl_shutdown(core: *mut tmprl_core_t) {
    let core = unsafe { &*core };
    core.tokio_runtime.block_on(core.core.shutdown())
}

/// String is UTF-8, not null terminated, and caller still owns
#[no_mangle]
pub extern "C" fn tmprl_shutdown_worker(
    core: *mut tmprl_core_t,
    task_queue: *const libc::c_char,
    task_queue_len: libc::size_t,
) {
    let core = unsafe { &*core };
    core.tokio_runtime.block_on(
        core.core
            .shutdown_worker(utf8_string_ref(task_queue, task_queue_len)),
    )
}
