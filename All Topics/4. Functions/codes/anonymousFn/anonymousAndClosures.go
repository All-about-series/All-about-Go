package anonymousfn

import (
	"fmt"
	"sort"
	"strings"
)

func Run() {
	//anonymous fns are also called lambda functions in other progarmming terminologies
	fmt.Println(strings.Repeat("-", 50), "closures") //closure are special type of anonymous fns
	//anonymous funcs and closures
	func() {
		a := 3
		func() {
			a = 3 //accessible because of closure
			a += 2
			fmt.Println(a)
			a := 3 //shadowing(this is not closure)
			a++
			fmt.Println(a)
		}()
		a++
		fmt.Println(a)
	}()

	func() {
		b := 1
		f := func() int {
			b += 10
			return b
		}
		fmt.Print("b: ")
		fmt.Println(f()) //passing closure value of another func to use it outside the funcs
	}()

	type Person struct {
		firstName string
		lastName  string
		age       int
	}

	person := []Person{{"a", "A", 19}, {"b", "B", 17}, {"c", "C", 10}}
	sort.Slice(person, func(i, j int) bool { //person is captured by the closure

		return person[i].age < person[j].age //sorting by age
	})

	fmt.Println(person)

}
