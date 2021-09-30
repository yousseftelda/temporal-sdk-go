#![allow(non_camel_case_types)]

extern crate libc;

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
    core: std::sync::Arc<dyn temporal_sdk_core::Core>,
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
pub struct tmprl_core_worker_config_t {
    /// UTF-8, not null terminated, not owned
    task_queue: *const libc::c_char,
    task_queue_len: libc::size_t,
    // TODO(cretz): Other stuff
}

#[no_mangle]
pub extern "C" fn tmprl_core_register_worker(
    core: *mut tmprl_core_t,
    config: *const tmprl_core_worker_config_t,
) -> *mut tmprl_error_t {
    match core_register_worker(core, config) {
        Ok(_) => std::ptr::null_mut(),
        Err(err) => Box::into_raw(Box::new(err)),
    }
}

fn core_register_worker(
    core: *mut tmprl_core_t,
    config: *const tmprl_core_worker_config_t,
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
