use core::panic;

use lazy_static::lazy_static;
use tera::Tera;

lazy_static! {
    pub static ref TEMPLATES: Tera = {
        match Tera::new("views/**/*.html") {
            Ok(t) => t,
            Err(e) => panic!("Parsing error(s): {}", e),
        }
    };
}
