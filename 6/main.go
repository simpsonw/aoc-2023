package main

import (
	"fmt"
	"math"
)

func main()  {
	/*
	timeToDistance := map[int]int{
		7: 9,
		15:40,
		30:200,
	}
	 */

	timeToDistance := map[int]int {
		40828492: 233101111101487,
	}

	productOfWinningCombinations := 1
	for time, recordDistance := range timeToDistance {
		fmt.Printf("Time: %d Record Distance: %d\n", time, recordDistance)
		// This uses the smart way of solving this that I found on Reddit
		//winningCombinations := countWaysToWin_Optimal(time, recordDistance)
		var speed, distance, remainingTime, winningCombinations int
		// Brute force
		for i :=0; i < time; i++ {
			remainingTime = time-i
			distance = remainingTime*speed
			if distance > recordDistance {
				fmt.Printf("\tHolding the button for %dms, travelling %dmm at a speed of %dmm/s\n", i, distance, speed)
				winningCombinations++
			}
			speed++
		}
		productOfWinningCombinations *= winningCombinations
	}
	fmt.Printf("The product of all winning combinations is %d\n", productOfWinningCombinations)
}

// Smarter way of doing this I found on Reddit
func countWaysToWin_Optimal(time int, record int) int {
	// a = -1 due to the step time being constant and a < 0 indicates a
	// parabola open downward (upside down U)
	a := float64(-1)
	// b is the duration of the race and we shift the parabola to the right of
	// the y-axis
	b := float64(time)
	// c shifts the roots of the parabola to the edges of the record we need to
	// beat
	c := float64(-(record + 1))

	// Use the quadratic formula to find the x roots
	x1 := (-b + math.Sqrt(b*b-4*a*c)) / (2 * a)
	x2 := (-b - math.Sqrt(b*b-4*a*c)) / (2 * a)

	return int(math.Floor(x2) - math.Ceil(x1) + 1)
}