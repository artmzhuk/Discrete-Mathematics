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

func main() {
	sol := findSolutions(inputGraph())
	sel(sol)
	/*	if solutions != nil {
		selectSolution(solutions)
	}*/
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
