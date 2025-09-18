# Go


## Start project in Go: Create a Module
```bash
mkdir myapp && cd myapp
go mod init github.com/cmrohityadav/myapp

# for learning
go mod init main

```
- **Equivalent of `npm init` in Go**

```bash
go mod init <module_name>
```

- `go mod init` is used to initialize a new module.
-  It creates a new `go.mod` file in the current directory.
-  The `go.mod` file contains information about the module, its dependencies, and the Go version.

## To Run go program

```bash
go run <main_file.go>
```
- **Equivalent of `npm run start` in Go**





## Compiled
- Go tool can run file directly , without VM
- Executable are different for OS

## what
- System apps to web apps - Cloud

### Don't bring baggage
- Yeah , I did that esrlier
- Similarity with lots of language C JAVA PASCAL


# Object Oriented 
- YES and NO

# Missing 
- No try catch
- lexer does a lot of work

# Run Code
```bash
go run main.go
go mod init main
go env
GOOS="windows" go build
go mod tidy
go mod verify
go list
go list all
go list -m all
go mod vendor
go run -mod=vendor main.go

```

# types
- case sensitive almost
- Variable type sholud be known in advance
- Everything is type


# Run Code
bash
go run main.go
go mod init main
go env
GOOS="windows" go build



# types
- case sensitive almost
- Variable type sholud be known in advance
- Everything is type

# Memory Management
- Memory allocation and Deallocation happens automatically
- Garbage Collector happens automatically
- Out of Scope or nil
## new()
- Allocate memory but not initialize
- you will get a memory address
- zeroed storage

## make()
- Allocate memory but initialize
- you will get a memory address
- non-zeroed storage