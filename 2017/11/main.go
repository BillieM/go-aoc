package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getInput() []string {
	var instructions []string

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var str string

	for scanner.Scan() {
		str = scanner.Text()
	}

	instructions = strings.Split(str, ",")

	return instructions
}

func problemOne(instructions []string) int {

	var vert float64
	var hori float64

	var steps []string

	for _, instruction := range instructions {
		switch instruction {
		case "n":
			vert++
		case "s":
			vert--
		case "ne":
			vert += 0.5
			hori++
		case "nw":
			vert += 0.5
			hori--
		case "se":
			vert -= 0.5
			hori++
		case "sw":
			vert -= 0.5
			hori--
		}
	}

	for vert != 0 {
		for hori != 0 {
			if hori > 0 {
				if vert > 0 {
					vert -= 0.5
					steps = append(steps, "ne")
				} else {
					vert += 0.5
					steps = append(steps, "se")
				}
				hori--
			} else {
				if vert > 0 {
					vert -= 0.5
					steps = append(steps, "nw")
				} else {
					vert += 0.5
					steps = append(steps, "sw")
				}
				hori++
			}
		}

		if vert > 0 {
			vert--
			steps = append(steps, "n")
		} else if vert < 0 {
			vert++
			steps = append(steps, "s")
		}
	}

	return len(steps)
}

func problemTwo(instructions []string) int {
	// loop through all possible slices and calc num steps
	maxSteps := 0

	for i := 0; i < len(instructions); i++ {
		slice := instructions[:i+1]
		sliceSteps := problemOne(slice)
		if sliceSteps > maxSteps {
			maxSteps = sliceSteps
		}
	}

	return maxSteps
}

func main() {
	instructions := getInput()
	// fmt.Println(problemOne(instructions))
	fmt.Println(problemTwo(instructions))
}
