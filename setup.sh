#!/usr/bin/env sh

THISDIR=$(dirname "$(readlink -f "$0")")
mkdir -p "$HOME/.local"
ln -snvf "$THISDIR/uni" "$HOME/.local/bin"
sudo chmod +x "$HOME"/.local/bin/*
