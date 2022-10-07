docker build -t gin-api .
docker stop gin-api
docker rm gin-api
docker run -p 8080:8080 -d --name gin-api gin-api