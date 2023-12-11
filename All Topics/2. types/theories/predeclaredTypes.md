# Predeclared Types
**Zero Value** - Go compiler assigns a default `zero value`, like most other langauges

## Literals-
*are **untyped**. eg-*
```go
var x float64 = 10 //assigning integer literal to floats are valid
var y float64 = 200.3 * 5 // mixing floats and int in literals are valid

//assigning string literal to a numeric type is NOT ALLOWED
//assigning float literal to an int type is NOT ALLOWED
// etc
```

1. `Integer` literals- are base 10 by default, but
   different prefixes are used to indicate other bases: `0b` for binary (base 2), `0o` for octal (base 8), or `0x` for hexadecimal (base 16) .
   Readability can be improved with `_`. eg- - numbers (base 10): `21_000` (group by thousands) - octal number (base 8): `234_345_437` - hexadecimal number (base 16): `F0A_567_9AB_CDEF`

2. `Floating` literals- decimal point to indicate a fractional portion.
   Can also have exponents specified with `e`.
   eg- `6.03e34`, `0x12.34p5`(in hexadecimal, `p` here is exponent) = `582.5`(in base10)

3. `rune` literal- represents character and is surrounded by single quotes, eg- `'a'`(single unicode characters), `'\141'`(8 bit octal no.), `'\x61'`(8-bit hexadecimal no.), `'\u0061'`(16-bit hexadecimal no.), `'\U00000062'`(32-bit unicode no.).
   Some Important ones are: `'\n'`(newline), `'\t'`(tab), `'\''`(single quote) and `'\\'`(backslash)
   Single quote and double quote in Go is different.

4. `string` literal-
   > `NOTE`: According to the language specification, Go source code is always written in `UTF-8`, unless you use hexadecimal escapes in a string literal.

   Two ways to use string literal-

   - `interpreted string literal`(contains zero or more `rune` literal): they are called interpreted as they interpret rune literal. The only chars that cannot appear are unescaped backslashes, unescaped newline and unescaped double quotes, hence in these cases, for ease of use, use `raw string literal`

     - eg. `"You will be a \"hero\""` (single quote escape `rune` literal is not valid in `string` literal, it is replaced by backslash escape for double quotes).

   - `raw string literal`: can contain any char except `\`(escape).
     - eg-
       ```go
       a := `You will bring the next
       "Revolution"`
       ```

5. `imaginary` literal *(a special type of numeric literal used in complex types)*: `numeric literal` with a suffix `i`. eg- `3i`, `3.2i`
   > Go is not a popular language for numerical computing, third-party `Gonum` package is better to use over standard lib, but its much better to choose other language first for numerical based computings.
   >
   >  Adoption has been limited for Go because other features (like matrix support) are not part of the language and libraries have to use inefficient
replacements, like slices of slices.
## Boolean types -
   `true` and `flase` are the literals in `bool` type

## Numeric types-

Go has a large number of numeric types: **12 types** (and a few special names).
Grouped into `3` categories-

1.  `Integer types` (zero-value is 0)-
     `int8, int16, int32, int64, uint8, uint16, uint32, uint64` (default integer type is `int`)

    The `special integer` types-
    - `byte`: an alias for `uint8`. It is allowed to assign, compare, or perform math operations b/w `byte` and `uint8` as it is platform indepedant.

    - `int`(platform dependant): as it is platform dependant hence, it is a `compile-time error` to assign, compare or perform math operations b/w `int` and `int32` unlike `byte`
      - on a `32-bit` CPU it is `int32`
      - on a `64-bit` CPU it is `int64` (Some uncommon 64-bit CPU architectures use a 32-bit signed integer for the int type. Go supports three of them: amd64p32, mips64p32, and mips64p32le) - `uint` - `rune` - `uintptr`

2.  `Floating-point types`(zero-value is 0)- `float32, float64` (default is `float64`)-
   All the math comparison opr works with floats except `%`.
   `0 / 0` is a `NaN`(not a number)
   `n / 0` is `+Inf` or `-Inf` based on sign of `n`
    > **NOTE**- `==` and `!=` should be used to compare floats as floats they are inexact in nature. Or use a maximum allowed variance (called `epsilon`- its value depends on the accuracy needs) and see if the difference b/w two floats is less than that

    > **NOTE**- Just like other languages, Go floating-point numbers
    have a huge range, but they cannot store every value in that range; they store the
    nearest approximation. Because floats aren’t exact, they can be used only in situations
    where inexact values are acceptable or the rules of floating point are well understood.
    That limits them to things like graphics, statistics, and scientific operations.
    >
    > To handle `exact decimal values` use third-party libraires. eg- `shopspring/decimal` is most commonly used for financial and business apps due to its precision.
3.  `Complex types`-
   `complex64` uses `float32` and `complex128` uses `float64` values to represent the real and imaginary parts.
   `complex128` is the default type.
    If one value of fn param is `float32` and the other value is an untyped constant or literal that can fit within a `float32`, then default type would be `complex64`
    It has the same precision limitation because of the floats

## Rune and String types-
String in Go is `immutable`, values can be reassigned to a string var, but cannot be changed.
`strings` are made of read only `[]bytes`, but is not its alias.
- `[]bytes` are mutable but not `strings`
   -  ```go
      s := "hello" //immutable
      s[0] = 'H' //error

      b := []byte("hello") //mutable
      b[0] = 'H'//error

      //convert string to []bytes then back to string to mutate the string
      bs := []byte(s)
      bs[0] = 'H'
      s = string(bs) //Hello (works)
      ```
- `rune` is an alias for `int32`
- use `rune` type for chars, not `int32` even though they are literally same. It improves Readability and avoid confusion
- slicing of `strings` has a problem-
  - code point that are of same byte long are ok to slice and work with, it will give consistent output
      > code point is the technical name given for each character and modifier
  - code points that are multiple bytes long, creates problem
    - ```go
      var s string = "Hello 🙂" //code point of smiley is not 1 byte long
      fmt.Println(s[6:]) //🙂
      fmt.Println(s[4:7]) //o � (as, sliced string has the code point that are multiple bytes long)
      fmt.Println(len(s)) //10 (not 7, because the smiley is 4 byte long in UTF-8 representation)
      ```
 - As of go 1.15, `go vet` blocks a type convertion to `string` from any other integer types other than `rune` and `byte`
   ```go
   var s string = "Hello, 🙂"
   var bs []byte = []byte(s) //[72 101 108 108 111 32 240 159 153 130] (UTF-8 bytes)
   var rs []rune = []rune(s) //[72 101 108 108 111 32 128578] (runes)
   ```
-  `strings` and `unicode/utf8` pacakges in standard lib should be used instead of direct slice and index expressions
## Explicit type conversion
   Some langauges allow implicit conversion between same type, but Go doesnot, it is strict and explicit type conversion is necessary.
   eg-
   ```go
   var a float32 = 54.3
   var b float64 = a //error- as floats are of different types here
   var b float64 = float64(a) // correct
   ```
   *this feature helps to reduce the bugs caused by coercion of numberic types to bool types, avoiding bugs*
