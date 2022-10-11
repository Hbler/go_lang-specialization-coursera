package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {

	var (
		time     float64
		acc      float64
		initVel  float64
		initDisp float64
	)

	fmt.Println("Type a value for aceleration, initial velocity and initial displacement separated by a space:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	setupi := scanner.Text()
	setupr := strings.NewReader(setupi)

	_, setupErr := fmt.Fscanf(setupr, "%f %f %f", &acc, &initVel, &initDisp)
	if setupErr != nil {
		fmt.Fprintf(os.Stderr, "Fscanf: %v\n", setupErr)
	}

	displacementFn := GenDisplaceFn(acc, initVel, initDisp)

	fmt.Println("Type a value for the time in seconds:")
	scanner.Scan()
	timei := scanner.Text()
	timer := strings.NewReader(timei)

	_, err := fmt.Fscanf(timer, "%f", &time)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
	}

	fmt.Printf("The displacement in %.2f seconds is %.2f", time, displacementFn(time))

}

func GenDisplaceFn(a, v, d float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (0.5 * a * math.Pow(t, 2)) + (v * t) + d
	}

	return fn
}
