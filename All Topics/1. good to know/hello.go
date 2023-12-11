package main //main is an "Executable" type package

// Types of Packages
// 1. Executable (these packages generates a file that we can run
// 2. Reusable (helper code/library/dependency)

//NOTE- Go imports only whole packages. You can’t limit the import to specific types,functions, constants, or variables within a package.

import "fmt" //including the standard "fmt" package, into our "main" package

func main() { //every package "main" must contain "main" as the function
	fmt.Println("Hello, 世界")
}
