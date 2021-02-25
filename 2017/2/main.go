package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInputData() [16][16]int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	n := 0
	var linesArr [16][16]int

	for scanner.Scan() {
		cells := strings.Fields(scanner.Text())
		for i, cell := range cells {
			linesArr[n][i], err = strconv.Atoi(cell)
		}
		n++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return linesArr
}

func problemOne(lines [16][16]int) int {
	sum := 0

	for _, line := range lines {
		min := line[0]
		max := line[0]

		for _, i := range line {
			if i < min {
				min = i
			}
			if i > max {
				max = i
			}
		}
		sum += (max - min)
	}

	return sum
}

func problemTwo(lines [16][16]int) int {
	sum := 0

	for _, line := range lines {
		for _, x := range line {
			for _, y := range line {
				if x != y {
					divis1 := float64(x) / float64(y)
					divis2 := float64(y) / float64(x)
					if divis1 == float64(int64(divis1)) {
						fmt.Println(x, y, divis1)
						sum += int(divis1)
					} else if divis2 == float64(int64(divis2)) {
						fmt.Println(x, y, divis2)
						sum += int(divis2)
					}
				}
			}
		}
	}

	return sum / 2
}

func main() {
	linesArr := getInputData()
	fmt.Println(problemOne(linesArr))
	fmt.Println(problemTwo(linesArr))
}
