package sum

import "testing"

func TestSum(t *testing.T) {
	sum := Sum(3, 4)
	expected := 7

	if sum != expected {
		t.Fatalf("Got %v expected %v", sum, expected)
	}
}
