package main

import "fmt"

func exercise1(x int) (int, bool) {
	return x / 2, (x % 2) == 0
}

func greater(x ...uint) uint {
	var greater uint = 0
	for _, i := range x {
		if greater < i {
			greater = i
		}
	}
	return greater
}

func main() {
	my_function := func(x int) (int, bool) {
		return x / 2, (x % 2) == 0
	}
	half, even := my_function(3)
	fmt.Printf("3/2 is %d, %t \n", half, even)

	lista_numeros := []uint{1, 2, 34, 455, 2345, 567}

	fmt.Printf("The greatest number between 3 2 1 0 is: %d", greater(lista_numeros...))

}
