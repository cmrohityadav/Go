# Go


## Start project in Go: Create a Module

```bash
mkdir myapp && cd myapp
go mod init github.com/cmrohityadav/myapp

# for learning
go mod init main

```
- **Equivalent of `npm init` in nodeJS**

```bash
go mod init <module_name>
```

- `go mod init` is used to initialize a new module.
-  It creates a new `go.mod` file in the current directory.
-  The `go.mod` file contains information about the module, its dependencies, and the Go version.

### **go.mod file**

- **Equivalent of `package.json` in nodeJS**
- It contains information about the module, its dependencies, and the Go version.



```bash
go mod tidy
```
- **Equivalent of `npm install` in nodeJS**
- `go mod tidy` is used to add missing and remove unused modules.
- It updates the go.mod file to use the latest version of the dependencies.

## To Run go program

```bash
go run <main_file.go>
```
- **Equivalent of `npm run start` in nodeJS**

## Installing package

```bash
go get <package_name>
```
- **Equivalent of `npm install <package_name>` in nodeJS**
- `go get` is not a package manager.
- `go get` is used to download and install packages from remote repositories.
- It does not handle versioning.
- This command fetches the package and its dependencies (if any)


## Primitve Data type
- int, float64, byte, string, rune & bool

| Type      | Example  | Description                                                   |
|-----------|---------|---------------------------------------------------------------|
| **int**   | `42`    | Integer (size depends on architecture: 32- or 64-bit)          |
| **float64** | `3.14` | Floating-point number                                         |
| **byte**  | `'A'`   | Alias for `uint8`; typically used for raw data or ASCII chars   |
| **rune**  | `'✓'`   | Alias for `int32`; represents a Unicode code point              |
| **string**| `"hi"`  | Immutable sequence of bytes (UTF-8 encoded text)                |
| **bool**  | `true`  | Boolean logic (`true` or `false`)                               |

## Variable

```go

// Explicit
var name string="rohit";


// Implicit type [Type inference (Go infers the type)]
var website = "cmrohityadav.in";

// No var style
// Short declaration (only inside functions)
numberOfUser := 30000;


```

## Constant
- constant is a name that’s bound to a fixed value at compile time.
Once defined, it cannot change during the lifetime of the program

```go
// Untyped constant
const pi=3.14;

// Typed constant
const maxUser int=420;

// Multiple constants in a block
const (
    portNumber=4000;
    greeting="Good Morning"
)


// iota – Constant Generator

type Weekday int

const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

```
- if we use **first Letter Capital** the it like: **Public**
```go
const LoginToken string = "abcdefg"
```

- numberOfUser := 30000 // This code is not allow in public,it can be only use in method / function

## Conditionals

### Basic if Statement
```go
age:=10;

if age>=18 {
    fmt.Println("You are Adult");
}
```
- Condition must evaluate to a boolean (true or false).

- Curly braces {} are mandatory even for a single line of code.

- No parentheses around the condition

### if-else Statement

```go
score:=65;

if score>=90 {
    fmt.Println("Grade A");
} else if score>=75 {
    fmt.Println("grade B");
} else if score>=50 {
    fmt.Println("grade C");
}else {
    fmt.Println("Grade F");
}

```
- else if handles multiple conditions sequentially

### Short Statement in if

- Go allows you to declare a variable inside the if condition
- which is only visible inside that block

```go
if temp:=30 ; temp>40 {
    fmt.Println("Its a hot day");
} else {
    fmt.Println("It cmfortable");
}
```
- temp exists only inside this if-else block.
- Use logical operator when possible [&&(and),||(or),!(not)]
- Go does not have ternary operator, you will have to use normal if else

## Switch

### Basic Switch
```go
day := "Tuesday"

    switch day {
    case "Monday":
        fmt.Println("Start of the week")
    case "Tuesday":
        fmt.Println("Second day of the week")
    case "Friday":
        fmt.Println("Almost weekend")
    default:
        fmt.Println("Midweek days")
    }
```
- default is optional but recommended
- No break needed: Go automatically stops after the first match

### Multiple Cases in One Line
```go
score := 85

switch {
case score >= 90:
    fmt.Println("Grade A")
case score >= 75 && score < 90:
    fmt.Println("Grade B")
case score >= 50 && score < 75:
    fmt.Println("Grade C")
default:
    fmt.Println("Grade F")
}
```
- Go allows condition expressions, not just simple values.
- Great for ranges or complex checks

### Type Switch (Advanced, for Interfaces)
```go
var val interface{} = 42

switch v := val.(type) {
case int:
    fmt.Println("Integer:", v)
case string:
    fmt.Println("String:", v)
default:
    fmt.Println("Unknown type")
}
```
- Checks the dynamic type of an interface.
- Essential for writing generic libraries or handling polymorphic data

### Fallthrough
```go
num := 2

switch num {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
    fallthrough
case 3:
    fmt.Println("Three")
}

```
- fallthrough forces execution to the next case, even if it doesn’t match.
- Use carefully—most of the time, Go’s default “break after case” is safer





## Loops
- Unlike some languages, Go has only one loop keyword: **for**
- Everything else is just a variation

### Basic for loop
```go
for i:=0; i<5; i++ {
    fmt.Println("Iteration : ",i);
}
```
### While Style Loop
```go
count:=0;
for count<5 {
    fmt.Println("counter: ",count);
}
```
### Infinite Loop
```go
for{
    fmt.PrintLn("Hello")
}
```
### Loop with **range** 
- Over Slices, Arrays, Maps, Strings,channels
- range is Go’s idiomatic way to iterate collections
- Returns two values: index/key and value
- **break** → exit the loop immediately

- **continue** → skip the rest of the current iteration and move to the next
```go
for i:=range 11{
    fmt.Println(i);
}
//0,1,2....10
```

```go
nums := []int{10, 20, 30, 40}
for index, value := range nums {
    fmt.Println("Index:", index, "Value:", value)
}
```
## Arrays
- An array is a fixed-size, ordered collection of elements all of the same type







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