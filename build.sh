#!/bin/bash

echo "pulling latest branch"
git pull
echo "building new docker image"
docker build . -t findingway
echo "build done"
