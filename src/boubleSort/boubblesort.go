package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Type at least two numbers separated by ',': ")
	userInput := bufio.NewScanner(os.Stdin)
	userInput.Scan()
	input := strings.Split(userInput.Text(), ",")
	sequence := convertToNum(input)
	BoubbleSort(sequence)

	fmt.Printf("%v", sequence)

}

func convertToNum(sli []string) []int {

	var seq []int

	for i := 0; i < len(sli); i++ {
		num, err := strconv.Atoi(sli[i])

		if err != nil {
			fmt.Println(err)
		}

		seq = append(seq, num)
	}

	return seq
}

func BoubbleSort(sli []int) {
	goOn := true
	for goOn {
		goOn = false
		for i := 0; i < len(sli)-1; i++ {

			if sli[i] > sli[i+1] {
				Swap(sli, i)
				goOn = true
			}

		}

	}
}

func Swap(sli []int, index int) {
	higher := sli[index]
	lower := sli[index+1]

	sli[index] = lower
	sli[index+1] = higher
}

// func BubbleSort(sli []int) []int {
// 	for j := len(sli) - 1; j > 0; j-- {
// 		for i := 0; i < j; i++ {
// 			if sli[i] > sli[i+1] {
// 				Swap(sli, i)
// 			}
// 		}
// 	}

// 	return sli
// }

// func Swap(sli []int, i int) {
// 	temp := sli[i]
// 	sli[i] = sli[i+1]
// 	sli[i+1] = temp
// }
