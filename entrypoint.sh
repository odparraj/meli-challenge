#!/bin/sh

go run main.go artisan migrate

CompileDaemon --build="go build main.go" --command=./main

