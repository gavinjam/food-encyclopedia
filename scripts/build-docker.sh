#!/bin/bash

# Make sure to be in the build folder
cd ../docker
cp first.Dockerfile Dockerfile
mv Dockerfile ../

cd ../
docker-compose stop

docker image rm -f food_image
docker build -f Dockerfile -t food_image .

# -d: docker running in background and shell into docker exits immediately
# Without -d, enter inside docker container while running (after leaving the 
# container, it is not running anymore)
# docker run -it --rm -d --name test-app --entrypoint=/bin/bash 84ed84f93d6c
docker run -it --rm -d --name test-app --entrypoint=bash food_image
# get id and pass into the script
# then run docker cp $docker_id:/Users/gavinjampani/go/src/src .

docker cp test-app:/go/src/food-encyclopedia/food-encyclopedia .

docker kill test-app
docker image rm food_image

rm Dockerfile
