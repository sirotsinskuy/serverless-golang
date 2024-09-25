#!/bin/bash

# cd to the directory of the script
cd "$(dirname "$0")" || exit

# save the working directory
WD=$(pwd)/..
export WD

# build hello_wold
cd "$WD"/services/hello_world || exit
GOARCH=arm64 GOOS=linux go build -o bootstrap
zip hello_world.zip bootstrap

# deploy
cd "$WD"/services || exit
sls deploy --stage dev
