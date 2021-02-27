package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInput() []int {

	var instructions []int

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		scanInt, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		instructions = append(instructions, scanInt)
	}

	return instructions
}

func problemOne(instructions []int) int {
	steps := 0
	curInst := 0
	lenInst := len(instructions)

	for {
		if curInst < 0 || curInst >= lenInst {
			break
		}
		nextInst := curInst + instructions[curInst]
		instructions[curInst]++
		steps++
		curInst = nextInst
	}

	return steps
}

func problemTwo(instructions []int) int {
	steps := 0
	curInst := 0
	lenInst := len(instructions)

	for {
		if curInst < 0 || curInst >= lenInst {
			break
		}
		nextInst := curInst + instructions[curInst]
		if instructions[curInst] >= 3 {
			instructions[curInst]--
		} else {
			instructions[curInst]++
		}

		steps++
		curInst = nextInst
	}

	return steps
}

func main() {
	instructions := readInput()
	fmt.Println(problemOne(instructions))
	fmt.Println(problemTwo(instructions))
}
