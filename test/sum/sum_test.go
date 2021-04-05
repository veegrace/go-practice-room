package sum

import (
	"testing"
)

func TestSum(t *testing.T) {
	sumB := Sum([]int{10, -2, 4})
	if sumB != 12 {
		t.Errorf("Wanted 12 but received %d", sumB)
	}
}
