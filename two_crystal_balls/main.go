package main

import (
	"fmt"
	"math"
)

// The two crystal ball problem can be solved similar to a binary search
// The breaks boolean is all false until the the first true, which signifies when it breaks
// So this is just like it is sorted.
// If you jump a sqrt(len(breaks)) amount of time then you can get O(sqrt(n)) running time for the problem.
// Once we land on our first break, we need to jump back the sqrt(n), and then walk forward until we find that first break.
// if we don't ever find one then just return -1
func two_crystal_balls(breaks []bool) int {
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))
	i := jumpAmount

	for ; i < len(breaks); i += jumpAmount {
		if breaks[i] {
			break
		}
	}

	i -= jumpAmount
	j := 0
	for j < jumpAmount && i < len(breaks) {
		if breaks[i] {
			return i
		}
		j++
		i++
	}

	return -1
}

func main() {
	breaks := [10]bool{false, false, false, false, false, false, false, false, true, true}

	breakIndex := two_crystal_balls(breaks[:])

	fmt.Printf("The break index is %d\n", breakIndex)
}
