package dp

const (
	inf = float64(1e9) // A large positive value for unreachable nodes
)

// return the minimunCost to travel to all cities and the optimal path
func DPTspIterative(graph [][]float64) (float64, []int) {

	n := len(graph)

	// Initialize DP table
	// 1<<n: The expression 1 << n performs a bitwise left shift operation on 1 by n positions.
	// This effectively calculates the value of 2^n,
	// which represents the number of rows in the two-dimensional slice.
	// Example if you have 4 cities you will have 2^4 combinations (I supose)
	dp := make([][]float64, 1<<n)
	for i := range dp {
		dp[i] = make([]float64, n)
		for j := range dp[i] {
			// initialize all distance with a large number
			dp[i][j] = inf
		}
	}

	// Initialize parent table to store the path
	parent := make([][]int, 1<<n)
	for i := range parent {
		parent[i] = make([]int, n)
	}

	// Base case: starting point to itself has distance 0
	dp[1][0] = 0

	// this loop iterates through all possible subsets of cities.
	//The variable mask is used as a bitmask to represent each subset,
	//where each bit in mask corresponds to a city in the graph.
	//The loop starts from the subset with only the first city included
	//(represented by mask := 1) and continues until it reaches the subset with all cities
	//included (represented by mask < (1 << n)).
	for mask := 1; mask < (1 << n); mask++ {
		for last := 0; last < n; last++ {
			// Check if last visited city is in the current subset
			if (mask>>last)&1 == 1 {
				// Try all possible next cities
				for next := 0; next < n; next++ {
					// Check if next city is not visited yet
					if (mask>>next)&1 == 0 {
						// Update the minimum distance
						if dp[mask][last]+graph[last][next] < dp[mask|(1<<next)][next] {
							dp[mask|(1<<next)][next] = dp[mask][last] + graph[last][next]
							parent[mask|(1<<next)][next] = last
						}
					}
				}
			}
		}
	}

	// Find the minimum cost of visiting all cities and returning to the starting point
	minCost := inf
	finalMask := (1 << n) - 1
	last := 0
	for i := 1; i < n; i++ {
		if dp[finalMask][i]+graph[i][0] < minCost {
			minCost = dp[finalMask][i] + graph[i][0]
			last = i
		}
	}

	// Reconstruct the path
	path := make([]int, n+1)
	path[n] = 0 // Starting point
	mask := finalMask
	for i := n - 1; i > 0; i-- {
		path[i] = last
		prev := parent[mask][last]
		mask ^= (1 << last)
		last = prev
	}

	return minCost, path
}
