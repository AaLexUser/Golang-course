package main

import "fmt"

func Sum(numbers []int) int {
	// TODO: верните сумму элементов, не меняя срез.
	res := 0
	for _, num := range numbers {
		res += num
	}
	return res
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	fmt.Println(Sum(numbers))                // 30
	fmt.Println(Sum([]int{}))                // 0
	fmt.Println(Sum(nil))                    // 0
	fmt.Println(Sum([]int{-2, -1, 0, 1, 2})) // 0

	// Add 12 to the slice
	numbers = append(numbers, 12)
	fmt.Println(len(numbers)) // 6
	fmt.Println(Sum(numbers)) // 42

	// Share the slice
	shared := numbers
	shared[0] = 100
	fmt.Println(numbers[0], shared[0]) // 100 100
}
