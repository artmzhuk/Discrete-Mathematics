package main

import (
	"fmt"
)

/* fraction struct is representing common fraction and has 2 fields:
*  1)numerator
*  2)denominator
*  This struct supports methods:
*  1)Abs (absolute value)
*  2)Add +
*  3)Subtract -
*  4)Multiply *
*  5)Divide /
*  6)Compare
 */
type frac struct {
	numerator   int
	denominator int
}

func main() {
	matrix, N := matrixInput()
	solutions := gaussMethod(matrix, N)
	if solutions != nil {
		for i := range solutions {
			fmt.Printf("%d/%d\n", solutions[i].numerator, solutions[i].denominator)
		}
	} else {
		fmt.Println("No solution") //system either inconsistent or has infinitely many solutions
	}
}

func intToAbs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func (f frac) Abs() frac {
	if f.numerator < 0 {
		res := frac{-f.numerator, f.denominator}
		return res
	} else {
		return f
	}
}

func GCD(a, b int) int {
	a = intToAbs(a)
	b = intToAbs(b)
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

func LCM(a, b int) int {
	res := intToAbs((a * b) / GCD(a, b))
	return res
}

func (f frac) Simplify() frac {
	if f.numerator != 0 {
		gcd := GCD(intToAbs(f.numerator), intToAbs(f.denominator))
		f.numerator /= gcd
		f.denominator /= gcd
	} else {
		f.denominator = 1
	}
	if f.denominator < 0 {
		f.numerator *= -1
		f.denominator *= -1
	}
	return f
}

func (f frac) Sum(b frac) frac {
	denominatorLCM := LCM(f.denominator, b.denominator)
	fCoefficient := denominatorLCM / f.denominator
	bCoefficient := denominatorLCM / b.denominator
	f.numerator = f.numerator*fCoefficient + b.numerator*bCoefficient
	f.denominator = denominatorLCM
	c := f.Simplify()
	return c
}

func (f frac) Subtract(b frac) frac {
	denominator := LCM(f.denominator, b.denominator)
	fCoefficient := denominator / f.denominator
	bCoefficient := denominator / b.denominator
	f.numerator = f.numerator*fCoefficient - b.numerator*bCoefficient
	f.denominator = denominator
	return f.Simplify()
}

func (f frac) Multiply(b frac) frac {
	f.numerator = f.numerator * b.numerator
	f.denominator = f.denominator * b.denominator
	return f.Simplify()
}

func (f frac) Divide(b frac) frac {
	f.numerator = f.numerator * b.denominator
	f.denominator = f.denominator * b.numerator
	return f.Simplify()
}

func (f frac) Compare(b frac) int {
	denominatorLCM := LCM(f.denominator, b.denominator)
	fCoefficient := denominatorLCM / f.denominator
	bCoefficient := denominatorLCM / b.denominator
	return f.numerator*fCoefficient - b.numerator*bCoefficient
}

func swapRow(matrix [][]frac, i, j int) {
	tmp := make([]frac, len(matrix[0]))
	copy(tmp, matrix[i][:])
	copy(matrix[i][:], matrix[j][:])
	copy(matrix[j][:], tmp)
}

func forwardElimination(matrix [][]frac, N int) bool { // returns whether matrix singular or not
	for k := 0; k < N; k++ {
		iMax := k
		vMax := matrix[iMax][k]

		for i := k + 1; i < N; i++ {
			if matrix[i][k].Abs().Compare(vMax) > 0 {
				vMax = matrix[i][k]
				iMax = i
			}
		}
		if matrix[k][iMax].numerator == 0 {
			return false
		}
		if iMax != k {
			swapRow(matrix, k, iMax)
		}
		for i := k + 1; i < N; i++ {
			var f frac
			f = matrix[i][k].Divide(matrix[k][k])
			for j := k + 1; j <= N; j++ {
				matrix[i][j] = matrix[i][j].Subtract(f.Multiply(matrix[k][j]))
			}
			matrix[i][k].numerator = 0
		}
	}
	return true
}

func backSubstitution(matrix [][]frac, N int) []frac {
	solutions := make([]frac, N)
	for i := range solutions {
		solutions[i].denominator = 1
	}

	for i := N - 1; i >= 0; i-- {
		solutions[i] = matrix[i][N]

		for j := i + 1; j < N; j++ {
			solutions[i] = solutions[i].Subtract(solutions[j].Multiply(matrix[i][j]))
		}
		solutions[i] = solutions[i].Divide(matrix[i][i])
	}
	return solutions
}

func matrixInput() ([][]frac, int) { //returns matrix [N][N+1] and N
	var N int
	_, _ = fmt.Scanln(&N)
	matrix := make([][]frac, N)
	for i := range matrix {
		matrix[i] = make([]frac, N+1)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N+1; j++ {
			_, _ = fmt.Scan(&matrix[i][j].numerator)
			matrix[i][j].denominator = 1
		}
	}
	return matrix, N
}

func gaussMethod(matrix [][]frac, N int) []frac {
	if forwardElimination(matrix, N) {
		solutions := backSubstitution(matrix, N)
		return solutions
	} else {
		return nil
	}
}
