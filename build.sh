#!/bin/bash
# shellcheck disable=SC2164
cd ./internal/
GOOS=linux GOARCH=amd64 go build  -o ../my_hotel_merger
