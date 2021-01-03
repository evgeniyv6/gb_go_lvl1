package fiblib

import "testing"

func TestFibonacciGoldenRatio(t *testing.T) {
	tests := []struct {
		name string
		in   uint64
		want uint64
	}{
		{"simple", 3, 2},
		{"negative", 0, 100},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := FibonacciGoldenRatio(tc.in)
			if got != tc.want {
				t.Fatalf("%s: expected %v; got - %v", tc.name, tc.want, got)
			}
		})
	}
}
