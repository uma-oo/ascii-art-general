#!/bin/bash
docker rmi -f $(docker images -aq)

# Remove all Docker containers
docker rm -f $(docker ps -aq)

# Build Our image, (--no-cache)=> each build step will be executed without retrieving already stored data.
docker build --no-cache -t ascii .

docker run -d -p 6500:6500 --name a1 ascii

# this command to execute a command in our container (i)=> let the standard input open, (t)=> create a virtual terminal in Our container
# docker exec -it a1  /bin/bash/