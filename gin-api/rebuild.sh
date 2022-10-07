docker build -t gin-api .
docker stop gin-api
docker rm gin-api
docker run --net gin-api -p 8080:8080 -d --name gin-api gin-api