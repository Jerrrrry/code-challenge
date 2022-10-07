docker build -t web .
docker stop web
docker run -it -p 80:80 --rm -d --name web web

