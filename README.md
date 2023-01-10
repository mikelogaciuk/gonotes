# gonotes

![gonotes_header](./img/GOPHER_SHARE.png)

## About

`Go` or so called `Golang` is a programming language developed in Google since 2007.

Authors of a language wanted to combine best parts of Python and C++ and take an advantage of multi-core CPU's.

Go is easy to learn, multipurpose language, compiled and statically typed.

It has concise syntax which uses only 25 keywords.

Up to now (v1.19) those keywords are:

```xf
break     default      func    interface  select
case      defer        go      map        struct
chan      else         goto    package    switch
const     fallthrough  if      range      type
continue  for          import  return     var
```

## Starting a project

In order to start coding, we need to install `Go`, for this you can use a little googling.

Next, just init a module and start coding:

```shell
cd repos && mkdir example && go mod init example && touch main.go
```

## Standard libraries

All standard libraries (which you can import into your code), can be found [here](https://pkg.go.dev/std).

## CLI

Standard `Go` CLI can be accessed via shell:

```shell
go
```

The commands are:

    bug         start a bug report
    build       compile packages and dependencies
    clean       remove object files and cached files
    doc         show documentation for package or symbol
    env         print Go environment information
    fix         update packages to use new APIs
    fmt         gofmt (reformat) package sources
    generate    generate Go files by processing source
    get         add dependencies to current module and install them
    install     compile and install packages and dependencies
    list        list packages or modules
    mod         module maintenance
    work        workspace maintenance
    run         compile and run Go program
    test        test packages
    tool        run specified go tool
    version     print Go version
    vet         report likely mistakes in packages

For more you can type:

```shell
go help
```

## Entry point and main package

Main package is initiated with a phrase:

```go
package main
```

Main entry point of an code, similar to languages like Kotlin, C, Java or C# goes into:

```go
package main

func main() {
	
}
```

In order to quick testing, compile on fly, we can use:

```shell
go run main.go
```

As you can imagine, we can build it with a command like:

```shell
go build main.go
```

## Variables

For variables declaration we use `var` for variables and `consts` for constants (you should know what it means).

```shell
var foo string // Without initialization.
var bar string = "I am an initialized variable." 
quaz := "I am short hard declared variable." // Only inside functions.
```

## Data types

In `Go` we can find such data types as:

```go
package main

//String
var thisIsTheString string = "I am a string"


// Bool
var isItTrue bool = true
var isItRobot bool = false

// Numeric types (ripped from -> https://kps.hashnode.dev/learn-go-the-complete-course)
var i int = 404                         // Platform dependent
var i8 int8 = 127                       // -128 to 127
var i16 int16 = 32767                   // -2^15 to 2^15 - 1
var i32 int32 = -2147483647             // -2^31 to 2^31 - 1
var i64 int64 = 9223372036854775807     // -2^63 to 2^63 - 1
var ui uint = 404                       // Platform dependent
var ui8 uint8 = 255                     // 0 to 255
var ui16 uint16 = 65535                 // 0 to 2^16
var ui32 uint32 = 2147483647            // 0 to 2^32
var ui64 uint64 = 9223372036854775807   // 0 to 2^64
var uiptr uintptr                       // Integer representation of a memory address

// Floats
var fl32 float32 = 1.1333
var fl64 float64 = 4.5551

// Complex
var compl128 complex128 = 20 + 2 
var compl64 complex64 = 10 + 46 
```

In case of integers, it is recommended to use an `int` if there is no other reason to use a specific signed or unsigned integers.

There are also two comples types, where for complex128 both parts (real and imaginary) are float64 and complex64 where real and imaginary are float32.

There are also two additional integer types: the `byte` (`uint8`) and `rune` (`int32`), which are simple aliases...

```go
package main

type byte = uint8
type rune = int32
```

## Type conversion

We can convert data types as easy as:

```go
package main

var inty int = 42
var f = float64(inty)
```

## Zero values

Quite opposite to other languages like Python or Ruby, the variables without an explicit initial values are given `zero value`.

For numbers the returned number will be `0`, while for booleans it will be `false` and `""` for strings.

## Variable printing

In order to print variables, we can use standard library `fmt`:

```go
package main

import "fmt"

var x string = "Hello <3"

func main () {
	fmt.Println("x")
}
```

Additionally we can use `f-strings in Python` equivalent:

```go
package main

import "fmt"

var x string = "Mike"

func main () {
	fmt.Printf("Hello: %s", x)
}
```

As you can think, the `%s` is the annotation verb which tell how to format arguments.
And... as you can imagine there are few of them e.g.:

```xf
bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p
```

For more info you can visit [this](https://pkg.go.dev/fmt) place.

![Gopher at beach](./img/This_is_Fine_Gopher.png)

## The flow

### If/else

The `if/else` clause works quite similar to other languages.

The parentheses `()` around expression is not mandatory.

Sample if:

```go
package main

func main() {
	
}
```

Another sample if:

```go
package main

import "fmt"

func main() {
	var numbs int = 666
	
	if numbs <= 333 {
		fmt.Println("Number is too low...")
    } else if numbs > 334 && <= 665 {
		fmt.Println("Number is kinda close...")
    } else if numbs >= 667 {
		fmt.Println("Seek lower...")
    } else {
		fmt.Println("Welcome to HELL.")
    }
}
```

Another way is an availability to use short-hand if/else:

```go
package main

import "fmt"

func main() {
	if x:= 555; x > 444 {
		fmt.Println("Variable 'x' is greater than 444...")
    }
}
```

