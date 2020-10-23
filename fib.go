package main

import "fmt"

// Return the value of the n-th fibonacii number.
// 1, 2, 3, 4, 5, 6, 7
// 0, 1, 1, 2, 3, 5, 8
func fibonacci(n int) (int, error) {
	if n < 1 {
		return -1, fmt.Errorf("Number is too low: %d", n)
	}

	if n == 1 {
		return 0, nil
	}

	if n == 2 {
		return 1, nil
	}

	j, k := 0, 1

	for i := 3; i <= n; i++ {
		j, k = k, j+k
	}

	return k, nil
}

func main() {
	rng := []int{1, 2, 3, 4, 5, 6}

	for x := range rng {
		result, err := fibonacci(x)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(result)
	}
}
