package main

import (
	"fmt"
)

func main() {
	/* TEST CODE*/
	mapa := make(map[int][]int)
	for i := 1; i < 1e5; i++ {
		mapa[i] = genChain(i)
	}
	fmt.Printf("es :%d \n", maxChain(mapa))
	printChain(mapa[maxChain(mapa)])
}

func printChain(values []int) {
	for key, val := range values {
		if key == 0 {
			fmt.Printf("%d", val)
			continue
		}
		fmt.Printf("->%d", val)
	}
	fmt.Println("")
	fmt.Printf("Number of steps%d", len(values))
	fmt.Println("")
}
func genChain(start int) []int {
	steps := []int{start}
	for i := 0; steps[len(steps)-1] != 1; i++ {
		if steps[len(steps)-1]%2 == 0 {
			steps = append(steps, steps[len(steps)-1]/2)
			continue
		}
		steps = append(steps, 3*steps[len(steps)-1]+1)
	}
	printChain(steps)
	return steps
}

func maxChain(info map[int][]int) int {
	var chainSteps int = 0
	var start int = 0

	for key, value := range info {
		size := len(value)
		if size > chainSteps {
			chainSteps = size
			start = key
		}
	}

	return start
}
