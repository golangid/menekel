# Builder
FROM golang:1.14.0-stretch

RUN apt update && apt upgrade && \
    apt install git gcc make

WORKDIR /app

COPY . .

RUN make full-test
