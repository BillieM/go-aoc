package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type program struct {
	name        string
	weight      int
	stackWeight int
	childrenStr []string
	children    []*program
	parents     []*program
}

func getInput() []string {
	var lines []string

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func getStructsMap(lines []string) map[string]*program {
	structsMap := make(map[string]*program)

	for _, line := range lines {
		p := parseLine(line)
		structsMap[p.name] = &p
	}

	for _, p := range structsMap {
		childrenAndParents(p, structsMap)
	}

	return structsMap
}

func parseLine(line string) program {
	stringFields := strings.Fields(line)
	name := stringFields[0]
	weight, err := strconv.Atoi(strings.TrimRight(strings.TrimLeft(stringFields[1], "("), ")"))
	var childrenStr []string
	if err != nil {
		log.Fatal(err)
	}
	if len(stringFields) > 2 {
		childrenStr = stringFields[3:]
	}
	p := program{name: name, weight: weight, childrenStr: childrenStr}
	return p
}

func childrenAndParents(p *program, structsMap map[string]*program) {
	/*
		loop through childrenStr
			set children to their program
			set their parents to this program
	*/

	if len(p.childrenStr) > 0 {
		for _, child := range p.childrenStr {
			trimChild := strings.TrimRight(child, ",")
			structsMap[trimChild].parents = append(structsMap[trimChild].parents, p)
			p.children = append(p.children, structsMap[trimChild])
		}
	}
}

func getStackWeight(p *program) int {
	childrenWeight := 0
	childrenWeight += p.weight

	for _, child := range p.children {
		childrenWeight += getStackWeight(child)
	}

	return childrenWeight
}

func isBalanced(p *program) bool {

	for _, a := range p.children {
		for _, b := range p.children {
			if a.weight != b.weight {
				return false
			}
		}
	}
	return true
}

func problemOne(structsMap map[string]*program) {
	/*
		loop through structsMap, return name of the program with 0 parents
	*/
	counter := 0
	for _, p := range structsMap {
		if len(p.parents) == 0 {
			counter++
		}
	}
}

func problemTwo(structsMap map[string]*program) {
	/*
		for all nodes, need to check stack weight of all children is equal

		common child between uneven stacks is problem node

		stack weight calculated by recursively adding all children wt for each node

	*/
	for _, p := range structsMap {
		p.stackWeight = getStackWeight(p)
	}

	var minP *program
	minVal := 1000000

	for _, p := range structsMap {
		if !isBalanced(p) && p.stackWeight < minVal {
			minP = p
		}
	}

	_ = minP

	for _, child := range structsMap["aobgmc"].children {
		fmt.Println(child)
	}
}

func main() {
	lines := getInput()
	structsMap := getStructsMap(lines)
	// problemOne(structsMap)
	problemTwo(structsMap)
}
