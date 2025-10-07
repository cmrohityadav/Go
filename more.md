| Type             | Zero-value      | Nil?  | make() effect                             |
| ---------------- | --------------- | ----- | ----------------------------------------- |
| int, float, bool | 0, 0.0, false   | ❌ no  | N/A                                       |
| string           | "" (empty)      | ❌ no  | N/A                                       |
| pointer `*T`     | nil             | ✅ yes | `new(T)` → alloc pointer, non-nil         |
| slice `[]T`      | nil             | ✅ yes | `make([]T,n)` → non-nil, len=n, cap≥n     |
| map `map[K]V`    | nil             | ✅ yes | `make(map[K]V)` → non-nil, ready-to-use   |
| channel `chan T` | nil             | ✅ yes | `make(chan T, n)` → non-nil, ready-to-use |
| interface        | nil             | ✅ yes | N/A                                       |
| struct           | all fields zero | ❌ no  | N/A                                       |




---

VALUE TYPE (int, string, bool)
+--------+
| x = 0  |  <- zero-value
+--------+
Stack memory

POINTER / REFERENCE TYPE (slice, map, channel)
Nil slice / map / channel
Stack header:
+---------------------+
| pointer = nil       |  <- not allocated
| length=0 / cap=0    |
+---------------------+
Heap: none allocated

Make slice/map/channel
Stack header:
+---------------------+
| pointer ------------> Heap allocated
| length >=0          |
| capacity >=0        |
+---------------------+
Heap (slice example):
+---+---+---+---+---+
| 0 | 0 | 0 | 0 | 0 |
+---+---+---+---+---+



---


# File
### Package Functions (os)
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


### *os.File Type Methods


#### Fd
```go
func (f *os.File) Fd() uintptr
```
- Returns the file descriptor.  

#### Sync
```go
func (f *os.File) Sync() error
```
- Flushes file contents to disk.  

#### Chmod
```go
func (f *os.File) Chmod(mode os.FileMode) error
```
- Changes file permissions.  

#### Chown
```go
func (f *os.File) Chown(uid, gid int) error
```
- Changes file owner.  

#### Truncate
```go
func (f *os.File) Truncate(size int64) error
```
- Truncates file to a specific size.  

#### Seek
```go
func (f *os.File) Seek(offset int64, whence int) (ret int64, err error)
```
- Moves the file pointer.  




### Process Type Methods

| Method | Description |
|--------|-------------|
| `Release() error` | Releases resources associated with the process. |
| `Signal(sig Signal) error` | Sends a signal to the process. |
| `Kill() error` | Kills the process. |
| `Wait() (*ProcessState, error)` | Waits for the process to exit. |

---

### Method Details

#### Release
```go
func (p *Process) Release() error
```
- Releases any resources associated with the process.  
- **Returns:** `error` → non-nil if releasing fails

#### Signal
```go
func (p *Process) Signal(sig os.Signal) error
```
- Sends a signal to the process.  
- **Parameters:** `sig os.Signal` → signal to send  
- **Returns:** `error` → non-nil if sending fails

#### Kill
```go
func (p *Process) Kill() error
```
- Kills the process.  
- **Returns:** `error` → non-nil if killing fails

#### Wait
```go
func (p *Process) Wait() (*os.ProcessState, error)
```
- Waits for the process to exit.  
- **Returns:**  
  - `*os.ProcessState` → state information of the exited process  
  - `error` → non-nil if waiting fails





### ProcessState Type Methods

| Method | Description |
|--------|-------------|
| `Pid() int` | Returns PID of exited process. |
| `Exited() bool` | Returns true if the process exited. |
| `Success() bool` | Returns true if the exit code was 0. |
| `Sys() interface{}` | Returns OS-specific exit information. |
| `String() string` | Returns a human-readable description. |
| `ExitCode() int` | Returns the exit code of the process. |

---

### Method Details

#### Pid
```go
func (ps *ProcessState) Pid() int
```
- Returns the PID of the exited process.  

#### Exited
```go
func (ps *ProcessState) Exited() bool
```
- Returns `true` if the process has exited.  

#### Success
```go
func (ps *ProcessState) Success() bool
```
- Returns `true` if the process exited with code 0.  

#### Sys
```go
func (ps *ProcessState) Sys() interface{}
```
- Returns OS-specific exit information.  

#### String
```go
func (ps *ProcessState) String() string
```
- Returns a human-readable description of the process state.  

#### ExitCode
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


