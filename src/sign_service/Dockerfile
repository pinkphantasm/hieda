FROM rust:1.75

RUN USER=root cargo new --bin index_service

WORKDIR /index_service

COPY ./Cargo.lock ./Cargo.lock
COPY ./Cargo.toml ./Cargo.toml

RUN cargo build --release

RUN rm src/*.rs

COPY ./src ./src
COPY ./views ./views

RUN rm ./target/release/deps/index_service*

RUN cargo install --path .

CMD ["index_service"]
