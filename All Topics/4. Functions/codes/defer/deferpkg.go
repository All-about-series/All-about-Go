package deferpkg

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Run() {
	//defer keyword
	fmt.Println(strings.Repeat("-", 50), "defer")
	//build and run-
	//go build main.go
	//./main.go main.go
	if len(os.Args) < 2 { // []string {<nameOfProgram> <validFile>}
		log.Fatal("no file specified") //print err and exit program
	}
	f, closer, err := getFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	//not using the closer() will be a compile-time error, this convention allows the user to always use `defer`
	defer closer() //defer delays the invocation until the surrounding function exits.(in our case, when getFile() returns)
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		os.Stdout.Write(data[:count])
	}

	fmt.Println(strings.Repeat("-", 50), "multiple defer(LIFO order)")
	//defer only gets executed when the parent func returns, in our case when `a` is returned
	//LIFO order of execution
	func() int {
		a := 10
		defer func(val int) {
			fmt.Println("first:", val)
		}(a)
		a = 20
		defer func(val int) {
			fmt.Println("second:", val)
		}(a)
		a = 30
		fmt.Println("exiting:", a)
		return a
	}()

	//no way to read the returned values from defer
	func() {
		defer func() int {
			return 2 //no way to read this value
		}()
	}()
}

// A common pattern in Go is for a function that allocates a resource to also return a closure that cleans up the resource.
// this way the returned Func() must be used by the user, making sure the cleanup happens
func getFile(name string) (*os.File, func(), error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	return f, func() {
		f.Close()
	}, err
}
