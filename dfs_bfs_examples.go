package main

import (
	"fmt"
)

// Graph represents an adjacency list graph
type Graph struct {
	vertices int
	adjList  map[int][]int
}

// NewGraph creates a new graph with given number of vertices
func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		adjList:  make(map[int][]int),
	}
}

// AddEdge adds an edge between two vertices (undirected graph)
func (g *Graph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u) // For undirected graph
}

// AddDirectedEdge adds a directed edge from u to v
func (g *Graph) AddDirectedEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
}

// TreeNode represents a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ================================
// DEPTH-FIRST SEARCH (DFS)
// ================================

// DFS traverses the graph using depth-first search
// Time Complexity: O(V + E) where V = vertices, E = edges
// Space Complexity: O(V) for visited array and recursion stack
func (g *Graph) DFS(start int) {
	visited := make(map[int]bool)
	fmt.Print("DFS Traversal: ")
	g.dfsUtil(start, visited)
	fmt.Println()
}

// dfsUtil is a recursive utility function for DFS
func (g *Graph) dfsUtil(vertex int, visited map[int]bool) {
	// Mark current vertex as visited and print it
	visited[vertex] = true
	fmt.Printf("%d ", vertex)

	// Recursively visit all adjacent vertices
	for _, neighbor := range g.adjList[vertex] {
		if !visited[neighbor] {
			g.dfsUtil(neighbor, visited)
		}
	}
}

// DFSIterative performs DFS using a stack (iterative approach)
func (g *Graph) DFSIterative(start int) {
	visited := make(map[int]bool)
	stack := []int{start}

	fmt.Print("DFS Iterative: ")

	for len(stack) > 0 {
		// Pop from stack
		vertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[vertex] {
			visited[vertex] = true
			fmt.Printf("%d ", vertex)

			// Add all unvisited neighbors to stack
			// Add in reverse order to maintain left-to-right traversal
			neighbors := g.adjList[vertex]
			for i := len(neighbors) - 1; i >= 0; i-- {
				if !visited[neighbors[i]] {
					stack = append(stack, neighbors[i])
				}
			}
		}
	}
	fmt.Println()
}

// DFS for Binary Tree - Preorder (Root -> Left -> Right)
func DFSPreorder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	DFSPreorder(root.Left)
	DFSPreorder(root.Right)
}

// DFS for Binary Tree - Inorder (Left -> Root -> Right)
func DFSInorder(root *TreeNode) {
	if root == nil {
		return
	}
	DFSInorder(root.Left)
	fmt.Printf("%d ", root.Val)
	DFSInorder(root.Right)
}

// DFS for Binary Tree - Postorder (Left -> Right -> Root)
func DFSPostorder(root *TreeNode) {
	if root == nil {
		return
	}
	DFSPostorder(root.Left)
	DFSPostorder(root.Right)
	fmt.Printf("%d ", root.Val)
}

// ================================
// BREADTH-FIRST SEARCH (BFS)
// ================================

// BFS traverses the graph using breadth-first search
// Time Complexity: O(V + E) where V = vertices, E = edges
// Space Complexity: O(V) for visited array and queue
func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
	queue := []int{start}
	visited[start] = true

	fmt.Print("BFS Traversal: ")

	for len(queue) > 0 {
		// Dequeue
		vertex := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", vertex)

		// Add all unvisited neighbors to queue
		for _, neighbor := range g.adjList[vertex] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	fmt.Println()
}

// BFS for Binary Tree - Level Order Traversal
func BFSLevelOrder(root *TreeNode) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}
	fmt.Print("BFS Level Order: ")

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Printf("%d ", node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	fmt.Println()
}

// BFS to find shortest path (unweighted graph)
func (g *Graph) BFSShortestPath(start, end int) int {
	if start == end {
		return 0
	}

	visited := make(map[int]bool)
	queue := [][]int{{start, 0}} // [vertex, distance]
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		vertex, distance := current[0], current[1]

		for _, neighbor := range g.adjList[vertex] {
			if neighbor == end {
				return distance + 1
			}

			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, []int{neighbor, distance + 1})
			}
		}
	}

	return -1 // Path not found
}

// ================================
// EXAMPLE APPLICATIONS
// ================================

// Check if graph has a cycle using DFS
func (g *Graph) HasCycleDFS() bool {
	visited := make(map[int]bool)
	recStack := make(map[int]bool)

	for vertex := 0; vertex < g.vertices; vertex++ {
		if !visited[vertex] {
			if g.hasCycleDFSUtil(vertex, visited, recStack) {
				return true
			}
		}
	}
	return false
}

func (g *Graph) hasCycleDFSUtil(vertex int, visited, recStack map[int]bool) bool {
	visited[vertex] = true
	recStack[vertex] = true

	for _, neighbor := range g.adjList[vertex] {
		if !visited[neighbor] && g.hasCycleDFSUtil(neighbor, visited, recStack) {
			return true
		} else if recStack[neighbor] {
			return true
		}
	}

	recStack[vertex] = false
	return false
}

// Count connected components using DFS
func (g *Graph) CountConnectedComponents() int {
	visited := make(map[int]bool)
	count := 0

	for vertex := 0; vertex < g.vertices; vertex++ {
		if !visited[vertex] {
			g.dfsUtil(vertex, visited)
			count++
		}
	}
	return count
}

// ================================
// DEMO FUNCTION WITH EXAMPLES
// ================================

func DemoDFSBFS() {
	fmt.Println("=== DFS and BFS Algorithms in Go ===\n")

	// Create a sample graph
	// Graph structure:
	//     0
	//    / \
	//   1   2
	//  /   / \
	// 3   4   5
	//         |
	//         6

	graph := NewGraph(7)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(5, 6)

	fmt.Println("Graph Adjacency List:")
	for vertex := 0; vertex < graph.vertices; vertex++ {
		fmt.Printf("Vertex %d: %v\n", vertex, graph.adjList[vertex])
	}
	fmt.Println()

	// DFS Examples
	fmt.Println("=== DEPTH-FIRST SEARCH (DFS) ===")
	graph.DFS(0)
	graph.DFSIterative(0)
	fmt.Println()

	// BFS Examples
	fmt.Println("=== BREADTH-FIRST SEARCH (BFS) ===")
	graph.BFS(0)
	fmt.Println()

	// Shortest path using BFS
	fmt.Println("=== BFS Shortest Path ===")
	distance := graph.BFSShortestPath(0, 6)
	fmt.Printf("Shortest path from 0 to 6: %d edges\n\n", distance)

	// Binary Tree Examples
	fmt.Println("=== BINARY TREE TRAVERSALS ===")

	// Create a binary tree:
	//       1
	//      / \
	//     2   3
	//    / \   \
	//   4   5   6
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	fmt.Print("DFS Preorder:  ")
	DFSPreorder(root)
	fmt.Println()

	fmt.Print("DFS Inorder:   ")
	DFSInorder(root)
	fmt.Println()

	fmt.Print("DFS Postorder: ")
	DFSPostorder(root)
	fmt.Println()

	BFSLevelOrder(root)
	fmt.Println()

	// Advanced Examples
	fmt.Println("=== ADVANCED APPLICATIONS ===")

	// Create a directed graph to test cycle detection
	directedGraph := NewGraph(4)
	directedGraph.AddDirectedEdge(0, 1)
	directedGraph.AddDirectedEdge(1, 2)
	directedGraph.AddDirectedEdge(2, 3)
	directedGraph.AddDirectedEdge(3, 1) // Creates a cycle

	fmt.Printf("Directed graph has cycle: %v\n", directedGraph.HasCycleDFS())

	// Count connected components
	disconnectedGraph := NewGraph(6)
	disconnectedGraph.AddEdge(0, 1)
	disconnectedGraph.AddEdge(2, 3)
	// Vertex 4 and 5 are isolated

	fmt.Printf("Connected components: %d\n", disconnectedGraph.CountConnectedComponents())
}
