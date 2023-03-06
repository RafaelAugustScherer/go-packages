package my_math

import "testing"

func TestSum(t *testing.T) {
	expected := 6
	received := Sum(3, 3)

	if expected != received {
		t.Errorf("Sum(3, 3): Expected %v. Received %v", expected, received)
	}
}
