all: test build 

build: 
	go build  -o bin/gobol -v cmd/gobol/gobol.go 
test:
	go test -v ./...
clean:
	rm bin/gobol 
	rm bin/gobol-windows		

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/gobol-windows -v cmd/gobol/gobol.go 
