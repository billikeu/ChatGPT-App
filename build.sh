####!/bin/bash

# 获取当前操作系统的名称
version="v1.1.1"

# 构建应用程序
# wails build -ldflags="-s -w" -platform="darwin/amd64" -o "ChatGPT-App-$version"
# wails build -ldflags="-s -w" -platform="linux/amd64" -o "ChatGPT-App-$version"
wails build -ldflags="-s -w" -platform="windows/amd64" -o "ChatGPT-App-$version.exe"