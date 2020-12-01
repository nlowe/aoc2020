package util

import "fmt"

func IntMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func IntClamp(low, n, high int) int {
	if low > high {
		panic(fmt.Errorf("IntClamp: low cannot be > high: %d > %d", low, high))
	}

	return IntMax(IntMin(n, high), low)
}
