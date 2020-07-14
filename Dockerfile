# Builder
FROM golang:1.14.0-stretch as builder

RUN apt --yes --force-yes update && apt --yes --force-yes upgrade && \
    apt --yes --force-yes install git \
    make openssh-client

WORKDIR /app

COPY . .

RUN make menekel

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app /menekel

WORKDIR /menekel

EXPOSE 9090

COPY --from=builder /app/menekel /app

CMD /app/menekel http