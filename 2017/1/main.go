package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getInputData() string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func problemOne(inputString string) int {
	sum := 0

	inputArr := []rune(inputString)

	for i := 0; i < len(inputArr)-1; i++ {
		if inputArr[i] == inputArr[i+1] {
			sum += int(inputArr[i] - '0')
		}
	}

	if inputArr[0] == inputArr[len(inputArr)-1] {
		sum += int(inputArr[0] - '0')
	}

	return sum
}

func getNeighborIndex(i int, travSteps int) int {
	if i < travSteps {
		return i + travSteps
	}
	return i - travSteps
}

func problemTwo(inputString string) int {
	sum := 0

	/*
		for each digit, need to calculate the index of its partner

		take the length of the array
		divide by 2, this is the number of steps we need to traverse

		something with minusing current position from length

	*/

	inputArr := []rune(inputString)

	travSteps := len(inputArr) / 2

	for i := 0; i < len(inputArr); i++ {
		neighborIndex := getNeighborIndex(i, travSteps)
		thisNum := int(inputArr[i] - '0')
		neighNum := int(inputArr[neighborIndex] - '0')
		if thisNum == neighNum {
			sum += thisNum
		}
	}

	return sum
}

func main() {
	inputString := getInputData()
	fmt.Println(problemOne(inputString))
	fmt.Println(problemTwo(inputString))
}
