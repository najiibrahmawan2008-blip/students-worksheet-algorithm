package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	if x * x * x == y { // multiply x by itself three times
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
