package arithmetic

import (
	"testing"
)

func TestAddition(t *testing.T) {
	arith := NewAdapter()
	var a, b int32 = 1, 1
	var want int32 = 2

	res, err := arith.Addition(a, b)
	if err != nil {
		t.Fatalf("expected: %v, got: %v\n", want, err)
	}
	if res != want {
		t.Errorf("expected: %v, got: %v\n", want, res)
	}
}

func TestSubtraction(t *testing.T) {
	arith := NewAdapter()
	want := int32(0)
	var a, b int32 = 1, 1
	res, err := arith.Subtraction(a, b)
	if err != nil {
		t.Fatalf("expected: %v, got: %v\n", want, err)
	}
	if res != want {
		t.Errorf("expected: %v, got: %v\n", want, res)
	}
}

func TestMultiplication(t *testing.T) {
	arith := NewAdapter()
	want := int32(1)
	var a, b int32 = 1, 1

	res, err := arith.Multiplication(a, b)
	if err != nil {
		t.Fatalf("expected: %v, got: %v\n", want, err)
	}
	if res != want {
		t.Errorf("expected: %v, got: %v\n", want, res)
	}
}

func TestDivision(t *testing.T) {
	arith := NewAdapter()
	want := int32(1)
	res, err := arith.Division(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v\n", want, err)
	}
	if res != want {
		t.Errorf("expected: %v, got: %v\n", want, res)
	}
}
