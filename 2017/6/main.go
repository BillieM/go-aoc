package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInput() [16]int {
	var banks [16]int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		banksStr := strings.Fields(scanner.Text())
		for i, bankStr := range banksStr {
			bank, err := strconv.Atoi(bankStr)
			if err != nil {
				log.Fatal(err)
			}
			banks[i] = bank
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return banks

}

func problemOne(banks [16]int) int {
	cycles := 0
	var prevStates [][16]int

	for {
		if checkIfInStates(banks, prevStates) {
			break
		}
		prevStates = append(prevStates, banks)

		firstMaxInd := 0
		firstMaxVal := banks[firstMaxInd]

		for i, val := range banks {
			if val > firstMaxVal {
				firstMaxInd = i
				firstMaxVal = val
			}
		}

		banks[firstMaxInd] = 0

		bankInd := getNextIndex(firstMaxInd)

		for i := 0; i < firstMaxVal; i++ {

			banks[bankInd]++

			bankInd = getNextIndex(bankInd)
		}

		cycles++
	}

	return cycles
}

func problemTwo(banks [16]int) int {
	cycles := 0
	var prevStates [][16]int
	var seenStates [][16]int
	for {
		if checkIfInStates(banks, seenStates) {
			break
		}
		if checkIfInStates(banks, prevStates) {
			seenStates = append(seenStates, banks)
		}
		prevStates = append(prevStates, banks)

		firstMaxInd := 0
		firstMaxVal := banks[firstMaxInd]

		for i, val := range banks {
			if val > firstMaxVal {
				firstMaxInd = i
				firstMaxVal = val
			}
		}

		banks[firstMaxInd] = 0

		bankInd := getNextIndex(firstMaxInd)

		for i := 0; i < firstMaxVal; i++ {

			banks[bankInd]++

			bankInd = getNextIndex(bankInd)
		}

		cycles++
	}

	return cycles
}

func checkIfInStates(banks [16]int, prevStates [][16]int) bool {
	for _, prevState := range prevStates {
		if banks == prevState {
			return true
		}
	}
	return false
}

func getNextIndex(curIndex int) int {
	if curIndex <= 14 {
		return curIndex + 1
	}
	return 0
}

func main() {
	banks := getInput()
	fmt.Println(problemOne(banks))
	fmt.Println(problemTwo(banks))
}
