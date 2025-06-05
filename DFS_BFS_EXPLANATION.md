# DFS and BFS Algorithms in Go

This document provides a comprehensive explanation of **Depth-First Search (DFS)** and **Breadth-First Search (BFS)** algorithms with practical Go implementations.

## Table of Contents
1. [Overview](#overview)
2. [Depth-First Search (DFS)](#depth-first-search-dfs)
3. [Breadth-First Search (BFS)](#breadth-first-search-bfs)
4. [Comparison](#comparison)
5. [When to Use Which](#when-to-use-which)
6. [Common Applications](#common-applications)
7. [Running the Examples](#running-the-examples)

## Overview

DFS and BFS are fundamental graph traversal algorithms used to visit all vertices in a graph. They form the foundation for many other algorithms and are essential for solving various computational problems.

## Depth-First Search (DFS)

### Concept
DFS explores as far as possible along each branch before backtracking. It uses a **stack** data structure (either explicit stack or recursion stack).

### Algorithm Steps:
1. Start from a given vertex
2. Mark it as visited
3. Visit an adjacent unvisited vertex
4. Recursively apply DFS to that vertex
5. If no unvisited adjacent vertices exist, backtrack

### Visual Example:
```
Graph:     0
          / \
         1   2
        /   / \
       3   4   5

DFS from 0: 0 → 1 → 3 → 2 → 4 → 5
```

### Time & Space Complexity:
- **Time Complexity**: O(V + E) where V = vertices, E = edges
- **Space Complexity**: O(V) for visited array + O(V) for recursion stack = O(V)

### Implementation Types:

#### 1. Recursive DFS
```go
func (g *Graph) DFS(start int) {
    visited := make(map[int]bool)
    g.dfsUtil(start, visited)
}

func (g *Graph) dfsUtil(vertex int, visited map[int]bool) {
    visited[vertex] = true
    fmt.Printf("%d ", vertex)
    
    for _, neighbor := range g.adjList[vertex] {
        if !visited[neighbor] {
            g.dfsUtil(neighbor, visited)
        }
    }
}
```

#### 2. Iterative DFS (using explicit stack)
```go
func (g *Graph) DFSIterative(start int) {
    visited := make(map[int]bool)
    stack := []int{start}
    
    for len(stack) > 0 {
        vertex := stack[len(stack)-1]  // Peek
        stack = stack[:len(stack)-1]   // Pop
        
        if !visited[vertex] {
            visited[vertex] = true
            fmt.Printf("%d ", vertex)
            
            // Add neighbors to stack (in reverse order)
            for i := len(g.adjList[vertex]) - 1; i >= 0; i-- {
                if !visited[g.adjList[vertex][i]] {
                    stack = append(stack, g.adjList[vertex][i])
                }
            }
        }
    }
}
```

### DFS for Trees:

#### Preorder (Root → Left → Right)
```go
func DFSPreorder(root *TreeNode) {
    if root == nil {
        return
    }
    fmt.Printf("%d ", root.Val)    // Process root
    DFSPreorder(root.Left)         // Traverse left
    DFSPreorder(root.Right)        // Traverse right
}
```

#### Inorder (Left → Root → Right)
```go
func DFSInorder(root *TreeNode) {
    if root == nil {
        return
    }
    DFSInorder(root.Left)          // Traverse left
    fmt.Printf("%d ", root.Val)    // Process root
    DFSInorder(root.Right)         // Traverse right
}
```

#### Postorder (Left → Right → Root)
```go
func DFSPostorder(root *TreeNode) {
    if root == nil {
        return
    }
    DFSPostorder(root.Left)        // Traverse left
    DFSPostorder(root.Right)       // Traverse right
    fmt.Printf("%d ", root.Val)    // Process root
}
```

## Breadth-First Search (BFS)

### Concept
BFS explores all vertices at the current depth before moving to vertices at the next depth level. It uses a **queue** data structure.

### Algorithm Steps:
1. Start from a given vertex and add it to queue
2. Mark it as visited
3. While queue is not empty:
   - Dequeue a vertex
   - Process it
   - Add all unvisited adjacent vertices to queue
   - Mark them as visited

### Visual Example:
```
Graph:     0
          / \
         1   2
        /   / \
       3   4   5

BFS from 0: 0 → 1 → 2 → 3 → 4 → 5
```

### Time & Space Complexity:
- **Time Complexity**: O(V + E) where V = vertices, E = edges
- **Space Complexity**: O(V) for visited array + O(V) for queue = O(V)

### Implementation:

#### Graph BFS
```go
func (g *Graph) BFS(start int) {
    visited := make(map[int]bool)
    queue := []int{start}
    visited[start] = true
    
    for len(queue) > 0 {
        vertex := queue[0]      // Front of queue
        queue = queue[1:]       // Dequeue
        fmt.Printf("%d ", vertex)
        
        // Add all unvisited neighbors
        for _, neighbor := range g.adjList[vertex] {
            if !visited[neighbor] {
                visited[neighbor] = true
                queue = append(queue, neighbor)
            }
        }
    }
}
```

#### Binary Tree Level Order Traversal
```go
func BFSLevelOrder(root *TreeNode) {
    if root == nil {
        return
    }
    
    queue := []*TreeNode{root}
    
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
}
```

## Comparison

| Aspect | DFS | BFS |
|--------|-----|-----|
| **Data Structure** | Stack (recursion/explicit) | Queue |
| **Memory Usage** | O(h) where h = height | O(w) where w = width |
| **Shortest Path** | No (unweighted graphs) | Yes (unweighted graphs) |
| **Implementation** | Recursive/Iterative | Iterative |
| **Space Efficiency** | Better for deep graphs | Better for wide graphs |
| **Use Cases** | Topological sort, cycle detection | Shortest path, level processing |

## When to Use Which

### Use DFS when:
- Finding a path (any path, not necessarily shortest)
- Detecting cycles in directed graphs
- Topological sorting
- Finding strongly connected components
- Solving maze problems
- Memory is limited (deep but narrow graphs)

### Use BFS when:
- Finding shortest path in unweighted graphs
- Level-order processing is needed
- Finding all nodes at a specific distance
- Finding connected components
- Social networking (degrees of separation)
- Web crawling (level by level)

## Common Applications

### DFS Applications:

#### 1. Cycle Detection in Directed Graph
```go
func (g *Graph) HasCycleDFS() bool {
    visited := make(map[int]bool)
    recStack := make(map[int]bool)  // Recursion stack
    
    for vertex := 0; vertex < g.vertices; vertex++ {
        if !visited[vertex] {
            if g.hasCycleDFSUtil(vertex, visited, recStack) {
                return true
            }
        }
    }
    return false
}
```

#### 2. Count Connected Components
```go
func (g *Graph) CountConnectedComponents() int {
    visited := make(map[int]bool)
    count := 0
    
    for vertex := 0; vertex < g.vertices; vertex++ {
        if !visited[vertex] {
            g.dfsUtil(vertex, visited)  // DFS from unvisited vertex
            count++
        }
    }
    return count
}
```

### BFS Applications:

#### 1. Shortest Path in Unweighted Graph
```go
func (g *Graph) BFSShortestPath(start, end int) int {
    if start == end {
        return 0
    }
    
    visited := make(map[int]bool)
    queue := [][]int{{start, 0}}  // [vertex, distance]
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
    
    return -1  // Path not found
}
```

#### 2. Level Order Processing
Perfect for problems that require processing nodes level by level, such as:
- Binary tree level order traversal
- Finding minimum depth of binary tree
- Printing binary tree in level order

## Running the Examples

To run the DFS and BFS examples:

1. **Compile and run the demonstration:**
```bash
go run dfs_bfs_examples.go
```

2. **Call the demo function from your main:**
```go
func main() {
    DemoDFSBFS()
}
```

This will show:
- Graph traversal using both DFS and BFS
- Binary tree traversals (DFS variants and BFS)
- Shortest path finding using BFS
- Cycle detection using DFS
- Connected components counting

## Key Takeaways

1. **DFS** goes deep first, uses stack, good for path finding and cycle detection
2. **BFS** goes wide first, uses queue, guarantees shortest path in unweighted graphs
3. Both have O(V + E) time complexity
4. Choice depends on the problem requirements and graph characteristics
5. Both are fundamental building blocks for more complex algorithms

## Practice Problems

Try implementing these algorithms for:
- Finding if a path exists between two nodes
- Detecting cycles in undirected graphs
- Finding the diameter of a tree
- Implementing topological sort
- Solving maze problems
- Finding bridges and articulation points in graphs

Remember: The best way to understand these algorithms is to implement them yourself and trace through the execution step by step! 