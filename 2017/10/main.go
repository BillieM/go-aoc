package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInput() []int {

	var inputArr []int

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		strArr := strings.Split(scanner.Text(), ",")
		for _, str := range strArr {
			intToAppend, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
			inputArr = append(inputArr, intToAppend)

		}
	}
	return inputArr
}

func problemOne(input []int) int {

	var list [256]int

	listLen := len(list)

	curPos := 0

	for i := 0; i < listLen; i++ {
		list[i] = i
	}

	for skipSize, length := range input {

		nextListCycle(&list, curPos, length, listLen)

		curPos = getNextPos(curPos, length, skipSize, listLen)

	}

	return list[0] * list[1]
}

func getNextPos(curPos int, length int, skipSize int, listLen int) int {
	return (curPos + length + skipSize) % listLen
}

func nextListCycle(list *[256]int, curPos int, length int, listLen int) {
	/*
		check if arr to reverse needs to wrap round
	*/

	var revArr []int

	if curPos+length >= listLen {
		cntFirstArr := 0
		cntSecondArr := 0
		for _, item := range list[curPos:listLen] {
			revArr = append(revArr, item)

			cntFirstArr++
		}

		for _, item := range list[0 : length-cntFirstArr] {
			revArr = append(revArr, item)

			cntSecondArr++
		}

		reverseArr(&revArr)
		revArrIndex := 0
		for i := 0; i < cntFirstArr; i++ {
			list[curPos+i] = revArr[revArrIndex]
			revArrIndex++
		}

		for i := 0; i < cntSecondArr; i++ {
			list[i] = revArr[revArrIndex]
			revArrIndex++
		}
	} else {
		revArr = list[curPos : curPos+length]
		reverseArr(&revArr)
		for i, item := range revArr {
			list[curPos+i] = item
		}

	}

}

func reverseArr(arr *[]int) {
	for i, j := 0, len(*arr)-1; i < j; i, j = i+1, j-1 {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}

func main() {
	input := getInput()
	// fmt.Println(problemOne([]int{3, 4, 1, 5}))
	fmt.Println(problemOne(input))

}
