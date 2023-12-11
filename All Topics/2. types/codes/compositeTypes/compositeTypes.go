package compositetypes

import (
	"fmt"
	"maps"
	"strings"
)

func CompositeTypes() {
	var nilSlice []int
	var nonNilSlice = []int{}
	fmt.Println(nilSlice == nil)    //true
	fmt.Println(nonNilSlice == nil) //false

	var y = [2]int{3, 4}
	var z = [12]int{1, 5: 4, 6, 10: 100, 15} //sparse array
	var a = [...]int{3, 4}
	fmt.Println(y == a) //true
	fmt.Println(z)
	// y = [3]int{3, 6, 3} //error, as array size are part of type, so [3]int and [2]int are different types and cannot be reassigned
	var x = []int{1, 5: 4, 6, 10: 3, 18, 19: 34}
	var y1 = []int{1, 5: 4, 6, 10: 3, 18, 19: 34}
	// fmt.Println(x1 == y1) //error, slices cannot be compared
	fmt.Println(x == nil) //false
	fmt.Println(cap(y1))  //false

	fmt.Println(strings.Repeat("-", 50))
	// len_cap()
	fmt.Println(strings.Repeat("-", 50))
	// sliceOps()
	fmt.Println(strings.Repeat("-", 50))
	// sliceOfStrings()
	fmt.Println(strings.Repeat("-", 50))
	// mapOps()
	fmt.Println(strings.Repeat("-", 50))
	structOps()

}

// structs
func structOps() {
	//structs can be defined inside of functions as well
	type person struct {
		name      string
		age       int
		isInIndia bool
	}

	//Unlike maps, there is no difference between assigning an empty struct literal and not assigning a value at all (like below eg p1 and p2). Both initialize all fields in the struct to their zero values.
	var p1 person
	p2 := person{}
	fmt.Println(p1 == p2) //true

	p3 := person{
		"John",
		30,
		false,
	}
	p2 = p3
	p3 = person{ //no neccessary to give provide key when using this representation
		age:       24,
		isInIndia: true,
	}
	fmt.Println(p3, p2)
	p3.name = "Aadarsh"
	fmt.Println(p3.name)

	//anonymous structs
	fmt.Println(strings.Repeat("-", 50), "anonymous structs")
	var animal struct {
		name     string
		age      int
		haveTail bool
	}
	animal.name = "Dog"
	animal.age = 5
	animal.haveTail = true
	fmt.Println(animal)

	reptile := struct {
		name            string
		age             int
		onlyLandReptile bool
	}{
		name:            "Crocodile",
		age:             10,
		onlyLandReptile: false,
	}
	fmt.Println(reptile)

	//structs type conversion
	fmt.Println(strings.Repeat("-", 50), "structs type conversion")
	type firstPerson struct {
		name string
		age  int
	}

	// firstPerson and secondPerson instances can be type converted as they have same fields and types, but they are of differnt types hence cannot be compared with `==`
	type secondPerson struct {
		name string
		age  int
	}

	//firstPerson and thirdPerson instances cannot be type converted as the order of fields is different
	type thirdPerson struct {
		age  int
		name string
	}

	//firstPerson and fourthPerson instances cannot be type converted as the fields are different
	type fourthPerson struct {
		FirstName string
		age       int
	}

	//firstPerson and fifthPerson instances cannot be type converted as fifthPerson has an additional field
	type fifthPerson struct {
		name      string
		age       int
		isInIndia bool
	}

	fmt.Println(firstPerson(secondPerson{name: "aadarsh", age: 24}))
	// fmt.Println(firstPerson(thirdPerson{})) //error
	fmt.Println(thirdPerson{age: 24, name: "aadarsh"})
	fmt.Println(fourthPerson{FirstName: "aadarsh", age: 24})
	fmt.Println(fifthPerson{name: "aadarsh", age: 24, isInIndia: true})

	//if atleast one of operand is anonymous structs it can be compared if they have the same field names, order and types
	fmt.Println(strings.Repeat("-", 50), "anonymous structs are comparable")
	// fmt.Println(animal == reptile) //name, order or type is different, hence cannot be compared
	var guy struct {
		name      string
		age       int
		isInIndia bool
	}

	//assignable and comparable as `guy` is an anonymous struct
	guy = p3
	fmt.Println(guy == p3)
}

// maps
func mapOps() {
	fmt.Println(strings.Repeat("-", 50), "map")
	var x map[string]int  //nil
	y := map[string]int{} //non-nil empty map
	fmt.Println(x == nil) //true
	// x["a"] = 1 //panic, as map is nil
	y["a"] = 1
	fmt.Println(x, len(x), y, len(y))

	z := make(map[string]int, 10) //len is still 0, can grow past the specified size
	fmt.Println(z, len(z))
	z["a"] = 1
	fmt.Println(z, len(z))

	fmt.Println(strings.Repeat("-", 50), "key check")
	//key check
	totalSum := map[string]int{
		"apples": 10,
	}
	v, ok := totalSum["apples"] //v = 10, ok = true
	fmt.Println(v, ok)
	v, ok = totalSum["mangoes"] //v = 0, ok = false
	fmt.Println(v, ok)

	totalSum["mangoes"]++
	fmt.Println(totalSum) //map[apples:10 mangoes:1]

	fmt.Println(strings.Repeat("-", 50), "delete and clear map")
	//delete from map
	delete(totalSum, "apples")
	fmt.Println(totalSum)

	//unlike slices, clear() makes makes the len as 0 for maps
	clear(totalSum)
	fmt.Println(totalSum)

	fmt.Println(strings.Repeat("-", 50), "compare maps using `maps` package in standard lib")
	//compare maps using standard library
	m1 := map[string]int{"one": 1, "two": 2}
	m2 := map[string]int{"two": 2, "one": 1}

	fmt.Println(maps.Equal(m1, m2)) //true

	//maps as set
	fmt.Println(strings.Repeat("-", 50), "maps as set")
	set := map[int]bool{} //disadvantage of bool is that it takes 1 byte of memory when empty
	vals := []int{1, 2, 1, 4, 2, 6, 7, 4, 9}
	for _, v := range vals {
		set[v] = true
	}
	fmt.Println(len(vals), len(set))

	preferredSet := map[int]struct{}{} //advantage of struct{} is that it takes 0 bytes of memory when empty
	for _, v := range vals {
		preferredSet[v] = struct{}{}
	}
	if _, ok := preferredSet[1]; ok {
		fmt.Println("1 is present")
	} else {
		fmt.Println("1 is not present")
	}
	fmt.Println(len(vals), len(preferredSet))
}

// sliceOps in strings
func sliceOfStrings() {
	// reference behaviors can be seen in slices of strings
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	x[1] = "y"
	y[0] = "x"
	z[1] = "z"
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	fmt.Print(strings.Repeat("-", 50), "\n")
	x1 := []string{"a", "b", "c", "d"}
	y1 := x1[:2]
	fmt.Println(len(x1), cap(x1), len(y1), cap(y1)) //only len is changed, cap is same as x1, since append to a non-nil slice will append from the len index, so the original slice is changed from that index
	y1 = append(y1, "z")
	fmt.Println("x1:", x1)
	fmt.Println("y1:", y1)

	//append should not be used with subslices
	fmt.Print(strings.Repeat("-", 50), "\n")
	x2 := make([]string, 0, 5)
	x2 = append(x2, "a", "b", "c", "d")
	y2 := x2[:2]
	z2 := x2[2:]
	fmt.Println(cap(x2), cap(y2), cap(z2)) //cap(z2) is 3, which is original slice cap  minus starting offset of subslice (i.e. 5-2 = 3)
	y2 = append(y2, "i", "j", "k")         //orginal slice is changed along with new len
	// y2 = append(y2, "i", "j", "k", "l", "m") //orginal slice is NOT changed as the size of the slice is increased beyond the cap
	fmt.Println(len(x2), cap(x2), len(y2), cap(y2), len(z2), cap(z2))
	x2 = append(x2, "x")
	// z2 = append(z2, "y")
	fmt.Println("x2:", x2)
	fmt.Println("y2:", y2)
	fmt.Println("z2:", z2)

	//or instead make sure you are using full slice expressions to avoid bugs
	//the full slice expression protects against append
	//this code is fine now
	fmt.Print(strings.Repeat("-", 50), "\n")
	x3 := make([]string, 0, 5)
	x3 = append(x3, "a", "b", "c", "d")
	y3 := x3[:2:2] //after 2nd : , the capacity is manually specified
	z3 := x3[2:4:4]
	fmt.Println(len(x3), cap(x3), len(y3), cap(y3), len(z3), cap(z3))
	y3 = append(y3, "i", "j", "k")
	x3 = append(x3, "x")
	z3 = append(z3, "y")
	fmt.Println("x3:", x3)
	fmt.Println("y3:", y3)
	fmt.Println("z3:", z3)
}

func sliceOps() {
	fmt.Println(strings.Repeat("-", 50), "make slice")
	x := make([]int, 5, 10)
	x = append(x, 1, 2, 2, 4, 5, 6) //append starts from 5th index
	fmt.Println(x, len(x), cap(x))  //initial cap is changed

	y := make([]int, 0, 5)
	y = append(y, 1, 2, 2, 4, 5, 6) //append starts from 0th index
	fmt.Println(y, len(y), cap(y))  //initial cap is changed

	clear(x)
	fmt.Println(x, len(x), cap(x)) //all values are set to its zero value, not changing the len and cap

	fmt.Println(strings.Repeat("-", 50), "copy slice")
	//copy slice
	x1 := []int{1, 2, 3, 4, 5}
	y1 := make([]int, 2)
	copy(y1, x1[2:4])       //no need to return len if not reqd.
	num := copy(y1, x1[1:]) //previous values are overwritten
	fmt.Println(x1, y1, num)
	num = copy(x1[:5], x1[2:]) //previous values are overwritten
	fmt.Println(x1, num)

	fmt.Println(strings.Repeat("-", 50), "copy array")
	//copy array (new mem location is created)
	x2 := []int{1, 2, 3, 4}
	d2 := [4]int{5, 6, 7, 8}
	y2 := make([]int, 2)
	copy(y2, d2[:])
	fmt.Println(y2)
	copy(d2[:], x2)
	fmt.Println(d2)

	//slices to array (new mem location is created)
	fmt.Println(strings.Repeat("-", 50), "slices to array")
	xSlice := []int{1, 2, 3, 4}
	xArray := [4]int(xSlice)
	smallArray := [2]int(xSlice)
	xArrayPointer := (*[4]int)(xSlice) //type converted to `poninter to an array`, changes in either will reflect in other
	// rdArray := [...]int(xSlice) //error, as the size of the array is not specified
	// panicArray := [5]int(xSlice) //Panis, as the size of the array is greater than the slice
	xSlice[0] = 10

	fmt.Println(xSlice)
	fmt.Println(xArray)
	fmt.Println(smallArray)
	fmt.Println(xArrayPointer)
}

// len and cap
func len_cap() {
	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))
}
