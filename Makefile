all: test build 

build: 
	go build -o bin/gobol -v 
test:
	go test -v ./...
clean:
	rm bin/gobol 
	rm bin/gobol-windows		

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/gobol-windows -v
