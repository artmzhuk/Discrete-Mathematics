package main

import "fmt"

func partition(low, high int, less func(i, j int) bool, swap func(i, j int)) int {
	i, j := low, low
	for j < high {
		if less(j, high) {
			swap(i, j)
			i++
		}
		j++
	}
	swap(i, high)
	return i
}

func qsortRec(low, high int, less func(i, j int) bool, swap func(i, j int)) {
	if low < high {
		q := partition(low, high, less, swap)
		qsortRec(low, q-1, less, swap)
		qsortRec(q+1, high, less, swap)
	}
}

func qsort(n int, less func(i, j int) bool, swap func(i, j int)) {
	qsortRec(0, n-1, less, swap)
}

func test() {
	var testArray = []int{557, 765, 705, 218, 54, 673, -4, 79, -866, 586}
	fmt.Println(testArray)
	qsort(len(testArray),
		func(i, j int) bool { return testArray[i] < testArray[j] },
		func(i, j int) { testArray[i], testArray[j] = testArray[j], testArray[i] })
	fmt.Println(testArray)
}

func main() {
	test()
}
