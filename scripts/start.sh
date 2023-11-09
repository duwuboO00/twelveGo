#!/bin/bash
# 編譯應用程序
go build -o ./build/server ./cmd/api 

# 運行應用程序
./build/server
