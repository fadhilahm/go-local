package main

import "fmt"

func main() {
	data := [...]int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	for index := 0; index < len(data); index++ {
		switch data[index] {
		case 5:
			fmt.Printf("this is a test: 5")
		case 7:
			fmt.Printf("this is a test: 7")
		case 9:
			fmt.Printf("this is a test: 9")
		default:
		}
		switch {
		case data[index] < 10: // <<<<< var,19,3,19,6,newVar,fail
			fmt.Printf("this is a test: 9")
		default:
		}
	}
}
