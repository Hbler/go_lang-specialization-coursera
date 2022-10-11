package main

import (
	"fmt"
	"sync"
	// "math/rand"
	// "sync"
	// "time"
)

var wg sync.WaitGroup

type ChopStk struct{ sync.Mutex }

type Philo struct {
	rightCS, leftCS *ChopStk
	id              int
}

func (p Philo) eat(eating chan int, canEat chan bool) {

	eating <- p.id

	if <-canEat {
		for i := 0; i < 3; i++ {
			p.rightCS.Lock()
			p.leftCS.Lock()

			fmt.Println("starting to eat ", p.id)
			fmt.Println("finishing eating ", p.id)

			p.leftCS.Unlock()
			p.rightCS.Unlock()
		}
	}

	defer wg.Done()

}

// func randomN() int {

// 	rand.Seed(time.Now().UnixNano())
// 	min := 1
// 	max := 5

// 	return rand.Intn(max-min+1) + min

// }

func host(eating chan int, canEat chan bool) {

	var areEating []int

	philo := <-eating

	areEating = append(areEating, philo)

	canEat <- len(areEating) < 3

	defer wg.Done()

}

func main() {
	CStks := make([]*ChopStk, 5)

	for i := 0; i < 5; i++ {
		CStks[i] = new(ChopStk)
	}

	Philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		Philos[i] = &Philo{CStks[i], CStks[(i+1)%5], i + 1}
	}

	eating := make(chan int)
	canEat := make(chan bool)

	wg.Add(6)

	for i := 0; i < 5; i++ {
		go Philos[i].eat(eating, canEat)
		go host(eating, canEat)
	}
	wg.Wait()
}
