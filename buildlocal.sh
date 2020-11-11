#!/bin/bash

docker build -t my-golang-app . -f Dockerfile.local --mount type=bind,source="$(pwd)"/bin,target=/binout