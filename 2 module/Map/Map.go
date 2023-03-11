package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	// Считываем карту
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	// Вычисляем минимальную длину пути
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}
	dp[0][0] = a[0][0]
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i > 0 && (dp[i-1][j] == -1 || dp[i-1][j] > dp[i][j]+a[i-1][j]) {
				dp[i-1][j] = dp[i][j] + a[i-1][j]
			}
			if j > 0 && (dp[i][j-1] == -1 || dp[i][j-1] > dp[i][j]+a[i][j-1]) {
				dp[i][j-1] = dp[i][j] + a[i][j-1]
			}
			if i < n-1 && (dp[i+1][j] == -1 || dp[i+1][j] > dp[i][j]+a[i+1][j]) {
				dp[i+1][j] = dp[i][j] + a[i+1][j]
			}
			if j < n-1 && (dp[i][j+1] == -1 || dp[i][j+1] > dp[i][j]+a[i][j+1]) {
				dp[i][j+1] = dp[i][j] + a[i][j+1]
			}
		}
	}
	fmt.Println(dp[n-1][n-1])
}
