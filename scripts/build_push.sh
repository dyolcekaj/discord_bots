#!/bin/bash

source ./scripts/common.sh

docker_build $1 
docker_push $1
