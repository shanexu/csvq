use std::ffi::{c_char, CString};

pub fn add(left: u64, right: u64) -> u64 {
    left + right
}

#[no_mangle]
pub extern "C" fn process_data(_input: *const c_char) -> *mut c_char {
    // Process the data
    // Remember to ensure memory safety!

    let output = "hello world";
    CString::new(output).unwrap().into_raw()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = add(2, 2);
        println!("hello world");
        assert_eq!(result, 4);
    }
}
