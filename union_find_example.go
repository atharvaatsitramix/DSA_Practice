package main

import (
	"fmt"
	"sort"
)

// ================================
// UNION-FIND (DISJOINT SET UNION)
// ================================

// UnionFind represents a Union-Find data structure
type UnionFind struct {
	parent []int // parent[i] = parent of element i
	rank   []int // rank[i] = approximate depth of tree rooted at i
	count  int   // number of disjoint sets
}

// NewUnionFind creates a new Union-Find data structure with n elements
func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)

	// Initially, each element is its own parent (self-loop)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}

	return &UnionFind{
		parent: parent,
		rank:   rank,
		count:  n, // Initially n disjoint sets
	}
}

// Find returns the root of the set containing x
// Uses path compression optimization
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		// Path compression: make parent[x] point directly to root
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// Union merges the sets containing x and y
// Uses union by rank optimization
func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	// Already in same set
	if rootX == rootY {
		return false
	}

	// Union by rank: attach smaller tree under larger tree
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		// Same rank: make one root and increase its rank
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}

	uf.count--
	return true
}

// Connected checks if x and y are in the same set
func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

// Count returns the number of disjoint sets
func (uf *UnionFind) Count() int {
	return uf.count
}

// GetComponents returns all elements grouped by their components
func (uf *UnionFind) GetComponents() map[int][]int {
	components := make(map[int][]int)

	for i := 0; i < len(uf.parent); i++ {
		root := uf.Find(i)
		components[root] = append(components[root], i)
	}

	return components
}

// ================================
// OPTIMIZED VERSIONS
// ================================

// WeightedUnionFind uses union by size instead of rank
type WeightedUnionFind struct {
	parent []int
	size   []int // size[i] = size of tree rooted at i
	count  int
}

func NewWeightedUnionFind(n int) *WeightedUnionFind {
	parent := make([]int, n)
	size := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}

	return &WeightedUnionFind{
		parent: parent,
		size:   size,
		count:  n,
	}
}

func (wuf *WeightedUnionFind) Find(x int) int {
	if wuf.parent[x] != x {
		wuf.parent[x] = wuf.Find(wuf.parent[x])
	}
	return wuf.parent[x]
}

func (wuf *WeightedUnionFind) Union(x, y int) bool {
	rootX := wuf.Find(x)
	rootY := wuf.Find(y)

	if rootX == rootY {
		return false
	}

	// Union by size: attach smaller tree to larger tree
	if wuf.size[rootX] < wuf.size[rootY] {
		wuf.parent[rootX] = rootY
		wuf.size[rootY] += wuf.size[rootX]
	} else {
		wuf.parent[rootY] = rootX
		wuf.size[rootX] += wuf.size[rootY]
	}

	wuf.count--
	return true
}

func (wuf *WeightedUnionFind) Connected(x, y int) bool {
	return wuf.Find(x) == wuf.Find(y)
}

func (wuf *WeightedUnionFind) GetSize(x int) int {
	return wuf.size[wuf.Find(x)]
}

// ================================
// PRACTICAL APPLICATIONS
// ================================

// Edge represents an edge in a graph
type Edge struct {
	From   int
	To     int
	Weight int
}

// NumberOfIslands counts connected components in a 2D grid
func NumberOfIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	uf := NewUnionFind(rows * cols)
	islands := 0

	// Count initial islands (1s)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				islands++
			}
		}
	}

	// Directions: up, down, left, right
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				// Check all 4 directions
				for _, dir := range directions {
					ni, nj := i+dir[0], j+dir[1]
					if ni >= 0 && ni < rows && nj >= 0 && nj < cols && grid[ni][nj] == '1' {
						id1 := i*cols + j
						id2 := ni*cols + nj
						if uf.Union(id1, id2) {
							islands-- // Two islands merged into one
						}
					}
				}
			}
		}
	}

	return islands
}

// KruskalMST finds Minimum Spanning Tree using Kruskal's algorithm
func KruskalMST(n int, edges []Edge) ([]Edge, int) {
	// Sort edges by weight
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	uf := NewUnionFind(n)
	mst := []Edge{}
	totalWeight := 0

	for _, edge := range edges {
		// If adding this edge doesn't create a cycle
		if uf.Union(edge.From, edge.To) {
			mst = append(mst, edge)
			totalWeight += edge.Weight

			// MST has exactly n-1 edges
			if len(mst) == n-1 {
				break
			}
		}
	}

	return mst, totalWeight
}

// DetectCycle detects if an undirected graph has a cycle
func DetectCycle(n int, edges []Edge) bool {
	uf := NewUnionFind(n)

	for _, edge := range edges {
		// If both vertices are already connected, adding this edge creates a cycle
		if uf.Connected(edge.From, edge.To) {
			return true
		}
		uf.Union(edge.From, edge.To)
	}

	return false
}

// FriendCircles counts the number of friend circles
func FriendCircles(friends [][]int) int {
	n := len(friends)
	uf := NewUnionFind(n)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if friends[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}

	return uf.Count()
}

// AccountsMerge merges accounts belonging to the same person
func AccountsMerge(accounts [][]string) [][]string {
	// Map email to account index
	emailToIndex := make(map[string]int)
	emailToName := make(map[string]string)
	index := 0

	// Assign unique index to each email
	for _, account := range accounts {
		name := account[0]
		for i := 1; i < len(account); i++ {
			email := account[i]
			if _, exists := emailToIndex[email]; !exists {
				emailToIndex[email] = index
				emailToName[email] = name
				index++
			}
		}
	}

	uf := NewUnionFind(index)

	// Union emails in the same account
	for _, account := range accounts {
		if len(account) > 1 {
			firstEmailIndex := emailToIndex[account[1]]
			for i := 2; i < len(account); i++ {
				uf.Union(firstEmailIndex, emailToIndex[account[i]])
			}
		}
	}

	// Group emails by their root
	indexToEmails := make(map[int][]string)
	for email, idx := range emailToIndex {
		root := uf.Find(idx)
		indexToEmails[root] = append(indexToEmails[root], email)
	}

	// Build result
	result := [][]string{}
	for _, emails := range indexToEmails {
		sort.Strings(emails)
		name := emailToName[emails[0]]
		account := append([]string{name}, emails...)
		result = append(result, account)
	}

	return result
}

// ================================
// DEMONSTRATION FUNCTIONS
// ================================

func DemoUnionFind() {
	fmt.Println("=== UNION-FIND (DISJOINT SET UNION) ALGORITHM ===\n")

	fmt.Println("Union-Find is a data structure that efficiently handles:")
	fmt.Println("1. Union: Merge two disjoint sets")
	fmt.Println("2. Find: Determine which set an element belongs to")
	fmt.Println("3. Connected: Check if two elements are in the same set\n")

	// Example 1: Basic operations
	fmt.Println("=== EXAMPLE 1: Basic Operations ===")
	uf := NewUnionFind(7)
	fmt.Printf("Initial sets: %d disjoint sets {0}, {1}, {2}, {3}, {4}, {5}, {6}\n", uf.Count())

	// Union operations
	operations := [][]int{{0, 1}, {2, 3}, {4, 5}, {1, 3}, {5, 6}}

	for _, op := range operations {
		x, y := op[0], op[1]
		fmt.Printf("Union(%d, %d): ", x, y)
		if uf.Union(x, y) {
			fmt.Printf("Merged! Now %d sets\n", uf.Count())
		} else {
			fmt.Printf("Already connected! Still %d sets\n", uf.Count())
		}
	}

	// Test connectivity
	fmt.Println("\nConnectivity tests:")
	testPairs := [][]int{{0, 3}, {4, 6}, {0, 4}, {2, 1}}
	for _, pair := range testPairs {
		x, y := pair[0], pair[1]
		fmt.Printf("Connected(%d, %d): %v\n", x, y, uf.Connected(x, y))
	}

	// Show final components
	fmt.Println("\nFinal components:")
	components := uf.GetComponents()
	for root, members := range components {
		fmt.Printf("Component %d: %v\n", root, members)
	}
	fmt.Println()

	// Example 2: Islands problem
	fmt.Println("=== EXAMPLE 2: Number of Islands ===")
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	fmt.Println("Grid:")
	for _, row := range grid {
		fmt.Printf("%s\n", string(row))
	}

	islands := NumberOfIslands(grid)
	fmt.Printf("Number of islands: %d\n\n", islands)

	// Example 3: Minimum Spanning Tree
	fmt.Println("=== EXAMPLE 3: Minimum Spanning Tree (Kruskal's Algorithm) ===")
	edges := []Edge{
		{0, 1, 4}, {0, 7, 8}, {1, 2, 8}, {1, 7, 11},
		{2, 3, 7}, {2, 8, 2}, {2, 5, 4}, {3, 4, 9},
		{3, 5, 14}, {4, 5, 10}, {5, 6, 2}, {6, 7, 1},
		{6, 8, 6}, {7, 8, 7},
	}

	fmt.Println("Edges (from, to, weight):")
	for _, e := range edges {
		fmt.Printf("(%d, %d, %d) ", e.From, e.To, e.Weight)
	}
	fmt.Println()

	mst, totalWeight := KruskalMST(9, edges)
	fmt.Printf("\nMinimum Spanning Tree (weight = %d):\n", totalWeight)
	for _, edge := range mst {
		fmt.Printf("(%d, %d, %d) ", edge.From, edge.To, edge.Weight)
	}
	fmt.Println("\n")

	// Example 4: Cycle detection
	fmt.Println("=== EXAMPLE 4: Cycle Detection ===")
	cyclicEdges := []Edge{{0, 1, 1}, {1, 2, 1}, {2, 0, 1}}
	acyclicEdges := []Edge{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}}

	fmt.Printf("Cyclic graph edges: ")
	for _, e := range cyclicEdges {
		fmt.Printf("(%d, %d) ", e.From, e.To)
	}
	fmt.Printf("Has cycle: %v\n", DetectCycle(3, cyclicEdges))

	fmt.Printf("Acyclic graph edges: ")
	for _, e := range acyclicEdges {
		fmt.Printf("(%d, %d) ", e.From, e.To)
	}
	fmt.Printf("Has cycle: %v\n\n", DetectCycle(4, acyclicEdges))

	// Example 5: Friend circles
	fmt.Println("=== EXAMPLE 5: Friend Circles ===")
	friends := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}

	fmt.Println("Friendship matrix:")
	for i, row := range friends {
		fmt.Printf("Person %d: %v\n", i, row)
	}

	circles := FriendCircles(friends)
	fmt.Printf("Number of friend circles: %d\n\n", circles)

	// Performance characteristics
	fmt.Println("=== ALGORITHM CHARACTERISTICS ===")
	fmt.Println("Time Complexity (with optimizations):")
	fmt.Println("- Find: O(α(n)) ≈ O(1) amortized")
	fmt.Println("- Union: O(α(n)) ≈ O(1) amortized")
	fmt.Println("- α(n) is the inverse Ackermann function (grows very slowly)")
	fmt.Println()
	fmt.Println("Space Complexity: O(n)")
	fmt.Println()
	fmt.Println("Key Optimizations:")
	fmt.Println("1. Path Compression: Make nodes point directly to root during Find")
	fmt.Println("2. Union by Rank/Size: Attach smaller tree to larger tree")
	fmt.Println()
	fmt.Println("Applications:")
	fmt.Println("- Network connectivity")
	fmt.Println("- Image processing (connected components)")
	fmt.Println("- Kruskal's MST algorithm")
	fmt.Println("- Percolation theory")
	fmt.Println("- Social network analysis")
	fmt.Println("- Dynamic connectivity problems")
}

// DemoAdvancedApplications shows more complex use cases
func DemoAdvancedApplications() {
	fmt.Println("\n=== ADVANCED APPLICATIONS ===")

	// Application 1: Account merging
	fmt.Println("1. ACCOUNT MERGING")
	accounts := [][]string{
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"Mary", "mary@mail.com"},
		{"John", "johnnybravo@mail.com"},
	}

	fmt.Println("Original accounts:")
	for i, account := range accounts {
		fmt.Printf("Account %d: %v\n", i, account)
	}

	merged := AccountsMerge(accounts)
	fmt.Println("\nMerged accounts:")
	for i, account := range merged {
		fmt.Printf("Merged %d: %v\n", i, account)
	}
	fmt.Println()

	// Application 2: Weighted Union-Find for size tracking
	fmt.Println("2. COMPONENT SIZE TRACKING")
	wuf := NewWeightedUnionFind(8)

	connections := [][]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {0, 2}, {4, 6}}

	for _, conn := range connections {
		x, y := conn[0], conn[1]
		fmt.Printf("Connect %d and %d\n", x, y)
		wuf.Union(x, y)

		// Show component sizes
		components := make(map[int][]int)
		for i := 0; i < 8; i++ {
			root := wuf.Find(i)
			components[root] = append(components[root], i)
		}

		for root, members := range components {
			if len(members) > 1 {
				fmt.Printf("  Component %d: %v (size: %d)\n", root, members, wuf.GetSize(root))
			}
		}
		fmt.Println()
	}

	// Application 3: Dynamic connectivity with operations trace
	fmt.Println("3. DYNAMIC CONNECTIVITY TRACE")
	uf2 := NewUnionFind(6)

	fmt.Printf("Initial: %d components\n", uf2.Count())

	operations2 := []string{
		"union(0,1)", "union(2,3)", "union(4,5)",
		"connected(0,3)?", "union(1,2)", "connected(0,3)?",
		"union(3,4)", "count",
	}

	opData := [][]int{{0, 1}, {2, 3}, {4, 5}, {0, 3}, {1, 2}, {0, 3}, {3, 4}, {}}

	for i, op := range operations2 {
		fmt.Printf("%s: ", op)

		if i < 3 || i == 4 || i == 6 { // Union operations
			x, y := opData[i][0], opData[i][1]
			uf2.Union(x, y)
			fmt.Printf("Done. Components: %d\n", uf2.Count())
		} else if i == 3 || i == 5 { // Connected queries
			x, y := opData[i][0], opData[i][1]
			fmt.Printf("%v\n", uf2.Connected(x, y))
		} else if i == 7 { // Count
			fmt.Printf("%d components\n", uf2.Count())
		}
	}
}
