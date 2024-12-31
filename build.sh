#!/bin/bash

build_uni() {
    local pkgname="$1"
    echo "Building [ $pkgname ] for Linux..."
    GOOS=linux GOARCH=amd64 go build -o "./uni/$pkgname" "./cmd/$pkgname"
    echo "Building [ $pkgname ] for Linux completed."
}

build_win() {
    local pkgname="$1"
    echo "Building [ $pkgname ] for Windows..."
    GOOS=windows GOARCH=amd64 go build -o "./win/$pkgname.exe" "./cmd/$pkgname"
    echo "Building [ $pkgname ] for Windows completed."
}

if [ $# -eq 1 ]; then
    build_uni "$1"
else
    for d in ./cmd/*; do
        if [ -d "$d" ]; then
            build_uni "$(basename "$d")"
            build_win "$(basename "$d")"
            echo "------------------------------------------"
        fi
    done
fi
