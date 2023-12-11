
## Modules and packages
*example projcet structure* -
```go
myproject/
├── go.mod             # Module declaration
├── main.go            # Package: main- Entry point of program
├── utils/             # Package: utils
│   ├── math.go
│   └── string.go
├── models/            # Package: models
│   ├── user.go
│   └── product.go
└── handlers/          # Package: handlers
    ├── user.go
    └── product.go
```
its `go.mod` file
```go
module example.com/myproject

go 1.20
```
- Modules can have multiple packages, provide a high-level scope for dependency management and versioning.
- Packages (*a reusable single unit of code*) lives in their own directories within a module
- `main` package- executes after all the packages are imported and all `init()`s of every pacakge is executed, `main()` fn executes first in the main pacakge. An example of order of execution below-

**Order of execution**-   
*example folder structure* -
```go
myproject/
├── main.go          # Main package
├── utils/
│   └── utils.go     # Utils package
└── db/
    └── db.go        # Database package
```
`main.go`-
```go
package main

import (
    "fmt"
    "example.com/myproject/utils"
    "example.com/myproject/db"
)

func main() {
    fmt.Println("Executing main...")
    utils.DoSomething()
    db.Connect()
}
```

`utils/utils.go`-
```go
package utils

import "fmt"

func init() { //executes before main fn
    fmt.Println("Initializing utils package...")
}

func DoSomething() {
    fmt.Println("Doing something in utils...")
}
```

`db/db.go`-
```go
package db

import "fmt"

func init() { //executes before main fn
    fmt.Println("Initializing db package...")
}

func Connect() {
    fmt.Println("Connecting to database...")
}
```
```md
# Output of Program
Initializing utils package...
Initializing db package...
Executing main...
Doing something in utils...
Connecting to database...
```
> **Points to note out**- 
>- `init() Functions`: Executed automatically before `main()`.
>Used for setup tasks like initializing variables, reading configs, or setting up connections. It cannot be called explicitly.
>- Import order in the `main` package determines the order of initialization for packages. If a package imports other packages, their initialization is done recursively.
>- Circular Dependencies: Go does not allow circular dependencies between packages, ensuring a clear and predictable order of execution.

Exampe of `circular` dependency-  
`packageA/a.go`
```go
package packageA
import (
	"example.com/myproject/packageB"
	"fmt"
)
func FuncA() {
	fmt.Println("Function A in package A")
	packageB.FuncB()
}
```
`packageB/b.go`

```go
package packageB

import (
 	"example.com/myproject/packageA"
 	"fmt"
)

func FuncB() {
 	fmt.Println("Function B in package B")
 	packageA.FuncA()
}
 ```
As `packageA` is dependent on `pacakgeB` and vice-versa, Go compiler throws an error `import cycle not allowed` 

## Some Go Tools 

0. **`go mod`**-
   - `go mod init [module]`- initiates a new Go module in the current dir, creates a `go.mod` file, which keeps track of the module's dependency
   - `go mod tidy`- ensures `go.mod` and `go.sum` files are up-to-date
     - Adds any missing module dependencies required by the project
	 - Removes any unused module dependencies from `go.mod` and `go.sum`
1. **`go build`**-
   - `go build` - compiles code
   - `go build -o hello` - changes the output binary filename to  **hello**
2. **`go run`**- compiles and run

3. **`go clean`**-
    - `go clean`- removes object and cached files
	- `go clean -modcache`- removes caches files in the module cache
	- `go clean -testcache`- remoes test binaries	

4. **`go fmt`**-     
	>`Russ Cox`, the development lead for Go, has publicly stated that better tooling was his original motivation.

	Go language enforcing a standard format
	makes it a great deal easier to write tools that manipulate source code. _This simplifies
	the compiler and allows the creation of some clever tools for generating code.

	**Secondary benefit**- Now developer cannot argue about making their formatting tool better, as developers have historically wasted extraordinary amounts of time on format wars. lol

	```bash
	go fmt ./...
	```

	`./...` tells a Go tool to apply the command to all the files in the current
	directory and all subdirectories

5. **`go vet`**-
	```go
	fmt.Printf("Hello, %s!\n") //wrong code, 'go vet' will output error
	```
	One of the thing that `go vet` detects-   whether a value exists for every placeholders in a formatting template

<hr/>

### Semicolon insertion Rule

The Go compiler adds `;` automatically, following a simple rule described in [Effective Go](https://go.dev/doc/effective_go#semicolons).
If the last token before a newline is either of-

- An identifier (which includes words like int and float64)
- A basic literal such as a number or string constant
- One of the tokens: break, continue, fallthrough, return, ++, --, ), or }

the lexer inserts a semicolon after the token

```go
func main() //;
{
    fmt.Println("Hello,world!") //;
} //;
```
**eg**- the above code is invalid as, after `)` and `}` the compiler inserts `;` and `func main();` is an invalid code
<hr/>

#### Makefile
[Great free makefile tutorial](https://makefiletutorial.com/)

- `.DEFAULT_GOAL`- defines which target is run when no target is specified
- `.PHONY` line keeps `make` from getting confused if a directory or file in your project has the same name as one
of the listed targets.
- `<a>:<b>`- target `b` runs before target `a`
- running the Makefile- `make vet`, `make fmt` and `make build`(or simply `make` for our case)

