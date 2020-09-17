#!/bin/bash

cd /Users/gavinjampani/go/src
# docker-compose stop

cd /Users/gavinjampani/go/docker
cp second.Dockerfile Dockerfile
mv Dockerfile /Users/gavinjampani/go/src

cd /Users/gavinjampani/go/src

# modify docker-compose.yml
# sed -i -e 's/test/app/g' docker-compose.yml
# docker exec -it food_web /bin/sh
docker build -f Dockerfile -t food_image .
# docker run -it --rm --name test-app 
docker-compose up -d
# Will need docker-compose stop to stop it
rm Dockerfile
rm src