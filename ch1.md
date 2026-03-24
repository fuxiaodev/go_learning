# Why We Write Go

## Speed Comparison

### Execution Speed (Interpretative vs Compiled)

**Interpreted (slower execution):**
- JavaScript
- Python
- Ruby
- PHP

**Compiled (faster execution):**
- C
- C++
- Rust
- Go

**Virtual Machine (compiled to bytecode, then run on VM — typically slower):**
- Java
- C#

## Why Go?

Go is a compiled language, and the compiled machine code runs directly on the local machine.

However, it is slightly slower than low-level compiled languages like **Rust** and **C** because of the **Go runtime**, which handles:
- Memory management
- Garbage collection
- Goroutine scheduling

Even so, Go typically uses **less memory than Java and C#**, since it does not require a full virtual machine.

> [!TIP]
> Go code generally runs faster than interpreted languages and compiles faster than other compiled languages

## Memory Footprint
Go programs are fairly lightweight. Each program includes a small amount of extra code that is included in the executable binary called the go runtime. The go runtime cleans up unused memory at runtime. It includes a garbage collector that automatically frees up memory that's no longer in use.

As a general rule, Java uses more memory than Go mainly because Java uses a virtual machine to interpret bytecode at runtime and typically allocates more on the heap.


## Run Go

### Running a go script
```
go run <filename>.go
```
- go compiles your code into a temporary binary
- runs that binary immediately
- deletes the temporary binary after execution

quick for testing but slower than running a prebuilt executable

### Running a compiled executable
```
go build <filename>.go
./<filename>
```
need to recompile after each code change, but fast execution for each run.

## Compilation process
- ```package main``` lets the go compiler know that we want this code to compile and run as a standalone program, as opposed to being a library that's imported by other programs
- ```func main``` defines the main function, the entry point for a go program

### Two kinds of Errors
- Compilation Errors: Occur when code is compiled. The resulting executable won't be created
- Runtime Errors: Occur when a program is running. These errors can cause your program to crash or behave unexpectedly.


# Variables

## Signed integers(no decimal)
```int int8 int16 int32 int64```

## Unsigned integers(non-negative numbers/no decimal)
```uint unit8 uint16 uint32 uint64 uintptr```

## Signed decimal numbers
```float32 float64```

## Complex numbers(a complex number has a real and imaginary part)
```complex64 complex128```

The size 8, 16, 32, 64, 128 represents the number of bits in memory would be allocated to store the variable
> [!TIP]
> The default int and uint types refer to their respective 32 or 64 bit sizes depending on the environment of the user

The standard sizes should be used unless you have a specific performance need are
- bool
- string
- int
- uint
- byte
- rune
- float64
- complex128

## Go is statically typed
Go enforces static typing meaning variable types are known before the code runs.
This mean your editor and the compiler can display type errors before the code is ever run, making development easier and faster
Python and Javascript are dynamically typed.

## Declaration preference
Prefer := inside functions
Prefer var with inference at package level
Specify types only when needed
