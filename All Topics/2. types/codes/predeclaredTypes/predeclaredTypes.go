package predeclaredtypes

import (
	"fmt"
	"strings"
)

func PredeclaredTypes() {
	var a int = 10_100
	var b float64 = 6.03e34
	var c float32 = 0x12.34p5
	var d rune = 'a'          //single unicode character
	var e rune = '\u0122'     //16 bit hexadecimal no.
	var f rune = '\x11'       //8 bit hexadecimal no.
	var g rune = '\U00101233' //32 bit unicode no.
	var h rune = '\122'       //8 bit octal no.

	//explicit type conversion
	var i int = int(c)
	var j float64 = float64(c) // float32 to float64
	fmt.Println(a, b, c, d, e, f, g, h, i, j)

	fmt.Println(strings.Repeat("-", 50), "ways to mutate strings")
	//mutabe a string
	s := "hello"
	// s[0] = 'H' //error, strings are immutable
	bs := []byte(s) //convert to byte slice
	bs[0] = 'H'
	s = string(bs) //convert back to string
	fmt.Println(s)

	fmt.Println(strings.Repeat("-", 50), "mutli-byte characters")
	//code point of multi-byte characters
	var ss string = "Hello 🙂" //code point of smiley is not 1 byte long
	fmt.Println(ss[6:])
	fmt.Println(ss[4:7])
	fmt.Println(len(ss))

	fmt.Println(strings.Repeat("-", 50), "string to bytes and runes conversion")
	//string to bytes and runes conversion
	var s1 string = "Hello 🙂"
	var b1 []byte = []byte(s1)
	var r1 []rune = []rune(s1)
	fmt.Println(b1)
	fmt.Println(r1)
}
