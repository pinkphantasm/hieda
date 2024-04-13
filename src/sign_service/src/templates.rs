use core::panic;

use axum::{http::StatusCode, response::Html};
use lazy_static::lazy_static;
use tera::{Context, Tera};

use crate::tool::Tool;

lazy_static! {
    pub static ref TEMPLATES: Tera = {
        match Tera::new("views/**/*.html") {
            Ok(t) => t,
            Err(e) => panic!("Parsing error(s): {}", e),
        }
    };
}
