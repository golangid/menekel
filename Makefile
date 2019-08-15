ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

# Database
MYSQL_USER ?= user
MYSQL_PASSWORD ?= password
MYSQL_ADDRESS ?= 127.0.0.1:33060
MYSQL_DATABASE ?= testing

.PHONY: lint-prepare
lint-prepare:
	@echo "Preparing Linter"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

.PHONY: lint
lint:
	./bin/golangci-lint run \
		--exclude-use-default=false \
		--enable=golint \
		--enable=gocyclo \
		--enable=goconst \
		--enable=unconvert \
		./...

.PHONY: mockery-prepare
mockery-prepare:
	@echo "Remove the existing one if exists"
	@rm -rf $(GOPATH)/bin/mockery
	@echo "Installing mockery"
	@go install github.com/vektra/mockery/.../

mockery-generate: 
	@$(GOPATH)/bin/mockery -name ArticleRepository
	@$(GOPATH)/bin/mockery -name ArticleUsecase 

.PHONY: mysql-test-up
mysql-test-up:
	@docker-compose up -d mysql_test

.PHONY: mysql-down-test
mysql-down-test:
	@docker-compose stop mysql_test

.PHONY: test
test: mysql-test-up
	@go test -v -race ./...
	@make mysql-down-test

.PHONY: unittest
unittest:
	@go test -v -short -race ./...

.PHONY: menekel
menekel:
	@go build -o menekel app/main.go

.PHONY: docker
docker:
	@docker build . -t menekel:latest

.PHONY: run
run:
	@docker-compose up -d
	@make mysql-down-test

.PHONY: stop
stop:
	@docker-compose down

.PHONY: migrate-prepare
migrate-prepare:
	@go get -u github.com/golang-migrate/migrate/v4
	@go build -a -o ./bin/migrate -tags 'mysql' github.com/golang-migrate/migrate/v4/cli

.PHONY: migrate-up
migrate-up:
	@migrate -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_ADDRESS))/$(MYSQL_DATABASE)" \
	-path=internal/database/mysql/migrations up	

.PHONY: migrate-down
migrate-down:
	@migrate -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_ADDRESS))/$(MYSQL_DATABASE)" \
	-path=internal/database/mysql/migrations down
 