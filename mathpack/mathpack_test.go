package mathpack

import (
	"testing"
)

func TestPow(t *testing.T) {

	i := 5.0
	n := 2.0
	actual := Pow(i, n)
	expected := 25.0

	if expected != actual {
		t.Errorf("mathpack.Pow failed, expected %v, got %v", expected, actual)
	}
}

func TestPowCanBeZero(t *testing.T) {

	i := 5.0
	n := 0.0
	actual := Pow(i, n)
	expected := 1.0

	if expected != actual {
		t.Errorf("mathpack.Pow failed, expected %v, got %v", expected, actual)
	}
}

func TestPowCanBeNegative(t *testing.T) {

	i := 5.0
	n := -1.0
	actual := Pow(i, n)
	expected := 0.2

	if expected != actual {
		t.Errorf("mathpack.Pow failed, expected %v, got %v", expected, actual)
	}
}
