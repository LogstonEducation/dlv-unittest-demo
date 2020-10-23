package main

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	rng := map[int]int{
		1: 1,
		2: 1,
		3: 2,
		4: 3,
	}
	for n, expected := range rng {
		n, expected := n, expected
		t.Run(fmt.Sprintf("fibonacci(%d)", n), func(t *testing.T) {
			t.Parallel()

			res, err := fibonacci(n)
			if err != nil {
				t.Errorf("%w", err)
			}
			if res != expected {
				t.Errorf("fibonacci(%d) = %d; got %d", n, expected, res)
			}
		})
	}
}
