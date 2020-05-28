package main

import (
	"fmt"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("wd :", err)
	}
	fmt.Println("wd :", wd)
}
