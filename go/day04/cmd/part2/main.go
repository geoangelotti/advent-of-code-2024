package main

import (
	"day04"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println(os.Getwd())
	}
	fmt.Println(day04.ProcessPart2(string(content)))
}
