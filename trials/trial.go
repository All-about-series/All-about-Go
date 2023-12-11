package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("file is not provided")
		return
	}
	count, err := fileLen(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}

func fileLen(file string) (int, error) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return 0, nil
	}
	defer f.Close()
	data := make([]byte, 2048)
	total := 0
	for {
		count, err := f.Read(data)
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			break
		}
		total += count
	}
	return total, err
}
