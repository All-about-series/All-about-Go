package ch03

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func Sol() {
	//ex 1
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	fmt.Println(greetings[:1])
	fmt.Println(greetings[1:4])
	fmt.Println(greetings[3:])

	//ex 2
	message := "Hi 👧 and 👦"
	runesFromString := []rune(message)
	fmt.Println(runesFromString, "\t", len(runesFromString))
	// fmt.Println(string(runes[3]))
	fmt.Println(string(runesFromString))

	//wrong approach (this way we are not fetching the correct code point of 👧)
	runesFromSliceOfBytes := []byte(message)
	fmt.Println(runesFromSliceOfBytes, "\t", len(runesFromSliceOfBytes)) //notice the smiley takes up multiple character bytes
	runeFromByte := rune(message[3])                                     //since message[3] is a byte (we are getting the correct code point of smiley)
	fmt.Println(string(runeFromByte))                                    //wrong

	//ex 3
	e1 := Employee{"aadarsh", "shah", 1}
	e2 := Employee{firstName: "aadarsh", lastName: "shah", id: 1}
	var e3 Employee
	e3.firstName = "aadarsh"
	e3.lastName = "shah"
	e3.id = 1

	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e3)

}
