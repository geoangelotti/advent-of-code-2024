package main

import (
	"day07"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println(os.Getwd())
	}
	fmt.Println(day07.ProcessPart2(string(content)))
}
