# menekel
menekel `/'m(ə)n(ə)k(ə)l/` stands for : `menej artikel` is a sample of article management services build from Golang.

## Index
- [Contribution](#contribution)
- [Run The Project](#run-the-project)

### Contribution 
You can file an [issue](https://github.com/golangid/menekel/issues/new) or submit a Pull Request

### Run The Project

Before run the aplication, make sure you already migrate the `article.sql` to your database. And edit the `config.json` based on your needs.

```bash
#move to directory
cd $GOPATH/src/github.com/golangid

# Clone into YOUR $GOPATH/src
git clone https://github.com/golangid/menekel.git

#move to project
cd menekel

# Install Dependencies
dep ensure

# Make File
make

# Run Project
./menekel http

```

Or

```bash
# GET WITH GO GET
go get github.com/golangid/menekel

# Go to directory

cd $GOPATH/src/github.com/golangid/menekel

# Install Dependencies
dep ensure

# Make File
make

# Run Project
./menekel http
```
