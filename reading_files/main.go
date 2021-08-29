package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Args", os.Args)

	if len(os.Args) != 2 {
		fmt.Println("Please specify the path")
		return
	}

	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("b ->", b)
}
