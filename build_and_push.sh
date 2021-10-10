#!/usr/bin/env bash

# build for ARM
docker build . -t joan41868/grpc-chat-server-arm -f rp4.dockerfile
docker push joan41868/grpc-chat-server-arm:latest


# build for x86
docker build . -t joan41868/grpc-chat-server -f dockerfile
docker push joan41868/grpc-chat-server:latest