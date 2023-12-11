package ch05functions

import (
	"errors"
	"fmt"
	"strconv"
)

func add(i int, j int) (int, error) { return i + j, nil }
func sub(i int, j int) (int, error) { return i - j, nil }
func mul(i int, j int) (int, error) { return i * j, nil }
func div(i int, j int) (int, error) {
	if j == 0 {
		err := errors.New("divisor cannot be zero")
		return 0, err
	}
	return i / j, nil
}

var opMap = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func Ex1() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
		{"10", "/", "0"},
	}

	for _, exp := range expressions {
		if len(exp) != 3 {
			fmt.Println("invalid expression: ", exp)
			continue
		}
		o1, err := strconv.Atoi(exp[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		o2, err := strconv.Atoi(exp[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		opr := exp[1]

		f, ok := opMap[opr]
		if !ok {
			fmt.Println("Operator not supported", opr)
			continue
		}
		res, err := f(o1, o2)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(res)

	}
}
