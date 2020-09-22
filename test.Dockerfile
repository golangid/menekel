# Builder
FROM golang:1.14.0-stretch

RUN apt --yes --force-yes update && apt --yes --force-yes upgrade && \
    apt --yes --force-yes install git \
    make openssh-client

WORKDIR /app

COPY . .

RUN make full-test
