docker rm -f $(docker ps -aq)
docker rmi pinggo
docker build . -t pinggo
