package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func getInput() [512][]string {
	/*
		create a 512 len array of slices
			loop through lines of txt file, gen slice and put in array
	*/

	var lines [512][]string
	_ = lines

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	n := 0
	for scanner.Scan() {
		phrases := strings.Fields(scanner.Text())
		for _, phrase := range phrases {
			lines[n] = append(lines[n], phrase)
		}
		n++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func problemOne(lines [512][]string) int {
	invalids := 0
	for _, line := range lines {
		var unique []string
		for _, phrase := range line {
			if stringInSlice(phrase, unique) {
				invalids++
				break
			} else {
				unique = append(unique, phrase)
			}
		}
	}
	return (512 - invalids)
}

func problemTwo(lines [512][]string) int {
	invalids := 0
	for _, line := range lines {
		var unique []string
		for _, phrase := range line {
			if anagramInSlice(phrase, unique) {
				invalids++
				break
			} else {
				unique = append(unique, phrase)
			}
		}
	}
	return (512 - invalids)
}
func stringInSlice(a string, slc []string) bool {
	for _, b := range slc {
		if a == b {
			return true
		}
	}
	return false
}

func anagramInSlice(a string, slc []string) bool {
	aSort := sortString(a)

	for _, b := range slc {
		bSort := sortString(b)
		if aSort == bSort {
			return true
		}
	}
	return false
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	lines := getInput()
	fmt.Println(problemOne(lines))
	fmt.Println(problemTwo(lines))

}
