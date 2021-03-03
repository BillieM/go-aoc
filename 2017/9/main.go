package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getTestInput1() []string {
	return []string{
		"{}",
		"{{{}}}",
		"{{},{}}",
		"{{{},{},{{}}}}",
		"{<a>,<a>,<a>,<a>}",
		"{{<ab>},{<ab>},{<ab>},{<ab>}}",
		"{{<!!>},{<!!>},{<!!>},{<!!>}}",
		"{{<a!>},{<a!>},{<a!>},{<ab>}}",
	}
}
func getTestInput2() []string {
	return []string{
		"<>",
		"<random characters>",
		"<<<<>",
		"<{!>}>",
		"<!!>",
		"<!!!>>",
	}
}

func getInput() string {

	var input string

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = scanner.Text()
	}

	return input
}

func problemOne(input string) int {

	score := 0
	skipNext := false
	garbageClosed := true
	groupLvl := 0

	for _, char := range input {
		if skipNext {
			skipNext = false
			continue
		}

		if char == '!' {
			skipNext = true
		}

		if char == '{' {

			if garbageClosed {
				groupLvl++
			}

		}

		if char == '}' {

			if garbageClosed {
				score += groupLvl
				groupLvl--
			}

		}

		if char == '<' {
			garbageClosed = false
		}

		if char == '>' {
			garbageClosed = true
		}

	}

	return score

}

func problemTwo(input string) int {

	skipNext := false
	garbageClosed := true
	garbageCount := 0

	for _, char := range input {
		if skipNext {
			skipNext = false
			continue
		}

		if char == '<' && garbageClosed {
			garbageClosed = false
			continue
		}

		if char == '>' {
			garbageClosed = true
			continue
		}

		if char == '!' {
			skipNext = true
			continue
		}

		if !garbageClosed {
			garbageCount++
		}

	}

	return garbageCount

}

func main() {
	input := getInput()
	fmt.Println(problemOne(input))
	fmt.Println(problemTwo(input))

}

/*
we track the current group level

if a group is broken, decrement the group level

if a group is closed, increase score and decrement group level
*/
