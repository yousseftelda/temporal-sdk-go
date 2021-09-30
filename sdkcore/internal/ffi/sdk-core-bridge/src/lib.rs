extern crate libc;

#[no_mangle]
pub extern "C" fn hello_rust() {
    println!("Hello, Rust!");
}
