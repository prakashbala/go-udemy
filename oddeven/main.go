package main

import "fmt"

func main() {

	numbers := []int{}

	for _, i := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		numbers = append(numbers, i)
	}

	for _, v := range numbers {
		if v%2 == 0 {
			fmt.Printf("%v is even number\n", v)
		} else {
			fmt.Printf("%v is odd number\n", v)
		}
	}

}
