package main

import "fmt"

func main(){
	slice1 := [] int {2, 3, 4, 6, 8, 9, 3, 4}
	nonSlice1 := slice1[1:3]
	slice2 := make([] int, len(slice1))
	fmt.Println(slice1)
	fmt.Println(nonSlice1)
	fmt.Println(slice2)
}