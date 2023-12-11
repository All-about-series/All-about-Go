package main

import (
	ch03 "LearningGo/Ch03_CompositeTypes"
	ch05functions "LearningGo/Ch05-Functions"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Repeat("-", 50), "Ch03_CompositeTypes")
	ch03.Sol()
	fmt.Println(strings.Repeat("-", 50), "ch05-functions")
	ch05functions.Ex1()

	//steps to run:
	//go build main.go
	//./main <fileName> (eg. main.go)
	ch05functions.Ex2()
}
