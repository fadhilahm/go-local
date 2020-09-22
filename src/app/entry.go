// package main

// import (
// 	"fmt"
// 	// "greet/de"
// 	_ "greet"
// 	child "greet/greet"
// )

// var integers [10]int

// // var (
// // 	a int = b
// // 	b int = f()
// // 	c int = 1
// // )

// // func f() int {
// // 	return c + 1
// // }

// func init() {
// 	fmt.Println("app/entry.go ===> init() [1]")
// }

// func init() {
// 	fmt.Println("app/entry.go ==> init() [2]")
// }

// func init() {
// 	fmt.Println("app/entry.go ==> init()")

// 	for i:= 0; i < 10; i++ {
// 		integers[i] = i
// 	}
// }

// func main() {
// 	// version := "1.0.1"
// 	// fmt.Println("version ==> ", version)
// 	// var a int = b
// 	// var b int = c
// 	// var c int = 2

// 	// fmt.Println(a, b, c)

// 	fmt.Println("app/entry.go ===> main()")
// 	// fmt.Println(integers)

// 	// fmt.Println(greet.Message)
// 	fmt.Println(child.Message)
// }

package main

import "fmt"

func init() {
	fmt.Println("app/entry.go ==> init()")
}

var myVersion = fetchVersion()

func main() {
	fmt.Println("app/fetch-version.go ==> fetchVersion.go")
	fmt.Println("version ==> ", myVersion)
}