#!/bin/bash

env CC=arm-linux-gnueabihf-gcc CXX=arm-linux-gnueabihf-g++  CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7  go build -v  -o deletePersonInfo-armv7