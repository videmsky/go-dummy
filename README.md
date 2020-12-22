## Branch & Build Training

Running the project:
* `go run .`
* http server listens on `localhost:8080`

Configure dependencies:
* `go mod init go-dummy`
* `go get -u -v ./...`

Testing the project:
* `go test -v -cover ./...`

Building the project:
* `go build -o app`

Build the docker image
* `docker build -t go-dummy .`

Run the docker container (port mappings... host:container)
* `docker run -p 8081:8080 -it go-dummy`
* http server listens on `localhost:8081`
