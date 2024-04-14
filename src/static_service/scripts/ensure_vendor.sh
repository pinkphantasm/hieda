#!/usr/bin/env bash

VENDOR_DIR="public/vendor"

# Check if vendor directory does not exist.
if [ ! -d "$VENDOR_DIR" ]; then
    echo "error: no vendor directory"
    echo "please run:"
    echo
    echo "mkdir -p \"$VENDOR_DIR\""
    exit 1
fi

# Fonts
FONTS_DIR="$VENDOR_DIR/fonts"
FONTS_INSTALLATION_URL="https://gwfh.mranftl.com/fonts"

FONTS_LIST=(
    "comfortaa-v45-cyrillic_latin-regular.woff2"
    "jetbrains-mono-v18-cyrillic_latin-regular.woff2"
    "roboto-v30-cyrillic_latin-700.woff2"
    "roboto-v30-cyrillic_latin-700italic.woff2"
    "roboto-v30-cyrillic_latin-700italic.woff2"
    "roboto-v30-cyrillic_latin-regular.woff2"
)

# Check if vendor fonts directory does not exist.
if [ ! -d "$FONTS_DIR" ]; then
    echo "error: no fonts directory"
    echo "please run:"
    echo
    echo "mkdir -p \"$FONTS_DIR\""
fi

# Check fonts
for FONT in "${FONTS_LIST[@]}"; do
    if [ ! -f "$FONTS_DIR/$FONT" ]; then
        echo "error: font \"$FONT\" not found"
        echo "expected path to exist: \"$FONTS_DIR/$FONT\""
        echo
        echo "please install it:"
        echo "$FONTS_INSTALLATION_URL"
        exit 1
    fi
done

# JavaScript Libraries
JS_DIR="$VENDOR_DIR/js"

# Check if JavaScript directory does not exist.
if [ ! -d "$JS_DIR" ]; then
    echo "error: no JavaScript directory"
    echo "please run:"
    echo
    echo "mkdir -p \"$JS_DIR\""
    exit 1
fi

# Alpine.js
ALPINE_PATH="$JS_DIR/alpine-3.13.7.min.js"
ALPINE_INSTALLATION_URL="https://cdn.jsdelivr.net/npm/alpinejs@3.13.7/dist/cdn.min.js"

if [ ! -f "$ALPINE_PATH" ]; then
    echo "error: Alpine.js not found"
    echo "expected path to exist: \"$ALPINE_PATH\""
    echo
    echo "please install it:"
    echo "$ALPINE_INSTALLATION_URL"
    exit 1
fi

# HTMX
HTMX_PATH="$JS_DIR/htmx.min.js"
HTMX_INSTALLATION_URL="https://unpkg.com/htmx.org/dist/htmx.min.js"

if [ ! -f "$HTMX_PATH" ]; then
    echo "error: HTMX not found"
    echo "expected path to exist: \"$HTMX_PATH\""
    echo
    echo "please install it:"
    echo "$HTMX_INSTALLATION_URL"
    exit 1
fi

# HTMX response-targets
HTMX_EXT_RT_PATH="$JS_DIR/htmx-response-targets.js"
HTMX_EXT_RT_INSTALLATION_URL="https://unpkg.com/htmx.org/dist/ext/response-targets.js"

if [ ! -f "$HTMX_EXT_RT_PATH" ]; then
    echo "error: HTMX extension response-targets not found"
    echo "expected path to exist: \"$HTMX_EXT_RT_PATH\""
    echo
    echo "please install it:"
    echo "$HTMX_EXT_RT_INSTALLATION_URL"
    exit 1
fi

echo "All requirements have been satisfied!"
