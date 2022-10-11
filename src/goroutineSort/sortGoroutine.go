package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func getInput(seq *[]int) {

	fmt.Printf("Type at least four numbers separated by ',': ")
	userInput := bufio.NewScanner(os.Stdin)
	userInput.Scan()
	input := strings.Split(userInput.Text(), ",")

	for i := 0; i < len(input); i++ {
		num, err := strconv.Atoi(input[i])

		if err != nil {
			fmt.Println(err)
		}

		*seq = append(*seq, num)
	}

}

func sortSlice(ch chan []int) {
	sli := <-ch

	fmt.Println("received part: ", sli)
	sort.Ints(sli)

	fmt.Println("sorted received part: ", sli)
	ch <- sli
	wg.Done()
}

func main() {
	var seq []int

	ch := make(chan []int, 5)

	getInput(&seq)

	if len(seq) < 4 {
		fmt.Printf("Your input has less than four numbers\n please try again.")
		return
	}

	step := len(seq) / 4

	fmt.Println(seq)
	ch <- seq[0:step]
	ch <- seq[step : step*2]
	ch <- seq[step*2 : step*3]
	ch <- seq[step*3:]

	wg.Add(4)
	go sortSlice(ch)
	go sortSlice(ch)
	go sortSlice(ch)
	go sortSlice(ch)
	wg.Wait()

	s1 := <-ch
	s2 := <-ch
	s3 := <-ch
	s4 := <-ch

	sorted := []int{}
	sorted = append(sorted, s1...)
	sorted = append(sorted, s2...)
	sorted = append(sorted, s3...)
	sorted = append(sorted, s4...)

	sort.Ints(sorted)

	fmt.Println("The sorted sequence is: ", sorted)
}
