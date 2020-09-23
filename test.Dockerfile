# Builder
FROM golang:1.14.0-stretch

RUN apt --yes --force-yes update && apt --yes --force-yes upgrade && \
    apt --yes --force-yes install git \
    make openssh-client

WORKDIR /app

COPY . .
ENV MYSQL_TEST_URL=root:root@tcp(mysql_test:3306)/testing?parseTime=1&loc=Asia%2FJakarta&charset=utf8mb4&collation=utf8mb4_unicode_ci
RUN make full-test
