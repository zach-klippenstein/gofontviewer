#!/bin/sh

# Exit on error.
set -e

echo "Building fonts…"
go generate -x fonts/fonts.go

echo
echo "Building javascript…"
gopherjs build -v
