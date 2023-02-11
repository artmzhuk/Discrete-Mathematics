/*Sources:
intToSlice, test functions are used to test add function
they might be stupid, but were written just to practise
*/

package main

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"reflect"
	"time"
)

func add(a, b []int32, p int) []int32 {
	maxLength := int(math.Max(float64(len(a)), float64(len(b)))) // max of a and b length's
	minLength := int(math.Min(float64(len(a)), float64(len(b)))) // min of a and b length's
	base := int32(p)
	over := int32(0)
	result := make([]int32, 0, maxLength)
	for i := 0; i < maxLength; i++ {
		if i < minLength {
			sum := a[i] + b[i] + over
			result = append(result, sum%base)
			over = sum / base
		} else { // minLength < i < maxLength
			if i >= len(a) {
				sum := b[i] + over
				result = append(result, sum%base)
				over = sum / base
			} else {
				sum := a[i] + over
				result = append(result, sum%base)
				over = sum / base
			}
		}
	}
	result = append(result, over)
	if result[len(result)-1] == 0 {
		result = result[0 : len(result)-1]
	}
	return result
}

func intToSlice(a int64) []int32 {
	result := make([]int32, 10)
	for n := a; n > 0; n /= 10 {
		result = append(result, int32(n%10))
	}

	return result
}

func test(n int) {
	for i := 0; i < n; i++ {
		rand.NewSource(time.Now().UnixNano())
		a := rand.Int31()
		a %= 1000
		rand.NewSource(time.Now().UnixNano())
		b := rand.Int31()
		b %= 1000
		var a1, b1, sum1 big.Int
		a1.SetInt64(int64(a))
		b1.SetInt64(int64(b))
		sum1.Add(&a1, &b1)
		sum64 := sum1.Int64()
		sumSlice := add(intToSlice(int64(a)), intToSlice(int64(b)), 10)

		if reflect.DeepEqual(sumSlice, intToSlice(sum64)) {
			fmt.Println(a, "+", b, "=", sum64, "Passed")
		} else {
			fmt.Println(intToSlice(int64(a)), "+", intToSlice(int64(b)), "!=", sum64, "Failed", "\n",
				add(intToSlice(int64(a)), intToSlice(int64(b)), 10), "!=", intToSlice(sum64))
		}
	}

}

func main() {
	test(1000)
}
