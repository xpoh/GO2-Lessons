package main

import (
	"fmt"
)

func main() {
	c, err := ReadConfig()
	if err != nil {
		fmt.Println(err)
	}
	err = SliceFile(*c)
	if err != nil {
		fmt.Println("Error slicing ", err)
	}
}
