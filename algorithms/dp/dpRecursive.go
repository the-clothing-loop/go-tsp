// https://www.geeksforgeeks.org/travelling-salesman-problem-using-dynamic-programming/
package dp

func DPTspRecursive(graph [][]float64) (float64, []int) {
	n := len(graph) - 1

	memo := make([][]float64, n+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]float64, 1<<(n+1))
	}

	// calculates the cost of the most efficient tour starting at node i
	// and visiting all unvisited nodes represented by the mask
	var dptspRecursive func(i, mask int) (float64, []int)
	dptspRecursive = func(i, mask int) (float64, []int) {

		// Base case: If only the ith bit and 1st bit are set in the mask,
		// it implies that we have visited all other nodes already.
		// In this case, we return the distance between node 1 and node i
		// along with the path [1, i].
		if mask == (1<<i | 3) {
			return graph[1][i], []int{0, i - 1}
		}

		// Memoization: If the result for the current node and mask has been calculated before,
		// we return the stored result to avoid redundant calculations.
		if memo[i][mask] != 0 {
			return memo[i][mask], nil
		}

		res := inf // Result of this sub-problem
		var resPath []int

		// We iterate over all nodes j in the mask and calculate the cost
		// of traveling from node j to node i, taking the shortest path.
		// We choose the minimum cost among all possible j nodes.
		for j := 1; j <= n; j++ {
			if mask&(1<<j) != 0 && j != i && j != 1 {

				cost, currPath := dptspRecursive(j, mask&^(1<<i))
				cost += graph[j][i]
				if cost < res {
					res = cost
					resPath = currPath
				}
			}
		}

		memo[i][mask] = res // Store the result in the memoization table
		resPath = append(resPath, i-1)

		return res, resPath
	}

	ans, path := inf, []int{}

	for i := 1; i <= n; i++ {
		// We try to go from node 1, visiting all nodes in between, to node i,
		// and then return from i taking the shortest route back to node 1.
		// We choose the minimum cost among all possible i nodes.
		cost, newPath := dptspRecursive(i, (1<<(n+1))-1)
		cost += graph[i][1]
		if cost < ans {
			ans = cost
			path = newPath
		}
	}

	return ans, path
}
