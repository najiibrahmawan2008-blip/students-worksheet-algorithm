package main

import "fmt"

func main() {
	first, second := 5, 3

	
	third := first + second
	fourth := second + third
	fifth := third + fourth

	fmt.Println(third, fourth, fifth)
}
