package util

import (
	"fmt"
)

func IntAbs(n int) int {
	if n < 0 {
		return -1 * n
	}

	return n
}

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

func ManhattanDistance(x1, y1, x2, y2 int) int {
	return IntAbs(x2-x1) + IntAbs(y2-y1)
}
