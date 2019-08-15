# menekel
menekel `/'m(ə)n(ə)k(ə)l/` stands for : `menej artikel` is a sample of article management services build from Golang.

## Index
- [Contribution](#contribution)
- [Run The Project](#run-the-project)

### Contribution 
You can file an [issue](https://github.com/golangid/menekel/issues/new) or submit a Pull Request

### Testing
**Integration Test**
```bash
$ make test
```
**Unit Test**
```bash
$ make unittest
```

### Run The Project

```bash
# Dockerize the app
$ make docker

# Create the config file
$ cp config.toml.example config.toml

# run the project
$ make run

# Migrate the schema
$ make migrate-prepare
$ make migrate-up
```

Now the application should be active. Try to access it.

```bash
$ curl http://localhost:9090/articles
```