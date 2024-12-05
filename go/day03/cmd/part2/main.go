package main

import (
	"day03"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println(os.Getwd())
	}
	fmt.Println(day03.ProcessPart2(string(content)))
}
