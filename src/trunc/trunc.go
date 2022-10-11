package main

import "fmt"

func main() {
	getFloat_printTrunc()
}

func getFloat_printTrunc() {
	var userNum float32

	fmt.Println("Type a floating point number:")
	fmt.Scan(&userNum)
	fmt.Printf("The truncated number is: %.0f\n", userNum)

}
