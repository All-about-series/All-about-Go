# Composite Types

1. [Arrays](#arrays)
2. [Slices](#slices)
3. [Maps](#maps)
4. [Structs](#structs)

## Arrays-

> `NOTE:` dont use it unless exact length is known ahead of time (ahead of compilation/at runtime)

```go
var x [3]int
var y = [2]int{3,4}
var z = [12]int{1, 5: 4, 6, 10: 100, 15} //sparse array(an arr where most elements are set to their zero value)
//Output- [1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15]
var a [2][3]int //declares x to be of len 2, whose internal type is of arr of ints of len 3. Go doesn't have a true matrix support
var b = [...]int{3,4} //same as y
```

Arrays in Go are rarely used because they have these limitaion-

- `[3]int` and `[4]int` are of different tyeps.
- Cannot even do _type converion_ to make two arrays the same size/type
- ofcorse then we cannot use same var to store differnt array sizes

```go
//converting array to slice
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[:]

//converting sub array to slice
slice := arr[1:3]
```

## Slices-

- Length of the `slice` is not part of the type, unlike `array`
- Dynamic in nature, with following internal implementation
  > `cap()` and `len()` fn:  
  > Slice has a `capacity` (`cap()` fn), the no. of consecutive
  > memory locations reserved. This can be larger than the `length` (`len()` fn). Each value added to `slice` increases its `length` by one. When the `length` == `capacity`, there’s no more room to put values, any new `append`s now uses the Go runtime to allocate a new backing `array` for the `slice` with a larger `capacity`. The values from the original backing `array` are copied to the new one, the new values are added to the end of the new backing `array`, and
  > the `slice` is updated to refer to the new backing `array`. Finally, the updated `slice` is returned.
- Zero-value can be `nil`(means lack of value, _eg. its **!=** untyped numeric constants and hence **cannot** be compared_)
  ```go
  var nilSlice []int
  var nonNilSlice = []int{}
  fmt.Println(nilSlice == nil)    //true
  fmt.Println(nonNilSlice == nil) //false
  ```
  For simplicity, favor `nil slices`. A `zero-length slice` is useful only when converting a `slice` to `JSON`
- error when slices are compared, only thing it can be compared with is `nil`. Use `slices.Equal()` fn instead

- `[...]` makes an array, `[]` makes a slice
- ```go
  var x = []int{1, 5: 4, 6, 10: 3, 18}
  var x [][]int
  ```

  `make` fn in Go, helps to create an empty slice (`non-nil`) and can be used to specify the `len` and `cap`

  ```go
  x := make([]int, 3) //len of 3
  x = append(x,10)//will make x = [0,0,0,10] (len of 4)
  y := make([]int, 0, 10) //len of 0, cap of 10
  y = append(y ,3,4,5) //will make y = [3,4,5] (as the len initially was 0)

  clear(y)//all values are set to its zero value, not changing the len and cap
  ```

  There are 3 possibilites for an idiomatic use of `make` for slices-

  1.  If slice are used as a buffer (as in “io and Friends”), then specify a `non-zero len`
  2.  If exact size is _sure_/_known_, then specify the `len` and index into slices to set values. Requires careful insertion.
  3.  In other situations, specify a `0` len with specified `cap`, allowing `append` to be used easily. `Advantage`- never will have any zero values left to fill at the end

  ### Slicing Slices

  Taking a slice from another slice, the subslice’s capacity is set to the capacity of the original slice, minus the starting offset of the subslice within the original slice.

  ```go
  x := make([]string, 0, 5)
  y := x[:2]
  z := x[2:]

  fmt.Println(cap(x)) //5
  fmt.Println(cap(y)) //5, which is 5-0 = 5
  fmt.Println(cap(z)) //3, which is original slice cap minus starting offset of subslice (i.e. 5-2 = 3)
  ```

  > `NOTE`- Be careful when taking a `slice of a slice` OR `slices of arrays` (not to be confused with `arrays of slice`- they are ok to work with) in both case slices share the same memory from its parent (either slice or arr), and changes to one are reflected in the other, use _full slice_ expressions to prevent `appends` from sharing capacity between slices
  >
  > `NOTE`- Dont use append with subslices or make sure you are using _full slice_ expressions
  >
  > eg. look at below code example

  ```go
  //Should not be used this way (as slice of subslices are used)

  x2:= make([]string, 0, 5)
  x2 = append(x2, "a", "b", "c", "d")
  y2 := x2[:2]
  z2 := x2[2:]
  fmt.Println(cap(x2), cap(y2), cap(z2))   //cap(z2) is 3, which is original slice cap  minus starting offset of subslice (i.e. 5-2 = 3)
  y2 = append(y2, "i", "j", "k") //orginal slice is changed along with new len
  // y2 = append(y2, "i", "j", "k", "l", "m") //orginal slice is NOT changed as the size of the slice is increased beyond the cap
  fmt.Println(len(x2), cap(x2), len(y2), cap(y2), len(z2), cap(z2))
  x2 = append(x2, "x")
  z2 = append(z2, "y")
  fmt.Println("x2:", x2)
  fmt.Println("y2:", y2)
  fmt.Println("z2:", z2)
  ```

  ```go
  //since full slice expressions are used, below code is perfectly fine to use
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
  ```

  ### `copy` fn

  Creates a new memory for new slice/array copy

  ### Slice to array

  Creates a new memory for the newly created array

## Maps-

- Zero-value of `map` is `nil` .

  ```go
  var x map[string]int //nil map
  y := map[string]int{} //empty map (non-nil)

  //write to a nil map will cause a panic
  x["key"] = 10 //panic, as x is nil
  y["key"] = 10 //works, as y is non-nil

  z := make(map[string]int) //empty map (non-nil)
  fmt.Println(z == nil) //false
  ```

- Maps are not comparable with `==` and `!=`, only thing it can be compared with is `nil`

  - But using `maps` package in standard lib, we can compare maps

  ```go
  m1 := map[string]int{"one": 1, "two": 2}
  m2 := map[string]int{"two": 2, "one": 1}

  fmt.Println(maps.Equal(m1, m2)) //true
  ```

- The key for a `map` can be any comparable type. This means `key` cannot be a `slice`, `map` and `func`(as they are not comparable).

- **Checking for a key**-

  ```go
  totalSum := map[string]int{
    "apples": 10,
  }

  //key exists or not
  v, ok := totalSum["apples"] //v = 10, ok = true
  v, ok = totalSum["mangoes"] //v = 0, ok = false

  totalSum["mangoes"]++
  fmt.Println(totalSum) //map[apples:10 mangoes:1]
  ```

- A cleared map has its `length` set to `0`, unlike a cleared slice. (using `clear()`)
- In Go, `map` is used as a set. This is because `map` is the only collection type that can be used as a set in Go.

  ```go
  set := map[int]bool{} //disadvantage of bool is that it takes 1 byte of memory when empty
  fmt.Println(set[1]) //false

  preferredSet := map[int]struct{}{} //advantage of struct{} is that it takes 0 bytes of memory when empty
  if _, ok := preferredSet[1]; ok {
    fmt.Println("1 is present in set")
  }else{
    fmt.Println("1 is not present in set")
  }
  ```

## Structs

To define related data together

Go doesn't have `classes` because it doesn't have `inheritance` (inheritence has many disadvantages), instead with struct Go prefers `composition`.

```go
type person struct {
		name      string
		age       int
		isInIndia bool
}

//Unlike maps, there is no difference between assigning an empty struct literal and not assigning a value at all (like below eg p1 and p2). Both initialize all fields in the struct to their zero values.
var p1 person
p2 := person{}
fmt.Println(p1 == p2) //true
```

- Anonymous `struct` are also possible
- Structs that are entirely composed of comparable types are comparable; those with `slice`, `map`, `func` or `chan` fields are not comparable.
- Go doesn’t allow comparisons between variables that represent structs of different types, just like different primitive types are not comparablable.

  But they can be type converted if they have same field `name`, `order` and `types`.
