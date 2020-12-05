package day4

import (
	"testing"

	"github.com/nlowe/aoc2020/challenge"
)

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = partA(challenge.FromFile())
	}
}

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = partB(challenge.FromFile())
	}
}
