package main

/*
Creating race conditions - Hugo Ernesto

We can say the code is under race conditions when it's outcome depends on the non-deterministic order of the interleaving created in the scheduling of the threads and the tasks in the code are communicating between each other.

In this exemple the code is under race conditions because both calls of the show the one in the main goroutine and the one in the second (created with the keyword go) are communicating through the variable x

In most runs (in my tests) the output will be the same: the prints of the call [1, 6, 11, 21, 31] happens so fast that the second call doesn't have time to print, but in other occasions, since the interlieaving is non-deterministic the values printed are altered due to the overlaping change of the value of x.
*/

import "fmt"

var x float32 = 1

func show(num *float32) {
	for i := 0; i < 5; i++ {
		fmt.Println(*num)
		add(num)
	}
}

func add(num *float32) {
	var z float32 = 1

	for i := 0; i < 5; i++ {
		*num = z + *num
	}
}

func main() {
	go show(&x)
	show(&x)
}
