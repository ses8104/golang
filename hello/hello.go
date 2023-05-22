package main

import "fmt"

const PI = 3.14
const FloatPI float64 = 3.14

func main() {

	//#1
	//fmt.Println("Hello World")

	//#2
	// var a int = 10
	// var msg string = "Hello Variable"

	// msg1 := "Alvin"

	// a = 20
	// msg = "Good Moring"
	// fmt.Println(msg, msg1, a)

	//#3
	var a int = PI * 100
	var b float64 = FloatPI * 100

	fmt.Println(a, b)

}
