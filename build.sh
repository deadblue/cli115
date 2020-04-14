#!/bin/sh

mkdir -p output/bin

for path in `find ./cmd -d 1 -type d`; do
  cmd=`basename $path`
  echo "Building app: ${cmd}"
  go build -o output/bin/${cmd} -ldflags="-s -w" $path/*.go
done