package main

import (
	"fmt"
	"sort"
)

func main() {
	SliceLoop()
}

func SliceLoop() {
	usrSlice := make([]int, 3)

	i := 0

	for {
		var input int

		fmt.Println("Enter an int to proceed or 'x' to quit:")
		n, err := fmt.Scan(&input)

		switch {
		case n != 1 || err != nil:
			return
		case i < 2:
			usrSlice[i] = input
		case i >= 2:
			usrSlice = append(usrSlice, input)
		}

		sort.Ints(usrSlice)
		fmt.Println(usrSlice)

		i++
	}
}
