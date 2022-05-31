# IMPORTANT NOTE 
# https://github.com/chemidy/smallest-secured-golang-docker-image/issues/5

FROM golang:alpine as builder

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go build

EXPOSE 3000

ENTRYPOINT ["./golang-simple"]