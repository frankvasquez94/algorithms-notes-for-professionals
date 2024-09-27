package main

import (
	"algorithms-notes-for-professionals/fizzbuzz"
	"context"
	"fmt"
)

func main() {
	fmt.Println("Fizz Buzz algorithm")

	fizzBuzzAlg := fizzbuzz.New()
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fizzBuzzAlg.Execute(context.Background(), numbers)
}
