package dsa

import "fmt"

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	// Get the gray code for n-1
	prev := grayCode(n - 1)
	result := make([]int, len(prev)*2)

	// Copy the previous sequence
	for i := 0; i < len(prev); i++ {
		result[i] = prev[i]
	}

	// Add the reversed sequence with 1 prepended (equivalent to adding 2^(n-1))
	power := 1 << (n - 1)
	for i := 0; i < len(prev); i++ {
		result[len(prev)+i] = prev[len(prev)-1-i] + power
	}

	return result
}

func sumOfDistancesInTree(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	// Build adjacency list
	adj := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// Arrays to store subtree sizes and distances
	count := make([]int, n) // count[i] = number of nodes in subtree rooted at i
	res := make([]int, n)   // res[i] = sum of distances from node i to all other nodes

	// First DFS: calculate count and initial distances from root (node 0)
	var dfs1 func(node, parent int) int
	dfs1 = func(node, parent int) int {
		count[node] = 1
		for _, child := range adj[node] {
			if child != parent {
				count[node] += dfs1(child, node)
				res[node] += res[child] + count[child]
			}
		}
		return count[node]
	}

	// Second DFS: calculate distances for all nodes using the re-rooting technique
	var dfs2 func(node, parent int)
	dfs2 = func(node, parent int) {
		for _, child := range adj[node] {
			if child != parent {
				res[child] = res[node] - count[child] + (n - count[child])
				dfs2(child, node)
			}
		}
	}

	// Start DFS from root (node 0)
	dfs1(0, -1)
	dfs2(0, -1)

	return res
}

func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)

	// dp[i][j] represents the length of the longest common subarray ending at nums1[i-1] and nums2[j-1]
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	maxLen := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
				}
			}
		}
	}

	return maxLen
}

func RunDSATest(algorithm string) {
	switch algorithm {
	case "grayCode":
		runGrayCodeTests()
	case "sumOfDistancesInTree":
		runSumOfDistancesInTreeTests()
	case "findLength":
		runFindLengthTests()
	default:
		fmt.Printf("Unknown algorithm: %s. Available algorithms: grayCode, sumOfDistancesInTree, findLength\n", algorithm)
	}
}

func runGrayCodeTests() {
	fmt.Println("=== Gray Code Tests ===")
	testCases := []int{1, 2, 3, 4}

	for _, n := range testCases {
		result := grayCode(n)
		fmt.Printf("n = %d, Result: %v\n", n, result)
		fmt.Println()
	}
}

func runSumOfDistancesInTreeTests() {
	fmt.Println("=== Sum of Distances in Tree Tests ===")
	// Test cases
	testCases := []struct {
		n     int
		edges [][]int
	}{
		{
			n:     6,
			edges: [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}},
		},
		{
			n:     1,
			edges: [][]int{},
		},
		{
			n:     2,
			edges: [][]int{{1, 0}},
		},
		{
			n:     4,
			edges: [][]int{{1, 2}, {0, 1}, {0, 3}},
		},
	}

	for i, tc := range testCases {
		result := sumOfDistancesInTree(tc.n, tc.edges)
		fmt.Printf("Test case %d:\n", i+1)
		fmt.Printf("n = %d, edges = %v\n", tc.n, tc.edges)
		fmt.Printf("Result: %v\n", result)
		fmt.Println()
	}
}

func runFindLengthTests() {
	fmt.Println("=== Find Length Tests ===")
	testCases := []struct {
		nums1 []int
		nums2 []int
	}{
		{
			nums1: []int{1, 2, 3, 2, 1},
			nums2: []int{3, 2, 1, 4, 7},
		},
		{
			nums1: []int{0, 0, 0, 0, 0},
			nums2: []int{0, 0, 0, 0, 0},
		},
		{
			nums1: []int{1, 2, 3},
			nums2: []int{4, 5, 6},
		},
		{
			nums1: []int{1, 2, 3, 4, 5},
			nums2: []int{3, 4, 5, 6, 7},
		},
	}

	for i, tc := range testCases {
		result := findLength(tc.nums1, tc.nums2)
		fmt.Printf("Test case %d:\n", i+1)
		fmt.Printf("nums1 = %v, nums2 = %v\n", tc.nums1, tc.nums2)
		fmt.Printf("Result: %d\n", result)
		fmt.Println()
	}
}
