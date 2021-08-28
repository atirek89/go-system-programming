// Getting and setting the working directory
package main

import (
	"fmt"
	"os"
)

func main() {
	wd, err := os.Getwd()

	if err != nil {
		fmt.Println("Error Getwd: ", err)
		return
	}

	fmt.Println("Current working dir:", wd)

	if err := os.Chdir("/"); err != nil {
		fmt.Println("Error Chdir: ", err)
		return
	}

	if wd, err = os.Getwd(); err != nil {
		fmt.Println("Error Getwd: ", err)
		return
	}

	fmt.Println("Final working dir:", wd)
}
