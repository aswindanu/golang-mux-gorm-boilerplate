FROM golang:alpine as builder

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go build

EXPOSE 3000

ENTRYPOINT ["./golang-simple"]