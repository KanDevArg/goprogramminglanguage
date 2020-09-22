package sumkit

import "testing"

// Test SumInts
func TestSumInts(t *testing.T) {
	s := SumInts(1, 2)
	if s != 3 {
		t.Errorf("Sum is not equals 3, but %v", s)
	}

	s = SumInts(1, -1)
	if s != 0 {
		t.Errorf("Sum is not 0 but %v", s)
	}

	s = SumInts()
	if s != 0 {
		t.Errorf("Sum of no numbers should be 0; got %v", s)
	}
}
