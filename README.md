![Gobol_logo](https://user-images.githubusercontent.com/12091079/56704865-a1293900-66c3-11e9-8c36-12ac2c585c0d.png)

# Gobol a high level language for Data analyst written in Go 

The proposed Gobol language is designed as a high level language optimized for the movement and manipulation of data.


GOBOL is inspired by verbs and features found in [COBOL](https://en.wikipedia.org/wiki/COBOL), [Go](https://golang.org/), [Python](https://www.python.org/) and the deprecated Warehouse language from [Taurus Software](https://taurus.com/) specifically. 

[GOBOL Language Reference Guide](https://github.com/the-kompany/gobol/wiki) 


## Install

You can just simply download the binary from the bin directory for your platform. Then create a Gobol file and write code in it. For example

```
MOVE "Gobol" TO VAR1 
DISPLAY VAR1 
DISPLAY "Hello world!"
MOVE "gobol is great" TO VAR2 

//this is a comment
UPSHIFT(VAR2)     //another comment 
DISPLAY VAR2 

MOVE "Gobol" TO VAR4
MOVE "Gobol" to VAR5
IF VAR4 = VAR5 THEN
   DISPLAY "OK"
   
END-IF

```

Then run it like this:

```
./gobol hello.gbl 
```

## Build

To build it from the source code you will need Go isntalled on your system. This commnd will build a binary for your specific platform in the bin directory.  

```
make build 
```
Or you can make a build for a specific platform like this 

```
make build-windows
```

## Running test 

```
make test 
```  
Or you can run 

```
go test ./... 

```

## Test and build 

```
make all 
```

## [Contributor](https://github.com/the-kompany/gobol/graphs/contributors) 

* [Shawn Gordon](https://github.com/smga3000)
* [Monir Zaman](https://github.com/monirz)  


[Contribution guidelines for this project](docs/CONTRIBUTING.md)


