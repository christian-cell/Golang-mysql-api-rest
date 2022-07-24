Flyweight docker container with golang:1.16-alpine :))

$ docker-compose build

$ docker-compose up -d

install mysql-driver and mux router generator

change module name if you want

$ docker-compose exec api github.com/go-sql-driver/mysql

$ docker-compose exec api github.com/gorilla/mux

$ docker-compose exec api go mod tidy

last thing : CompileDaemon is already installed in docker image "take a look in the Dockerfile",
ok ... , AS hot-reloader it's ok , but sometimes does not render errors in run time , in that case i use

$ docker logs --since 30s -f <container_name_or_id>

check and replace container_name_or_id , you'll get all the errors


export GO111MODULE="on"
