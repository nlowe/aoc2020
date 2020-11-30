package util

import "strconv"

func MustAtoI(a string) int {
	result, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}

	return result
}
