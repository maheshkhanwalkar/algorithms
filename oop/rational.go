package oop

// Algorithms 1.2 Data Abstraction: https://algs4.cs.princeton.edu/12oop/
// Creative Problems 16

// Rational number
type Rational struct {
	Numerator   int
	Denominator int
}

// Plus add a and b together
func (a Rational) Plus(b Rational) Rational {
	denominator := a.Denominator * b.Denominator
	numerator := a.Numerator*b.Denominator + b.Numerator*a.Denominator

	gcdFactor := gcd(numerator, denominator)
	return Rational{Numerator: numerator / gcdFactor, Denominator: denominator / gcdFactor}
}

// Minus subtract b from a
func (a Rational) Minus(b Rational) Rational {
	// a - b = a + (-b)
	return a.Plus(Rational{Numerator: -b.Numerator, Denominator: b.Denominator})
}

// Multiply a and b together
func (a Rational) Multiply(b Rational) Rational {
	numerator := a.Numerator * b.Numerator
	denominator := a.Denominator * b.Denominator

	gcdFactor := gcd(numerator, denominator)
	return Rational{Numerator: numerator / gcdFactor, Denominator: denominator / gcdFactor}
}

// Divide a by b
func (a Rational) Divide(b Rational) Rational {
	// (a1/a2) / (b1/b2) = (a1/a2) * (b2/b1)
	return a.Multiply(Rational{Numerator: b.Denominator, Denominator: b.Numerator})
}

// gcd compute the GCD of two integers using Euclid's algorithm
func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}

	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	remainder := a % b
	return gcd(b, remainder)
}
