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

func problemOne(instructions []string) {
	var instructionMap = map[string]int{
		"UD": 0, // up down
		"LR": 0, // left right
		"RD": 0, // right diagonal /
		"LD": 0, // left diagonal \
	}

	for _, instruction := range instructions {
		switch instruction {
		case "n":
			instructionMap["UD"]++
		case "s":
			instructionMap["UD"]--
		case "e":
			instructionMap["LR"]++
		case "w":
			instructionMap["LR"]--
		case "ne":
			instructionMap["RD"]++
		case "sw":
			instructionMap["RD"]--
		case "nw":
			instructionMap["LD"]++
		case "se":
			instructionMap["LD"]--
		}
	}

	fmt.Println(instructionMap)
}

func main() {
	instructions := getInput()
	problemOne(instructions)
}
