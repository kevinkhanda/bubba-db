# Graph DB Go implementation
Project is developed as an assignment for Advanced Databases course in Innopolis University

## DB Storage structure
<img src="docs/storage.png" height="250" width="500">

## Requirements
Apple macOS: Install [Go](https://storage.googleapis.com/golang/go1.9.darwin-amd64.pkg)

Microsoft Windows: Install [Go](https://storage.googleapis.com/golang/go1.9.windows-amd64.msi)

Linux: Install [Go](https://storage.googleapis.com/golang/go1.9.linux-amd64.tar.gz)

### Installation
```bash
$ git clone https://github.com/KKhanda/graph-db.git
```

### Building and running the application
```bash
$ go build ./cmd/graph-db/
$ ./graph-db
```

### Running tests
```bash
$ go test ./... -cover
```

## API description
Example of API usage is attached in "main.go" file.

API calls are described in package "/api" in "storage-controller.go" file:

* `api.CreateDatabase` - takes two parameters: database title and db mode ("local", "distributed")
* `api.SwitchDatabase` - parameter is database title. Switches only if database exist.
* `api.DropDatabase` - parameter is database title. Drops database if it exists. 


