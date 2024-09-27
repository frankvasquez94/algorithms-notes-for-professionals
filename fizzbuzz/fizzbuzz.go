package fizzbuzz

import (
	"context"
	"fmt"
)

type FizzBuzz interface {
	Execute(ctx context.Context, numbers []int)
}

type fizzbuzz struct {
}

func New() *fizzbuzz {
	return &fizzbuzz{}
}

func (fb fizzbuzz) Execute(ctx context.Context, numbers []int) {
	p := fmt.Println
	for _, number := range numbers {
		// This is the same as number % 3 && number % 5
		// But when we are evaluating algorithm efficiency is a key point.
		// In this case we reduce the two operations to one, so we can say that is most efficient.
		if number%15 == 0 {
			p(fmt.Sprintf("%d, Fizz Buzz", number))
		} else if number%3 == 0 {
			p(fmt.Sprintf("%d, Fizz", number))
		} else if number%5 == 0 {
			p(fmt.Sprintf("%d, Buzz", number))
		} else {
			p(fmt.Sprintf("%d", number))
		}
	}

}
