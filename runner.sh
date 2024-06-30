#!/bin/sh

set -e
tmpFile=$(mktemp)
go build -o "$tmpFile" ./*.go
chmod +x "$tmpFile"
exec "$tmpFile" "$@"
