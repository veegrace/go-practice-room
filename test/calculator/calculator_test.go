package calculator

import "testing"

var num Numbers = Numbers{16.0, 8.0}

func TestCalculator_Add(t *testing.T) {
	got := num.add()
	want := 24.0

	if got != want {
		t.Errorf("Wanted %v but got %v", want, got)
	}
	t.Log("test passed")

}

func TestCalculator_Subtract(t *testing.T) {
	got := num.subtract()
	want := 8.0

	if got != want {
		t.Errorf("Wanted %v but got %v", want, got)
	}
	t.Log("test passed")
}

func TestCalculator_Multiply(t *testing.T) {
	got := num.multiply()
	want := 128.0

	if got != want {
		t.Errorf("Wanted %v but got %v", want, got)
	}
	t.Log("test passed")
}

func TestCalculator_Divide(t *testing.T) {
	got := num.divide()
	want := 2.0

	if got != want {
		t.Errorf("Wanted %v but got %v", want, got)
	}
	t.Log("test passed")
}
