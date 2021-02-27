package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	adjReg string
	adjIns string
	adjAmt int
	cmpReg string
	cmpOpr string
	cmpAmt int
}

func getInput() [1000]instruction {

	var instructions [1000]instruction

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	i := 0

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		adjReg := fields[0]
		adjIns := fields[1]
		adjAmt, _ := strconv.Atoi(fields[2])
		cmpReg := fields[4]
		cmpOpr := fields[5]
		cmpAmt, _ := strconv.Atoi(fields[6])
		instruction := instruction{
			adjReg: adjReg, adjIns: adjIns, adjAmt: adjAmt, cmpReg: cmpReg, cmpOpr: cmpOpr, cmpAmt: cmpAmt,
		}
		instructions[i] = instruction
		i++
	}
	return instructions
}

func problemOne(instructions [1000]instruction) int {
	registers := make(map[string]int)

	for _, instruction := range instructions {
		if checkCmp(instruction, registers) {
			if instruction.adjIns == "inc" {
				registers[instruction.adjReg] += instruction.adjAmt
			} else {
				registers[instruction.adjReg] -= instruction.adjAmt
			}
		}
	}

	max := 0

	for _, register := range registers {
		if register > max {
			max = register
		}
	}

	return max
}

func problemTwo(instructions [1000]instruction) int {
	registers := make(map[string]int)

	max := 0

	for _, instruction := range instructions {
		if checkCmp(instruction, registers) {
			if instruction.adjIns == "inc" {
				registers[instruction.adjReg] += instruction.adjAmt
			} else {
				registers[instruction.adjReg] -= instruction.adjAmt
			}
			if registers[instruction.adjReg] > max {
				max = registers[instruction.adjReg]
			}
		}

	}

	return max
}

func checkCmp(instruction instruction, registers map[string]int) bool {
	/*
		> >= < <= == !=
	*/
	cmpReg := instruction.cmpReg
	cmpAmt := instruction.cmpAmt

	cmpRegVal := registers[cmpReg]

	switch instruction.cmpOpr {
	case ">":
		if cmpRegVal > cmpAmt {
			return true
		}
	case ">=":
		if cmpRegVal >= cmpAmt {
			return true
		}
	case "<":
		if cmpRegVal < cmpAmt {
			return true
		}
	case "<=":
		if cmpRegVal <= cmpAmt {
			return true
		}
	case "==":
		if cmpRegVal == cmpAmt {
			return true
		}
	case "!=":
		if cmpRegVal != cmpAmt {
			return true
		}
	}
	return false
}

func main() {
	instructions := getInput()
	fmt.Println(problemOne(instructions))
	fmt.Println(problemTwo(instructions))

}
