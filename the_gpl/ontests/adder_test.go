package adder

import "testing"

func TestAdder(t *testing.T) {
	var tests = []struct {
		x    int
		y    int
		want int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{5, 5, 10},
		{10, 10, 20},
	}
	for _, test := range tests {
		if got := Adder(test.x, test.y); got != test.want {
			t.Errorf("adder(%d, %d) = %d", test.x, test.y, got)
		}
	}
}
