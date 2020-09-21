package main

import "fmt"

func main() {
	var a[10] int

	fmt.Println("Array a: ", a)

	a[1] = 5 // modify second element

	fmt.Println("Array a: ", a)

	fmt.Println(a[2])
}
