# Go
## Content
---
[File](##start-project-in-go-create-a-module)
[File](#file)
---

---

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


## Pointers
- A pointer is a variable that stores the memory address of another variable
- Instead of copying a value, you pass around a “reference” to where that value lives in memory
- Think of it like sharing a Google Drive link instead of sending a huge video file. You can edit the video in one place and everyone sees the update
### compare with c/c++
```bash
# c/c++
int* pMyPointer
int *pMyPointer

# go
var pMyPointer *int

# note: here’s no alternative syntax
var name *Type 
```
- get address/reference : &anyVariable
- dereference: *anyPointerVariable
### example of pointers
```go
x := 42
fmt.Println(x) //42
fmt.Println(&x) // 0xc0000120a0 
p := &x        // &x → address of x
fmt.Println(p) // 0xc0000120a0
fmt.Println(*p) // 42
y := *p        // *p → value stored at that address (42)
fmt.Println(y) //42
```
```go
var x int          // declare x as int, default value is 0
    x = 42             // assign a value

    fmt.Println(x)     // 42  (value of x)

    fmt.Println(&x)    // 0xc0000120a0 (address of x — varies each run)

    var p *int         // declare p as pointer to int
    p = &x             // assign address of x to p
    fmt.Println(p)     // same as &x (pointer value)

    fmt.Println(*p)    // 42  (value stored at the address p points to)

    var y int          // declare y as int
    y = *p             // copy the value pointed to by p
    fmt.Println(y)     // 42
```
- **func changeByReference(&num int) is: ❌ Invalid in Go**
```bash
C++	void f(int &num) → call with f(x)

C	void f(int *num) → call with f(&x)
Go	func f(num *int) → call with f(&x)
```



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
- Define size (mandatory in Go)
```go
var arr[4]int
// array length
fmt.Println(len(arr));

arr[0]=1;

fmt.Println(arr[0]);

fmt.Println(arr);



var vals[4]bool;
vals[2]=true;
fmt.Println(vals);
// [false false true false]
// bydefault all will false
// int -> 0, string-> "",bool ->false

nums:=[3]string{"rohit","yadav","yadav"};


// 2D array
nums:=[2][2]int{{3,4},{5,6}};





```
## Slice
- Dynamic Array
```go
nums:=[]string{"rohit","yadav"};



var rollNo[]int;
// Above uninitialized slice is nill

var prices=make([]int,2);
//it set all 2 element to Zero: eg.  [0 0]

fmt.Println(cap(prices));
// capacity -> maximium numbers of elements can fit






```


## Map
- A map is Go’s built-in hash table or dictionary type.
- It stores data as key–value pairs so you can quickly look up a value by its key

### create
```go
//Using make
m:=make(map[keyType]keyValueType);
m:=make(map[int]string);

// using Map literal
m:=map[string]int{"rohit":1,"yadav":2};

// declareing then assigning
var m map[string]int;
m=make(map[string]int);


```
### Access
```go 
fmt.Println(m["rohit"]);
// if key does not exists in the map then it returns ZERO value

```
### Add/update
```go 
m["india"]=91;

```
### Iteration
```go
for k,v:=range m {
    fmt.Printf("For key %v, value is %v\n", key, value)
}

```

### Deletion
```go
delete(m,"rohit");
```
### Checking key
```go
m:=map[int]string{1:"rohit",2:"yadav"};

k,ok:=m[1];
// k:value of key,ok:bool
if ok {
    fmt.Println("all okay");
}else {
    fmt.Println("Not Okay");
}
```


## Struct
- Struct is a way to group together different pieces of data (fields) under one name
- Think of it like a blueprint for an object
```go
package main

import (
    "fmt"
    "time"
)

// Define the struct
type order struct {
    id        string
    amount    float32
    status    string
    createdAt time.Time // nanoseconds precision
}

func main() {
    // ===============================
    // WAY 1: Zero value / default struct
    // ===============================
    var mobileOrder order
    mobileOrder.id = "101"
    mobileOrder.amount = 12000.50
    mobileOrder.status = "pending"
    mobileOrder.createdAt = time.Now()
    fmt.Println("Mobile Order:", mobileOrder)

    // ===============================
    // WAY 2: Struct literal with field names
    // ===============================
    laptopOrder := order{
        id:        "64",
        amount:    50000.14,
        status:    "on the way",
        createdAt: time.Now(),
    }
    fmt.Println("Laptop Order:", laptopOrder)

    // ===============================
    // WAY 3: Struct literal without field names
    // Must follow the order of fields in struct!
    // ===============================
    bookOrder := order{"201", 999.99, "delivered", time.Now()}
    fmt.Println("Book Order:", bookOrder)

    // ===============================
    // WAY 4: Using pointer to struct
    // ===============================
    penOrder := &order{
        id:        "301",
        amount:    49.99,
        status:    "shipped",
        createdAt: time.Now(),
    }

    // Go auto-dereferences
    fmt.Println("Pen Order (pointer):", penOrder)
    fmt.Println("Pen Order ID via pointer:", penOrder.id) 
    // Dereferencing is automatic for struct field access

    // Modifying through pointer
    penOrder.status = "delivered"
    fmt.Println("Updated Pen Order:", penOrder)

    // ===============================
    // WAY 5: Using new() to create pointer to struct
    // ===============================
    bagOrder := new(order) // returns *order
    bagOrder.id = "401"
    bagOrder.amount = 1200.50
    bagOrder.status = "pending"
    bagOrder.createdAt = time.Now()
    fmt.Println("Bag Order (new):", bagOrder)

    // ===============================
    // WAY 6: Anonymous struct
    // ===============================
    anonymousOrder := struct {
        id     string
        amount float32
    }{
        id:     "501",
        amount: 350.75,
    }
    fmt.Println("Anonymous Order:", anonymousOrder)

    // ===============================
    // WAY 7: Using methods with struct
    // ===============================
    laptopOrder.PrintSummary()
}

// Method attached to struct
func (o order) PrintSummary() {
    fmt.Printf("Order %s of amount %.2f is %s\n", o.id, o.amount, o.status)
}

// (o order) → value receiver → gets a copy of the struct.

// If you change o.status inside this method, the original struct won’t change.

// (o *order) → pointer receiver → gets the address, can modify original struct

```

- **If x is a pointer to a struct, x.f is shorthand for (*x).f**

- Just like constructor in other programming language
```go
func newOrder(id string,amount float32,status)*orde{
    // initial setup goes here...
    myOrder:=order{
        id:id,
        amount:amount,
        status:status,
    }
    return &myOrder;
}

```
### Struct Embedding
```go
type Address struct {
    City  string
    State string
}

type Person struct {
    Name string
    Age  int
    Address   // embedded struct (no field name)
}

``` 
## Interfaces
- It is just Contract
- An interface is a type that specifies a set of method signatures—just the method names and parameters, no code
- If a type (a struct, usually) has all those methods with the same signatures, it automatically satisfies the interface.
- No keywords like “implements” are needed

```go
package main

import "fmt"

// 1. Define the interface (the contract)
type Speaker interface {
    Speak() string
}

// 2. Create types that match the contract
type Person struct{}
func (p Person) Speak() string {
    return "Hello, I'm a person."
}

type Dog struct{}
func (d Dog) Speak() string {
    return "Woof!"
}

// 3. Use the interface as a parameter
func SaySomething(s Speaker) {
    fmt.Println(s.Speak())
}

func main() {
    p := Person{}
    d := Dog{}

    SaySomething(p)  // works
    SaySomething(d)  // works
}


```

| Concept              | C++                               | Go                                          |
|----------------------|------------------------------------|---------------------------------------------|
| Declare behaviour    | Pure virtual class                | `interface`                                  |
| Explicit inheritance | `class Circle : public Shape`     | Nahi hota (automatic / implicit)             |
| Method override      | `override` keyword optional       | Bas method signature match karo              |
| Multiple inheritance | Limited / complex                 | Easy, ek type multiple interfaces satisfy kar sakta |


## Enums
- Go doesn’t have a native enum keyword like C, Java, or Rust.
- Instead, Go gives you simple primitives—constants + iota + custom types—that you combine to achieve the same (and often better) result.
```go
// Step 1: create a named type
type OrderStatus int

// Step 2: declare constants using iota
const (
    Pending OrderStatus = iota  // 0
    Processing                  // 1
    Shipped                     // 2
    Delivered                   // 3
    Cancelled                   // 4
)

```

## Generics
- Generics let you write one function or type that works with many data types, instead of writing the same code multiple times
- Generics = code that works for many types
### Without generics
- two functions that do the exact same thing, only the types differ
```go
func SumInt(a, b int) int {
    return a + b
}
func SumFloat(a, b float64) float64 {
    return a + b
}

```
### With generics
- One function now works for both int and float64
- T is a type parameter—a placeholder for any type that matches the rule int | float64.
```go
func Sum[T int | float64](a, b T) T {
    return a + b
}
```


## Functions
```go
func add(a int,b int) int{
    return a+b
}

// if same datatype of parameter
func add(a ,b int) int{
    return a+b
}

```

```go
// Multiple Returns
func multipleReturn()(string,string,int){
    return "first","second",3
}

first,second,third:=multipleReturn();


// Using the Blank Identifier: _
// If you don’t need all the results:

_, second,third := multipleReturn();
fmt.Println(second, number);

```

### Variadic Function
- Accepts any number of arguments of the same type.
- Inside, the parameter behaves like a slice

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
fmt.Println(sum(1,2,3,4)) // 10
```

### Anonymous / Closure
- Functions without a name can be stored in variables
- They can also close over variables from the surrounding scope
```go
square := func(x int) int {
    return x * x
}
fmt.Println(square(4)) // 16
```

### Functions in Go are first-class citizens
#### Assign to a Variable

```go
package main
import "fmt"

func greet(name string) string {
    return "Hello, " + name
}

func main() {
    // assign function to a variable
    sayHello := greet
    fmt.Println(sayHello("Rohit")) // Hello, Rohit
}
```

#### Pass a Function as an Argument
```go
func operate(a, b int, anyFunName func(int, int) int) int {
    return anyFunName(a, b)
}

func main() {
    add := func(x, y int) int { return x + y }
    fmt.Println(operate(3, 4, add)) // 7
}
```
#### Return a Function (Closure)
```go 
func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func main() {
    times2 := multiplier(2)
    fmt.Println(times2(5)) // 10
}
```
- func(int) int: Return type: instead of a simple type (like int), it returns a function.
That returned function itself takes one int and returns an int


## Goroutines
- Think of a goroutine as a tiny, lightweight worker that Go can run at the same time as your main program
- They’re like “threads,” but cheaper and easier
- You can start thousands of them without eating much memory

## Wait Group
### The Problem WaitGroup Solves
```go
package main

import (
    "fmt"
    "time"
)

func worker(id int) {
    time.Sleep(1 * time.Second)
    fmt.Println("Worker", id, "done")
}

func main() {
    for i := 1; i <= 3; i++ {
        go worker(i)
    }
    // main returns immediately → goroutines may not finish
}

```
- Sometimes you’ll see no output, because main ends and the process exits
### Meet sync.WaitGroup
- A WaitGroup is like a counter + gate
- You “add” the number of goroutines you plan to start
- Each goroutine “done” lowers(-) the counter
- Wait() blocks until the counter hits zero
### Basic Steps
1. Create a WaitGroup.
2. Add the number of goroutines.
3. Run goroutines and call Done() when finished.
4. Wait in main until all are done

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // mark this worker as finished when the function ends
    time.Sleep(time.Second)
    fmt.Println("Worker", id, "done")
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1)               // step 1: increment counter
        go worker(i, &wg)       // step 2: start goroutine
    }

    wg.Wait()                   // step 3: block until all Done() calls happen
    fmt.Println("All workers finished!")
}

```

## Channels
- A goroutine is a lightweight, concurrent worker
- If several goroutines run at the same time, they often need to talk to each other
- Without a safe way to communicate, you’d have to share variables and worry about “race conditions.”
That’s where channels come in
- Think of a `channel` as a `pipe` that lets one goroutine send a value and another goroutine receive it.
```go
ch := make(chan int)     // channel that carries int
// You use make to create it, just like a slice or map.

// Send a value
ch <- 42    // send 42 into the channel


// Receive a value
v := <- ch  // take a value out of the channel


```

## Mutex
- Mutex stands for Mutual Exclusion
- It’s like a key to the whiteboard
- Only the person holding the key can write
- Everyone else must wait until the key is returned


## File
- In Go, whenever you want to open, read, or write files, you use the os package

### Package Functions (os)
| Function | Description |
|----------|-------------|
| `Create(name string) (*File, error)` | Creates a new file (truncates if exists). |
| `Open(name string) (*File, error)` | Opens a file for reading. |
| `OpenFile(name string, flag int, perm FileMode) (*File, error)` | Opens a file with specific flags and permissions. |
| `Mkdir(name string, perm FileMode) error` | Creates a single directory. |
| `MkdirAll(path string, perm FileMode) error` | Creates a directory path recursively. |
| `Remove(name string) error` | Deletes a file or empty directory. |
| `RemoveAll(path string) error` | Deletes a directory and all its contents. |
| `Rename(oldpath, newpath string) error` | Renames a file or directory. |
| `ReadDir(name string) ([]DirEntry, error)` | Reads directory entries. |
| `Stat(name string) (FileInfo, error)` | Returns file info. |
| `Lstat(name string) (FileInfo, error)` | Like Stat, but does not follow symlinks. |
| `Chmod(name string, mode FileMode) error` | Changes file permissions. |
| `Chown(name string, uid, gid int) error` | Changes file owner. |
| `Chtimes(name string, atime, mtime time.Time) error` | Changes access/modification time. |
| `Getwd() (dir string, err error)` | Returns current working directory. |
| `Chdir(dir string) error` | Changes working directory. |
| `Hostname() (name string, err error)` | Returns host machine name. |
| `UserHomeDir() (string, error)` | Returns user’s home directory. |
| `Getenv(key string) string` | Gets environment variable. |
| `Setenv(key, value string) error` | Sets environment variable. |
| `Unsetenv(key string) error` | Deletes environment variable. |
| `Environ() []string` | Returns all environment variables. |
| `TempDir() string` | Returns the default temp directory. |
| `CreateTemp(dir, pattern string) (*File, error)` | Creates a temp file. |
| `StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)` | Starts a new process. |
| `Getpid() int` | Returns current process ID. |
| `Getppid() int` | Returns parent process ID. |
| `FindProcess(pid int) (*Process, error)` | Finds a process by PID. |

---

#### Create
```go
func Create(name string) (*os.File, error)
```
- Creates a new file. Truncates if the file already exists.  
- **Parameters:** `name string` → path of new file.  
- **Returns:**  
  - `*os.File` → writable file handle.  
  - `error` → non-nil if creation fails.

```go
f, err := os.Create("newfile.txt")
if err != nil {
    log.Fatal(err)
}
defer f.Close()
```

#### Open
```go
func Open(name string) (*os.File, error)
```
- Opens a file for reading.  
- **Parameters:** `name string` → path to file.  
- **Returns:**  
  - `*os.File` → read-only file handle.  
  - `error` → non-nil if opening fails.

```go
f, err := os.Open("example.txt")
if err != nil {
    log.Fatal(err)
}
defer f.Close()
```

#### OpenFile
```go
func OpenFile(name string, flag int, perm os.FileMode) (*os.File, error)
```
- Opens a file with specific flags and permissions.  
- **Parameters:**  
  - `name string` → file path  
  - `flag int` → mode (e.g., `os.O_RDONLY`, `os.O_CREATE|os.O_WRONLY`)  
  - `perm os.FileMode` → file permissions  
- **Returns:**  
  - `*os.File` → file handle  
  - `error` → non-nil if opening fails

```go
f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
    log.Fatal(err)
}
defer f.Close()
f.WriteString("New log entry\n")
```

#### Mkdir
```go
func Mkdir(name string, perm os.FileMode) error
```
- Creates a single directory.  
- **Parameters:** `name string`, `perm os.FileMode`  
- **Returns:** `error` → non-nil if creation fails

```go
err := os.Mkdir("testdir", 0755)
if err != nil {
    log.Fatal(err)
}
```

#### MkdirAll
```go
func MkdirAll(path string, perm os.FileMode) error
```
- Creates a directory path recursively.  
- **Parameters:** `path string`, `perm os.FileMode`  
- **Returns:** `error` → non-nil if creation fails

```go
err := os.MkdirAll("parent/child/grandchild", 0755)
if err != nil {
    log.Fatal(err)
}
```

#### Remove
```go
func Remove(name string) error
```
- Deletes a file or empty directory.  
- **Parameters:** `name string`  
- **Returns:** `error` → non-nil if deletion fails

```go
err := os.Remove("file.txt")
if err != nil {
    log.Fatal(err)
}
```

#### RemoveAll
```go
func RemoveAll(path string) error
```
- Deletes a directory and all its contents.  
- **Parameters:** `path string`  
- **Returns:** `error` → non-nil if deletion fails

```go
err := os.RemoveAll("parent")
if err != nil {
    log.Fatal(err)
}
```

#### Rename
```go
func Rename(oldpath, newpath string) error
```
- Renames a file or directory.  
- **Parameters:** `oldpath string`, `newpath string`  
- **Returns:** `error` → non-nil if rename fails

```go
err := os.Rename("old.txt", "new.txt")
if err != nil {
    log.Fatal(err)
}
```

#### ReadDir
```go
func ReadDir(name string) ([]os.DirEntry, error)
```
- Reads directory entries.  
- **Parameters:** `name string`  
- **Returns:**  
  - `[]os.DirEntry` → slice of entries  
  - `error` → non-nil if reading fails

```go
entries, err := os.ReadDir(".")
if err != nil {
    log.Fatal(err)
}
for _, e := range entries {
    fmt.Println(e.Name())
}
```

#### Stat
```go
func Stat(name string) (os.FileInfo, error)
```
- Returns file info.  

#### Lstat
```go
func Lstat(name string) (os.FileInfo, error)
```
- Like Stat, but does not follow symlinks.  

#### Chmod
```go
func Chmod(name string, mode os.FileMode) error
```
- Changes file permissions.  

#### Chown
```go
func Chown(name string, uid, gid int) error
```
- Changes file owner.  

#### Chtimes
```go
func Chtimes(name string, atime, mtime time.Time) error
```
- Changes access and modification time.  

#### Getwd
```go
func Getwd() (dir string, err error)
```
- Returns current working directory.  

#### Chdir
```go
func Chdir(dir string) error
```
- Changes working directory.  

#### Hostname
```go
func Hostname() (name string, err error)
```
- Returns host machine name.  

#### UserHomeDir
```go
func UserHomeDir() (string, error)
```
- Returns user’s home directory.  

#### Getenv
```go
func Getenv(key string) string
```
- Gets environment variable.  

#### Setenv
```go
func Setenv(key, value string) error
```
- Sets environment variable.  

#### Unsetenv
```go
func Unsetenv(key string) error
```
- Deletes environment variable.  

#### Environ
```go
func Environ() []string
```
- Returns all environment variables.  

#### TempDir
```go
func TempDir() string
```
- Returns the default temp directory.  

#### CreateTemp
```go
func CreateTemp(dir, pattern string) (*os.File, error)
```
- Creates a temporary file.  

#### StartProcess
```go
func StartProcess(name string, argv []string, attr *os.ProcAttr) (*os.Process, error)
```
- Starts a new process.  

#### Getpid
```go
func Getpid() int
```
- Returns current process ID.  

#### Getppid
```go
func Getppid() int
```
- Returns parent process ID.  

#### FindProcess
```go
func FindProcess(pid int) (*os.Process, error)
```
- Finds a process by PID.  


# *os.File Type Methods

| Method | Description |
|--------|-------------|
| `Read(b []byte) (n int, err error)` | Reads bytes from the file. |
| `ReadAt(b []byte, off int64) (n int, err error)` | Reads bytes from a specific offset. |
| `ReadDir(n int) ([]DirEntry, error)` | Reads directory entries. |
| `ReadFrom(r io.Reader) (n int64, err error)` | Reads data from an io.Reader. |
| `Write(b []byte) (n int, err error)` | Writes bytes to the file. |
| `WriteAt(b []byte, off int64) (n int, err error)` | Writes bytes at a specific offset. |
| `WriteString(s string) (ret int, err error)` | Writes a string to the file. |
| `WriteTo(w io.Writer) (n int64, err error)` | Writes data to an io.Writer. |
| `Close() error` | Closes the file. |
| `Sync() error` | Flushes file contents to disk. |
| `Chmod(mode FileMode) error` | Changes file permissions. |
| `Chown(uid, gid int) error` | Changes file owner. |
| `Truncate(size int64) error` | Truncates the file to a specific size. |
| `Seek(offset int64, whence int) (ret int64, err error)` | Moves the file pointer. |
| `Name() string` | Returns the file name. |
| `Fd() uintptr` | Returns the file descriptor. |
| `Stat() (FileInfo, error)` | Returns file info. |

---

## Method Details

### Read
```go
func (f *os.File) Read(b []byte) (n int, err error)
```
- Reads bytes from the file.  
- **Parameters:** `b []byte` → buffer to store data.  
- **Returns:**  
  - `n int` → number of bytes read  
  - `err error` → error if reading fails

```go
buf := make([]byte, 100)
n, err := f.Read(buf)
if err != nil {
    log.Fatal(err)
}
```

### ReadAt
```go
func (f *os.File) ReadAt(b []byte, off int64) (n int, err error)
```
- Reads bytes from a specific offset in the file.  
- **Parameters:** `b []byte`, `off int64`  
- **Returns:** `n int`, `err error`

```go
n, err := f.ReadAt(buf, 50) // read 100 bytes from offset 50
if err != nil {
    log.Fatal(err)
}
```

### ReadDir
```go
func (f *os.File) ReadDir(n int) ([]os.DirEntry, error)
```
- Reads directory entries from an opened directory.  
- **Parameters:** `n int` → number of entries to read, -1 for all.  
- **Returns:** `[]os.DirEntry`, `error`

```go
entries, err := f.ReadDir(-1)
if err != nil {
    log.Fatal(err)
}
for _, e := range entries {
    fmt.Println(e.Name())
}
```

### ReadFrom
```go
func (f *os.File) ReadFrom(r io.Reader) (n int64, err error)
```
- Reads from an io.Reader and writes to the file.  
- **Parameters:** `r io.Reader`  
- **Returns:** `n int64`, `err error`

```go
n, err := f.ReadFrom(reader)
if err != nil {
    log.Fatal(err)
}
```

### Write
```go
func (f *os.File) Write(b []byte) (n int, err error)
```
- Writes bytes to the file.  
- **Parameters:** `b []byte`  
- **Returns:** `n int`, `err error`

```go
n, err := f.Write([]byte("Hello World"))
if err != nil {
    log.Fatal(err)
}
```

### WriteAt
```go
func (f *os.File) WriteAt(b []byte, off int64) (n int, err error)
```
- Writes bytes at a specific offset in the file.  
- **Parameters:** `b []byte`, `off int64`  
- **Returns:** `n int`, `err error`

```go
n, err := f.WriteAt([]byte("GoLang"), 10)
if err != nil {
    log.Fatal(err)
}
```

### WriteString
```go
func (f *os.File) WriteString(s string) (ret int, err error)
```
- Writes a string to the file.  

### WriteTo
```go
func (f *os.File) WriteTo(w io.Writer) (n int64, err error)
```
- Writes file contents to an io.Writer.  

### Close
```go
func (f *os.File) Close() error
```
- Closes the file.  

### Sync
```go
func (f *os.File) Sync() error
```
- Flushes file contents to disk.  

### Chmod
```go
func (f *os.File) Chmod(mode os.FileMode) error
```
- Changes file permissions.  

### Chown
```go
func (f *os.File) Chown(uid, gid int) error
```
- Changes file owner.  

### Truncate
```go
func (f *os.File) Truncate(size int64) error
```
- Truncates file to a specific size.  

### Seek
```go
func (f *os.File) Seek(offset int64, whence int) (ret int64, err error)
```
- Moves the file pointer.  

### Name
```go
func (f *os.File) Name() string
```
- Returns the file name.  

### Fd
```go
func (f *os.File) Fd() uintptr
```
- Returns the file descriptor.  

### Stat
```go
func (f *os.File) Stat() (os.FileInfo, error)
```
- Returns file info.  



# FileInfo Interface Methods

| Method | Description |
|--------|-------------|
| `Name() string` | Returns file or directory name. |
| `Size() int64` | Returns file size in bytes. |
| `Mode() FileMode` | Returns file permissions/mode. |
| `ModTime() time.Time` | Returns last modification time. |
| `IsDir() bool` | Returns true if it’s a directory. |
| `Sys() interface{}` | Returns OS-specific information. |

---

## Method Details

### Name
```go
func (fi FileInfo) Name() string
```
- Returns the name of the file or directory.  

### Size
```go
func (fi FileInfo) Size() int64
```
- Returns the size of the file in bytes.  

### Mode
```go
func (fi FileInfo) Mode() os.FileMode
```
- Returns the file permissions/mode.  

### ModTime
```go
func (fi FileInfo) ModTime() time.Time
```
- Returns the last modification time.  

### IsDir
```go
func (fi FileInfo) IsDir() bool
```
- Returns `true` if the file is a directory.  

### Sys
```go
func (fi FileInfo) Sys() interface{}
```
- Returns OS-specific information about the file.  





# Process Type Methods

| Method | Description |
|--------|-------------|
| `Release() error` | Releases resources associated with the process. |
| `Signal(sig Signal) error` | Sends a signal to the process. |
| `Kill() error` | Kills the process. |
| `Wait() (*ProcessState, error)` | Waits for the process to exit. |

---

## Method Details

### Release
```go
func (p *Process) Release() error
```
- Releases any resources associated with the process.  
- **Returns:** `error` → non-nil if releasing fails

### Signal
```go
func (p *Process) Signal(sig os.Signal) error
```
- Sends a signal to the process.  
- **Parameters:** `sig os.Signal` → signal to send  
- **Returns:** `error` → non-nil if sending fails

### Kill
```go
func (p *Process) Kill() error
```
- Kills the process.  
- **Returns:** `error` → non-nil if killing fails

### Wait
```go
func (p *Process) Wait() (*os.ProcessState, error)
```
- Waits for the process to exit.  
- **Returns:**  
  - `*os.ProcessState` → state information of the exited process  
  - `error` → non-nil if waiting fails

# Process Type Methods

| Method | Description |
|--------|-------------|
| `Release() error` | Releases resources associated with the process. |
| `Signal(sig Signal) error` | Sends a signal to the process. |
| `Kill() error` | Kills the process. |
| `Wait() (*ProcessState, error)` | Waits for the process to exit. |

---

## Method Details

### Release
```go
func (p *Process) Release() error
```
- Releases any resources associated with the process.  
- **Returns:** `error` → non-nil if releasing fails

### Signal
```go
func (p *Process) Signal(sig os.Signal) error
```
- Sends a signal to the process.  
- **Parameters:** `sig os.Signal` → signal to send  
- **Returns:** `error` → non-nil if sending fails

### Kill
```go
func (p *Process) Kill() error
```
- Kills the process.  
- **Returns:** `error` → non-nil if killing fails

### Wait
```go
func (p *Process) Wait() (*os.ProcessState, error)
```
- Waits for the process to exit.  
- **Returns:**  
  - `*os.ProcessState` → state information of the exited process  
  - `error` → non-nil if waiting fails


# ProcessState Type Methods

| Method | Description |
|--------|-------------|
| `Pid() int` | Returns PID of exited process. |
| `Exited() bool` | Returns true if the process exited. |
| `Success() bool` | Returns true if the exit code was 0. |
| `Sys() interface{}` | Returns OS-specific exit information. |
| `String() string` | Returns a human-readable description. |
| `ExitCode() int` | Returns the exit code of the process. |

---

## Method Details

### Pid
```go
func (ps *ProcessState) Pid() int
```
- Returns the PID of the exited process.  

### Exited
```go
func (ps *ProcessState) Exited() bool
```
- Returns `true` if the process has exited.  

### Success
```go
func (ps *ProcessState) Success() bool
```
- Returns `true` if the process exited with code 0.  

### Sys
```go
func (ps *ProcessState) Sys() interface{}
```
- Returns OS-specific exit information.  

### String
```go
func (ps *ProcessState) String() string
```
- Returns a human-readable description of the process state.  

### ExitCode
```go
func (ps *ProcessState) ExitCode() int
```
- Returns the exit code of the process.  






### Reading from a file
```go
buffer := make([]byte, 100)    // create a buffer of 100 bytes
dataReadedSize, err := f.Read(buf)       // read into buffer
if err != nil {
    log.Fatal(err)
}
fmt.Println("Readed data in Byte: ",dataReadedSize)
fmt.Println("Buffer in string ": ,string(dataReadedSize))

```

### Writing to a file
```go
f, err := os.Create("output.txt") // creates a new file (or overwrites if exists)
if err != nil {
    log.Fatal(err)
}
defer f.Close()

_, err = f.Write([]byte("Hello, Techie!")) 
if err != nil {
    log.Fatal(err)
}
```
- `f.Write` → writes `bytes` into it.






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