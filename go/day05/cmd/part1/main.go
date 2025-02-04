package main

import (
	"day05"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println(os.Getwd())
	}
	fmt.Println(day05.ProcessPart1(string(content)))
}
