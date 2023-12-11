Combined [func code examples](../../codes/basics/basics.go)
# Functions

- `func` are values, can be anonymous
- Go doesn't have: `named and optional input params` for `func`.
  - Define a `struct` to emutate a `named and optional input params` (not having this feature is not a limitation, as a function
    shouldn’t have more than a few params, and named and optional params are mostly useful when a function has many inputs)
- have named returned values and the names are only local to the `func` at it should be
- can return blank values, but one should never (it is against readablility convention), in this case the most current values are used

## Variadic Input Params and Slices

Must be the only or the last param in the input params list, `...` is nothing but representing a slice internally

## Go is call by value
Go always make a `copy` of the value of the variable when passed through a function.
- Types in Go are pass-by-value
- composite types (like maps and slice) are implemented with pointers hence the origial values gets changed, if a value inside the func is modified 

<hr/>

1. [Anonymous func and Closure](../2.%20anonymousFunc/anonymousAndClosure.md)
  
