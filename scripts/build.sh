#!/bin/bash

if [ ! -d "./cmd" ] || [ ! -d "./build" ]; then 
    echo "must be run from project root directory"
    exit 1
fi

CMD=$1

if [ -z "${CMD}"]; then 
    echo "cmd argument required"
    exit 1
fi

if [ ! -d "./cmd/${CMD}" ]; then
    echo "./cmd/${CMD} command not found"
    exit 1;
fi

docker build -t dyolcekaj/images:discord-${CMD} --build-arg COMMAND_NAME=${CMD} -f build/Dockerfile .