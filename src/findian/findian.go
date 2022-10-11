package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	findIan()
}

func findIan() {

	fmt.Printf("Type a string: ")
	userStr := bufio.NewScanner(os.Stdin)
	userStr.Scan()
	lowerStr := strings.ToLower(userStr.Text())

	if strings.HasPrefix(lowerStr, "i") && strings.HasSuffix(lowerStr, "n") && strings.Contains(lowerStr, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
