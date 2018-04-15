#!/usr/bin/env bash
env GOOS=linux GOARCH=amd64 go build -o /tmp/main github.com/rwirdemann/3skills.time/lamda
zip -j /tmp/main.zip /tmp/main

# aws lambda create-function --function-name get-projects --runtime go1.x \
# --role arn:aws:iam::464797721797:role/lambda-books-executor \
# --handler main --zip-file fileb:///tmp/main.zip

aws lambda update-function-code --function-name get-projects --zip-file fileb:///tmp/main.zip

aws lambda invoke --function-name get-projects /tmp/output.json