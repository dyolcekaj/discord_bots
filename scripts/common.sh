#!/bin/bash

REPO="dyolcekaj/images"
IMG_PREFIX="discord"

check_root() {
    if [ ! -d "./cmd" ] || [ ! -d "./build" ]; then 
        echo "must be run from project root directory"
        exit 1
    fi
}

check_cmd_arg() {
    if [ -z "$1" ]; then 
        echo "cmd argument required"
        exit 1
    fi

    if [ ! -d "./cmd/$1" ]; then
        echo "./cmd/$1 command not found"
        exit 1;
    fi
}

docker_build() {
    CMD=$1

    check_root
    check_cmd_arg ${CMD}

    IMG="${REPO}:${IMG_PREFIX}-${CMD}"
    docker build -t ${IMG} --build-arg COMMAND_NAME=${CMD} -f build/Dockerfile .
}

docker_push() {
    CMD=$1

    check_root
    check_cmd_arg ${CMD}

    IMG="${REPO}:${IMG_PREFIX}-${CMD}"
    docker image inspect ${IMG} &>/dev/null

    if [ $? -ne 0 ]; then  
        echo "no image for ${CMD} exists"
    fi

    docker push ${IMG}
}