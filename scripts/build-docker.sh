#!/bin/bash

# Make sure to be in the build folder
cd ../docker
cp first.Dockerfile Dockerfile
mv Dockerfile /Users/gavinjampani/go/src

cd /Users/gavinjampani/go/src
docker-compose stop

docker image rm -f food_image
docker build -f Dockerfile -t food_image .

# -d: docker running in background and shell into docker exits immediately
docker run -it --rm -d --name test-app --entrypoint=/bin/bash food_image
# get id and pass into the script
# then run docker cp $docker_id:/Users/gavinjampani/go/src/src .

docker cp test-app:/Users/gavinjampani/go/src/src .

docker kill test-app
docker image rm food_image

rm Dockerfile
