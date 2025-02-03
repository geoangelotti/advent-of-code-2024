package main

import (
	"day06"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println(os.Getwd())
	}
	fmt.Println(day06.ProcessPart2(string(content)))
}
