package fizzbuzz

import "context"

type FizzBuzz interface {
	Execute(ctx context.Context, numbers []int)
}

type fizzbuzz struct {
}

func (fb fizzbuzz) Execute(ctx context.Context, numbers []int) {

}
