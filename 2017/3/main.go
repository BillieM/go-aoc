package main

import (
	"fmt"
	"math"
)

func problemOne(input float64) float64 {

	var i float64 = 1

	var out float64 = 0

	for {

		var uPow float64 = 2*i + 1
		upperBound := math.Pow(uPow, 2)

		if upperBound > input {

			var lPow float64 = 2*(i-1) + 1
			lowerBound := math.Pow(lPow, 2) + 1

			var mid1 float64 = lowerBound + (i - 1)
			var mid2 float64 = mid1 + (2 * i)
			var mid3 float64 = mid2 + (2 * i)
			var mid4 float64 = mid3 + (2 * i)

			mid1Diff := math.Abs(mid1 - input)
			mid2Diff := math.Abs(mid2 - input)
			mid3Diff := math.Abs(mid3 - input)
			mid4Diff := math.Abs(mid4 - input)

			min12 := math.Min(mid1Diff, mid2Diff)
			min34 := math.Min(mid3Diff, mid4Diff)
			minOut := math.Min(min12, min34)

			out += i
			out += minOut

			break
		}

		i++
	}

	return out
}

func problemTwo(input float64) {

}

func main() {
	var input float64 = 312051

	fmt.Println(problemOne(input))
	problemTwo(input)
}
