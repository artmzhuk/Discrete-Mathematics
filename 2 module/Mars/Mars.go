package main

import (
	"fmt"
	"math"
)

type solutions struct {
	colors  [][]int
	balance []int
}

type Graph struct {
	adjList [][]int
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func inputGraph() *Graph {
	var n int
	_, _ = fmt.Scan(&n)
	list := make([][]int, n)
	for i := 0; i < n; i++ {
		list[i] = make([]int, 0)
		for j := 0; j < n; j++ {
			var sign string
			_, _ = fmt.Scan(&sign)
			if ([]rune(sign))[0] == '+' {
				list[i] = append(list[i], j)
			}
		}
	}
	g := new(Graph)
	g.adjList = list
	return g
}

func findSolutions(g *Graph) solutions {
	visited := make([]bool, len(g.adjList))
	possibleSolutions := make([][]int, 0)
	balance := make([]int, 0)
	for i := 0; i < len(g.adjList); i++ {
		if visited[i] {
			continue
		}
		isBipartite := true
		color := make([]int, len(g.adjList))
		thisBalance := 1
		/*		if color[i] != 0 {
				continue
			}*/
		queue := make([]int, 0)
		color[i] = 1
		visited[i] = true
		queue = append(queue, i)
		for len(queue) != 0 && isBipartite {
			popped := queue[0]
			queue = queue[1:]
			for j := range g.adjList[popped] {
				if color[g.adjList[popped][j]] == 0 {
					if color[popped] == 1 {
						color[g.adjList[popped][j]] = 2
						thisBalance--
					} else {
						color[g.adjList[popped][j]] = 1
						thisBalance++
					}
					visited[g.adjList[popped][j]] = true
					queue = append(queue, g.adjList[popped][j])
				} else {
					if color[g.adjList[popped][j]] == color[popped] {
						isBipartite = false
						break
					}
				}
			}
		}
		if isBipartite {
			possibleSolutions = append(possibleSolutions, color)
			balance = append(balance, thisBalance)
		}
	}
	if len(possibleSolutions) == 0 {
		fmt.Println("No solution")
		return solutions{
			colors:  nil,
			balance: nil,
		}
	} else {
		return solutions{
			colors:  possibleSolutions,
			balance: balance,
		}
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func main() {
	sol := findSolutions(inputGraph())
	fmt.Println(sol)
	//sel(sol)
	/*	if solutions != nil {
		selectSolution(solutions)
	}*/
}

func sel(solutions solutions) int {
	maxInt := math.MaxInt
	flag := 1
	sum := 0
	A := solutions.balance
	n := len(A)
	flipped := make([][]bool, 2)
	dp := make([]map[int]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make(map[int]int)
		//dp[1] = make(map[int]int)
		flipped[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		sum += A[i]
	}
	for i := -sum; i <= sum; i++ {
		dp[0][i] = maxInt
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := -sum; j <= sum; j++ {
			dp[flag][j] = maxInt
			if j-A[i-1] <= sum && j-A[i-1] >= -sum { //sign isn't flipped for A[i-1]
				dp[flag][j] = dp[flag^1][j-A[i-1]]
				flipped[flag][i-1] = false
			}
			if j+A[i-1] <= sum && j+A[i-1] >= -sum && dp[flag^1][j+A[i-1]] != maxInt { //sign is flipped for A[i-1]
				dp[flag][j] = min(dp[flag][j], dp[flag^1][j+A[i-1]]+1)
				flipped[flag][i-1] = true
			}
		}
		flag ^= 1
	}
	for i := 0; i <= sum; i++ {
		if dp[flag^1][i] != maxInt {
			return dp[flag^1][i]
		}
	}
	return n - 1
}

/*func selectSolution(s [][]int) {
  	diff := make([]int, len(s))
  	for i := range s {
  		count1 := 0
  		count2 := 0
  		for j := range s[i] {
  			if s[i][j] == 1 {
  				count1++
  			}
  			if s[i][j] == 2 {
  				count2++
  			}
  			if s[i][j] == 0 {
  				fmt.Println("panic")
  				os.Exit(10)
  			}
  		}
  		diff[i] = count1 - count2
  	}
  	similarDiffs := make([][]int, 0)
  	similarDiffsIds := make([]int, 0)
  	minDiff := diff[0]
  	for i := range diff {
  		if abs(diff[i]) < minDiff {
  			minDiff = diff[i]
  			similarDiffs = similarDiffs[0:0]
  			similarDiffsIds = similarDiffsIds[0:0]
  		}
  		if abs(diff[i]) == minDiff {
  			similarDiffs = append(similarDiffs, s[i])
  			similarDiffsIds = append(similarDiffsIds, i)
  		}
  	}
  	if len(similarDiffs) == 1 {
  		for i := range similarDiffs[0] {
  			if (diff[similarDiffsIds[0]] < 0 && s[similarDiffsIds[0]][i] == 2) ||
  				(diff[similarDiffsIds[0]] > 0 && s[similarDiffsIds[0]][i] == 1) ||
  				(diff[similarDiffsIds[0]] == 0 && s[similarDiffsIds[0]][i] == s[similarDiffsIds[0]][0]) {
  				fmt.Print(i+1, " ")
  			}
  		}
  	} else {
  		firstGroups := make([][]int, len(similarDiffs))
  		for i := range similarDiffs {
  			for j := range s[similarDiffsIds[i]] {
  				if (diff[similarDiffsIds[i]] < 0 && s[similarDiffsIds[i]][j] == 2) ||
  					(diff[similarDiffsIds[i]] > 0 && s[similarDiffsIds[i]][j] == 1) ||
  					(diff[similarDiffsIds[i]] == 0 && s[similarDiffsIds[i]][j] == s[similarDiffsIds[i]][0]) {
  					firstGroups[i] = append(firstGroups[i], j)
  				}

  			}
  		}
  		swap := reflect.Swapper(firstGroups)
  		for i := range firstGroups[0] {
  			for j := range firstGroups {
  				if firstGroups[j][i] > firstGroups[0][i] {
  					swap(j, 0)
  					firstGroups = firstGroups[1:]
  				} else if firstGroups[j][i] < firstGroups[0][i] {
  					swap(j, 0)
  				}
  			}
  		}
  		for i := range firstGroups[0] {
  			fmt.Print(firstGroups[0][i]+1, " ")
  		}
  	}
  }
*/

/*func flipSigns(arr []int) [][]int {
	n := len(arr)
	s := 0
	for _, x := range arr {
		s += x
	}

	// Create dp array
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, 2*s+1)
	}
	dp[0][s] = true

	// Compute dp array
	for i := 1; i < n; i++ {
		for j := 0; j <= 2*s; j++ {
			if j+arr[i] <= 2*s && dp[i-1][j+arr[i]] {
				dp[i][j] = true
			}
			if j-arr[i] >= 0 && dp[i-1][j-arr[i]] {
				dp[i][j] = true
			}
		}
	}

	// Find indices of flipped elements
	var indices [][]int
	for i := 0; i < n; i++ {
		if dp[i][s] {
			indices = append(indices, []int{})
			j := s
			for k := i; k >= 1; k-- {
				if dp[k-1][j+arr[k]] {
					j += arr[k]
					indices[len(indices)-1] = append(indices[len(indices)-1], k-1)
				} else if dp[k-1][j-arr[k]] {
					j -= arr[k]
					indices[len(indices)-1] = append(indices[len(indices)-1], k-1)
				} else {
					// This should not happen
					panic("Invalid dp array")
				}
			}
			if dp[0][j+arr[0]] {
				indices[len(indices)-1] = append(indices[len(indices)-1], 0)
			} else if dp[0][j-arr[0]] {
				indices[len(indices)-1] = append(indices[len(indices)-1], 0)
			} else {
				// This should not happen
				panic("Invalid dp array")
			}
			// Reverse the order of indices
			for j, k := 0, len(indices[len(indices)-1])-1; j < k; j, k = j+1, k-1 {
				indices[len(indices)-1][j], indices[len(indices)-1][k] = indices[len(indices)-1][k], indices[len(indices)-1][j]
			}
		}
	}

	return indices
}*/
