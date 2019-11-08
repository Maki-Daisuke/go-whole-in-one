#!/bin/sh

go build -o bin/wio-init ../wio-init
go build -o bin/wio-generate ../wio-generate
PATH="./bin:${PATH}" wio-generate
if [ $1 = "install" ]; then
    go install
else
    go build
fi
