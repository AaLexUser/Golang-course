package main

import "fmt"

func Increment(value *int) {
	// TODO: увеличьте число по указателю и обработайте nil.
	if value != nil {
		*value++
	}
}

func main() {
	number := 10
	Increment(&number)
	fmt.Println(number) // 11

	Increment(&number)
	fmt.Println(number) // 12

	Increment(nil) // программа продолжает работать
}
