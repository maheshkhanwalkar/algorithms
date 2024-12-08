package oop_test

import (
	"algorithms/oop"
	"testing"
)

func TestPlus(t *testing.T) {
	a := oop.Rational{Numerator: 1, Denominator: 2}
	b := oop.Rational{Numerator: 2, Denominator: 6}

	res := a.Plus(b)

	if res.Numerator != 5 || res.Denominator != 6 {
		t.Errorf("expect 5/6, got %d/%d", res.Numerator, res.Denominator)
	}
}

func TestMinus(t *testing.T) {
	a := oop.Rational{Numerator: 1, Denominator: 2}
	b := oop.Rational{Numerator: 1, Denominator: 3}

	res := a.Minus(b)

	if res.Numerator != 1 || res.Denominator != 6 {
		t.Errorf("expect 1/6, got %d/%d", res.Numerator, res.Denominator)
	}
}

func TestMultiply(t *testing.T) {
	a := oop.Rational{Numerator: 3, Denominator: 2}
	b := oop.Rational{Numerator: 2, Denominator: 9}

	res := a.Multiply(b)

	if res.Numerator != 1 || res.Denominator != 3 {
		t.Errorf("expect 1/3, got %d/%d", res.Numerator, res.Denominator)
	}
}

func TestDivide(t *testing.T) {
	a := oop.Rational{Numerator: 3, Denominator: 2}
	b := oop.Rational{Numerator: 1, Denominator: 3}

	res := a.Divide(b)

	if res.Numerator != 9 || res.Denominator != 2 {
		t.Errorf("expect 9/2, got %d/%d", res.Numerator, res.Denominator)
	}
}
