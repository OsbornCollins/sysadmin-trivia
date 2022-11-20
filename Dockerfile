# Install golang to the docker container
FROM golang:1.19.0

# The directory of Docker since docker is linux based
WORKDIR /usr/src/app

# Air package allows us to hot reload the project
RUN go install github.com/cosmtrek/air@latest

COPY . .
# To allow all pakcages and interdepencies to be loaded
RUN go mod tidy