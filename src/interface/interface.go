package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Name() string
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name, food, locomotion, noise string
}

func (c Cow) Name() string {
	return c.name
}

func (c Cow) Eat() {
	fmt.Println(c.food + "\n")
}
func (c Cow) Move() {
	fmt.Println(c.locomotion + "\n")
}
func (c Cow) Speak() {
	fmt.Println(c.noise + "\n")
}

type Bird struct {
	name, food, locomotion, noise string
}

func (b Bird) Name() string {
	return b.name
}

func (b Bird) Eat() {
	fmt.Println(b.food + "\n")
}
func (b Bird) Move() {
	fmt.Println(b.locomotion + "\n")
}
func (b Bird) Speak() {
	fmt.Println(b.noise + "\n")
}

type Snake struct {
	name, food, locomotion, noise string
}

func (s Snake) Name() string {
	return s.name
}

func (s Snake) Eat() {
	fmt.Println(s.food + "\n")
}
func (s Snake) Move() {
	fmt.Println(s.locomotion + "\n")
}
func (s Snake) Speak() {
	fmt.Println(s.noise + "\n")
}

func createAnimal(name, animal string) Animal {
	err := "I couldn't understand, please verify your input and try again!"

	switch {
	case animal == "cow":
		return Cow{name, "grass", "walk", "moo"}
	case animal == "bird":
		return Bird{name, "worms", "fly", "peep"}
	case animal == "snake":
		return Snake{name, "mice", "slither", "hsss"}
	}

	return Cow{name, err, err, err}
}

func containsType(slice []string, t string) bool {
	for _, s := range slice {
		if t == s {
			return true
		}
	}
	return false
}

type Result struct {
	includes bool
	found    Animal
}

func containsName(slice []Animal, n string) Result {

	for _, a := range slice {
		if n == a.Name() {
			return Result{true, a}
		}
	}
	return Result{false, Cow{}}
}

func main() {
	var command, name, option string
	var animals []Animal

	types := []string{"cow", "bird", "snake"}
	created := "Animal created: "

	c := 0

	for {
		fmt.Println("\nEnter 'newanimal name type' to create an animal, or 'x' to stop.")
		fmt.Println("The name can be whatever you want, the options of animal type are: cow, bird, snake")
		if c > 0 {
			fmt.Println("Or enter 'query name action' to get a result, or 'x' to stop.")
			fmt.Println("The name must be of a created animal, the options of action are: eat, move, speak")
		}

		fmt.Scanln(&command, &name, &option)

		switch {
		case strings.ToLower(strings.ToLower(command)) == "x":
			return
		case strings.ToLower(strings.ToLower(command)) == "newanimal":
			if containsType(types, option) {
				anAnimal := createAnimal(strings.ToLower(name), strings.ToLower(option))
				animals = append(animals, anAnimal)
				fmt.Println(created, name, "the "+option)
			} else {
				fmt.Println("Oops! Wrong type. The options of animal type are: cow, bird, snake")
			}
		case strings.ToLower(strings.ToLower(command)) == "query":
			search := containsName(animals, strings.ToLower(name))
			if search.includes {
				switch {
				case strings.ToLower(option) == "eat":
					search.found.Eat()
				case strings.ToLower(option) == "move":
					search.found.Move()
				case strings.ToLower(option) == "speak":
					search.found.Speak()
				default:
					fmt.Println("Oops! Wrong action. The options of action type are: eat, move, speak")
				}
			} else {
				fmt.Println("Oops! Name not found. Try again!")
			}
		default:
			fmt.Println("I couldn't understand, please verify your command and try again!")
		}

		c++
	}
}
