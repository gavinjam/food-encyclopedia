#!/bin/bash

cd ../
docker-compose stop
docker rm -f food_web
docker rm -f food_db

cd docker
cp second.Dockerfile Dockerfile
mv Dockerfile ../

cd ../

# modify docker-compose.yml
# sed -i -e 's/test/app/g' docker-compose.yml
# Enter insides a running docker container: docker exec -it food_web /bin/sh
docker build -f Dockerfile -t food_image .
# docker run -it --rm --name test-app 
docker-compose up -d
# Will need docker-compose stop to stop it
rm Dockerfile
rm food-encyclopedia