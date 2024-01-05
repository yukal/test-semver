package main

import (
	"fmt"
	"testing"
)

func TestIntMaxBasic(t *testing.T) {
	ans := IntMax(2, -2)

	if ans != 2 {
		t.Errorf("IntMax(2, -2) = %d; want 2", ans)
	}
}

func TestIntMaxTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 1},
		{1, 0, 1},
		{2, -2, 2},
		{0, -1, 0},
		{-1, 0, 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)

		t.Run(testname, func(t *testing.T) {
			ans := IntMax(tt.a, tt.b)

			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarkIntMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMax(1, 2)
	}
}
