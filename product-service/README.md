# Introduction-to-Microservices


### Running

The applicaiton can be run with `go run`

```
➜ go run main.go
products-api 2020/02/16 16:15:11 Starting server on port 9090
curl localhost:9090/products
```

----------------------------

### Generating Swagger Documentation

Swagger documentation is generated from the code annotations inside the source using go-swagger.

Go swagger can be installed with the following command:

```
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

You can generate the documentation using the command:

```
swagger generate spec -o ./swagger.yaml --scan-models
```

After running the application:

```
go run main.go
```

Swagger documentation can be viewed using the ReDoc UI in your browser at [http://localhost:9090/docs](http://localhost:9090/docs).