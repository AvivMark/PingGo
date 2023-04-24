docker rm -f $(docker ps -aq)
docker rmi pinger-ui
docker build . -t pinger-ui
