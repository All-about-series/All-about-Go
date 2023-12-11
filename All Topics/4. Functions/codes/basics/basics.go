package basics

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Run is the entry point for the basics module
func Run() {
	fmt.Println(strings.Repeat("-", 50), "Basics")

	fmt.Println(strings.Repeat("-", 50), "named and optional params using structs")
	// Named and optional params
	NamedAndOptsParams(MyFuncOpts{LastName: "Patel", Age: -50})
	NamedAndOptsParams(MyFuncOpts{LastName: "Shah"})

	fmt.Println(strings.Repeat("-", 50), "variadic inputs")
	// Variadic input params and slices
	fmt.Println(Variadic(5, 1, 2, 3, 4, 5))

	fmt.Println(strings.Repeat("-", 50), "funcs are values")
	// Functions as values
	result, _, err := MultiReturnValue(3, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result, err)

	fmt.Println(strings.Repeat("-", 50), "Named return values")
	x, y, z := NamedReturnValue(3, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x, y, z)

	fmt.Println(strings.Repeat("-", 50), "Blank returns values")
	x, y, z = BlankReturns(4, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x, y, z)

	fmt.Println(strings.Repeat("-", 50), "func decl can be changed")
	//func decl can be changed
	res := add(2, 3)
	fmt.Println(res)
	changeAdd()
	res1 := add(2, 3)
	fmt.Println(res1)

	fmt.Println(strings.Repeat("-", 50), "funcs returning funcs")
	//funcs returning funcs
	fmt.Println(makeMult(2)(9))
	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 1; i < 4; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}

	fmt.Println(strings.Repeat("-", 50), "clean code example")
	// Expression evaluation eg of clean code
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := OpMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(result)
	}

	fmt.Println(strings.Repeat("-", 50), "pass by value nature of go")
	//go is pass by value
	person := MyFuncOpts{FirstName: "aadarsh"}
	i := 2
	name := "aadarsh"
	passByValue(i, name, person)
	fmt.Println(i, name, person) //values are not changed

	//maps, slice are also pass by value, but their implementation have pointers inside, so a copy of address is passed
	mp := map[int]string{1: "atul", 2:"shah"}
	defaultPassByRef(mp) //value changes
	fmt.Println(mp)

}

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

// named and optional params can be emutated with structs
func NamedAndOptsParams(opts MyFuncOpts) error {
	if opts.Age < 0 {
		return errors.New("invalid age")
	}
	fmt.Print("Hello")
	if opts.FirstName != "" {
		fmt.Print(" ", opts.FirstName)
	}
	if opts.LastName != "" {
		fmt.Print(" ", opts.LastName)
	}
	if opts.Age != 0 {
		fmt.Print(", you are ", opts.Age, " years old")
	}
	fmt.Println(".")
	return nil
}

// variadic input params and slices
func Variadic(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

// multiple return vales
func MultiReturnValue(num int, den int) (int, int, error) {
	if den == 0 {
		return 0, 0, errors.New("cannot divide by 0")
	}
	return num / den, num % den, nil
}

func NamedReturnValue(num int, den int) (result int, rem int, _ error) { // error is nameless with _
	if den == 0 {
		return result, rem, errors.New("cannot divide by 0") //named return values are init to their zero values when created
	}
	result, rem = num/den, num%den
	return result, rem, nil //returned names are local to the fn
}

func BlankReturns(num int, den int) (result int, rem int, err error) {
	if den == 0 {
		err = errors.New("cannot divide by 0")
		return
	}
	result, rem = num/den, num%den
	return //returns the values assigned to the named returns
}

var add = func(i int, j int) int { return i + j }

func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }

var OpMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func changeAdd() {
	add = func(i, j int) int { return i + j + j }
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func passByValue(i int, name string, person MyFuncOpts) {
	i = 19
	name = "atul"
	person.FirstName = "atul"
}

func defaultPassByRef(mp map[int]string) {
	mp[1] = "aadarsh"
}
