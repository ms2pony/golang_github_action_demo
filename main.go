package main

import "fmt"

func Cat() string {
	return "Wang~~~"
}

func main() {
	saying := Cat()
	fmt.Println(saying)
}
