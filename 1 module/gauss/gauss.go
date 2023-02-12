package main

import (
	"fmt"
)

type frac struct {
	num   int
	denom int
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func (f frac) abs() frac {
	if f.num < 0 {
		res := frac{-f.num, f.denom}
		return res
	} else {
		return f
	}
}

func gcd(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

func lcm(a, b int) int {
	res := (a * b) / gcd(a, b)
	return res
}

func (f *frac) simplify() {
	if f.num != 0 {
		gcd := gcd(intAbs(f.num), intAbs(f.denom))
		f.num /= gcd
		f.denom /= gcd
	}
}

func (f frac) sum(a, b frac) frac {
	denomLCM := lcm(a.denom, b.denom)
	aCoef := denomLCM / a.denom
	bCoef := denomLCM / b.denom
	f.num = a.num*aCoef + b.num*bCoef
	f.denom = denomLCM
	f.simplify()
	return f
}

func (f frac) subtract(a, b frac) frac {
	denomLCM := lcm(a.denom, b.denom)
	aCoef := denomLCM / a.denom
	bCoef := denomLCM / b.denom
	f.num = a.num*aCoef - b.num*bCoef
	f.denom = denomLCM
	f.simplify()
	return f
}

func (f frac) mult(a, b frac) frac {
	f.num = a.num * b.num
	f.denom = a.denom * b.denom
	f.simplify()
	return f
}

func (f frac) div(a, b frac) frac {
	f.num = a.num * b.denom
	f.denom = a.denom * b.num
	f.simplify()
	return f
}

func (f frac) compare(a, b frac) int {
	denomLCM := lcm(a.denom, b.denom)
	aCoef := denomLCM / a.denom
	bCoef := denomLCM / b.denom
	return a.num*aCoef - b.num*bCoef
}

func (f frac) print() {
	fmt.Printf("%d/%d", f.num, f.denom)
}

func swapRow(matrix [][]frac, i, j int) {
	tmp := make([]frac, len(matrix[0]))
	copy(tmp, matrix[i][:])
	copy(matrix[i][:], matrix[j][:])
	copy(matrix[j][:], tmp)
}

func forwardElim(matrix [][]frac, N int) bool { // returns whether matrix singular or not
	for k := 0; k < N; k++ {
		iMax := k
		vMax := matrix[iMax][k]

		for i := k + 1; i < N; i++ {
			if matrix[i][k].compare(matrix[i][k].abs(), vMax) > 0 {
				vMax = matrix[i][k]
				iMax = i
			}
		}
		if matrix[k][iMax].num == 0 {
			return false
		}
		if iMax != k {
			swapRow(matrix, k, iMax)
		}
		for i := k + 1; i < N; i++ {
			var f frac
			f = f.div(matrix[i][k], matrix[k][k])
			for j := k + 1; j <= N; j++ {
				matrix[i][j] = f.subtract(matrix[i][j], f.mult(matrix[k][j], f))
			}
			matrix[i][k].num = 0
		}
	}
	return true
}

func backSub(matrix [][]frac, N int) {
	x := make([]frac, N)
	for i := range x {
		x[i].denom = 1
	}

	for i := N - 1; i >= 0; i-- {
		x[i] = matrix[i][N]

		for j := i + 1; j < N; j++ {
			x[i] = x[i].subtract(x[i], x[i].mult(matrix[i][j], x[j]))
		}
		x[i] = x[i].div(x[i], matrix[i][i])
	}
	fmt.Println("Solutions:")
	for i := 0; i < N; i++ {
		x[i].print()
		fmt.Println()
	}
}

func matrixInput() ([][]frac, int) { //returns matrix [N][N+1] and N
	var N int
	fmt.Scanln(&N)
	matrix := make([][]frac, N)
	for i := range matrix {
		matrix[i] = make([]frac, N+1)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N+1; j++ {
			fmt.Scan(&matrix[i][j].num)
			matrix[i][j].denom = 1
		}
	}
	return matrix, N
}

func printMatrix(matrix [][]frac) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[i][j].print()
			fmt.Print("\t")
		}
		fmt.Println()
	}
}

func main() {
	/*	var a = frac{-2, 9}
		var b = frac{4, 18}
		c := a.abs()
		fmt.Println(c)
		fmt.Println(a.compare(a, b))
		a.print()*/
	matrix, N := matrixInput()
	forwardElim(matrix, N)
	backSub(matrix, N)
}
