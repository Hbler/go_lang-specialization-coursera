package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Move() string {
	return a.locomotion
}

func (a Animal) Talk() string {
	return a.noise
}

func createAnimal(animal string) Animal {
	err := "I couldn't understand, please verify your input and try again!"

	switch {
	case animal == "cow":
		return Animal{"grass", "walk", "moo"}
	case animal == "bird":
		return Animal{"worms", "fly", "peep"}
	case animal == "snake":
		return Animal{"mice", "slither", "hsss"}
	}

	return Animal{err, err, err}
}

func main() {
	var animal, action string

	for {
		fmt.Println("\nEnter an animal and action separated by a space to get a result, or 'x' to stop.")
		fmt.Println("The options of animal are: cow, bird, snake")
		fmt.Println("The options of action are: eat, move, talk")
		fmt.Scanln(&animal, &action)

		switch {
		case strings.ToLower(animal) == "x":
			return
		default:
			anAnimal := createAnimal(animal)
			switch {
			case strings.ToLower(action) == "eat":
				fmt.Println(anAnimal.Eat() + "\n")

			case strings.ToLower(action) == "move":
				fmt.Println(anAnimal.Move() + "\n")

			case strings.ToLower(action) == "talk":
				fmt.Println(anAnimal.Talk() + "\n")

			default:
				fmt.Println("I couldn't understand, please verify your input and try again!")
			}

		}
	}

}
