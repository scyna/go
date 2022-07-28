#!/bin/bash
if [ "$1" != '' ];
then
    image=registry.vin3s.vn:5000/scyna/manager:$1
    docker build -t ${image} -f Dockerfile .
    docker push ${image}
else
    echo 'Please pass version of image!'
fi