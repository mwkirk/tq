package container

import "testing"

func TestSliceEqual_equalSlices(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	if !sliceEqual(a, b) {
		t.Error("expected slices to be equal")
	}
}

func TestSliceEqual_differentLength(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2}
	if sliceEqual(a, b) {
		t.Error("expected slices to be different")
	}
}

func TestSliceEqual_differentValues(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 4}
	if sliceEqual(a, b) {
		t.Error("expected slices to be different")
	}
}
