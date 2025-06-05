# Topological Sort Algorithm in Go

## Overview

**Topological Sort** is a linear ordering of vertices in a **Directed Acyclic Graph (DAG)** such that for every directed edge (u,v), vertex u comes before vertex v in the ordering.

### Key Requirements:
1. **Must be a DAG** (Directed Acyclic Graph)
2. **Not unique** - multiple valid orderings may exist
3. **Cycle detection** - if cycle exists, topological sort is impossible

## Why Use Topological Sort?

Topological sort solves dependency problems where:
- Tasks have prerequisites
- Order of execution matters
- Dependencies form a directed graph

### Real-World Examples:
- **Course Prerequisites**: Take Math before Physics
- **Build Systems**: Compile dependencies before main program
- **Task Scheduling**: Complete design before coding
- **Package Management**: Install dependencies before main package

## Algorithms

### 1. DFS-Based Approach

**Core Idea**: Use DFS and process vertices in post-order (after visiting all descendants).

#### Algorithm Steps:
1. Perform DFS on all unvisited vertices
2. When finishing a vertex (post-order), push it to stack
3. Pop all vertices from stack to get topological order

#### Implementation:
```go
func (g *DirectedGraph) TopologicalSortDFS() []int {
    visited := make(map[int]bool)
    stack := []int{}
    
    // DFS on all unvisited vertices
    for vertex := 0; vertex < g.vertices; vertex++ {
        if !visited[vertex] {
            g.dfsUtil(vertex, visited, &stack)
        }
    }
    
    // Reverse stack for topological order
    result := make([]int, len(stack))
    for i := 0; i < len(stack); i++ {
        result[i] = stack[len(stack)-1-i]
    }
    return result
}

func (g *DirectedGraph) dfsUtil(vertex int, visited map[int]bool, stack *[]int) {
    visited[vertex] = true
    
    // Visit all neighbors first
    for _, neighbor := range g.adjList[vertex] {
        if !visited[neighbor] {
            g.dfsUtil(neighbor, visited, stack)
        }
    }
    
    // Add to stack after visiting all descendants (post-order)
    *stack = append(*stack, vertex)
}
```

**Complexity**: 
- Time: O(V + E)
- Space: O(V) for recursion stack

### 2. Kahn's Algorithm (BFS-Based)

**Core Idea**: Repeatedly remove vertices with no incoming edges (in-degree = 0).

#### Algorithm Steps:
1. Calculate in-degree for all vertices
2. Add all vertices with in-degree 0 to queue
3. While queue is not empty:
   - Remove vertex from queue, add to result
   - Decrease in-degree of all neighbors
   - If neighbor's in-degree becomes 0, add to queue
4. If result contains all vertices, return it; otherwise, cycle detected

#### Implementation:
```go
func (g *DirectedGraph) TopologicalSortKahn() []int {
    // Calculate in-degrees
    inDegree := make([]int, g.vertices)
    for vertex := 0; vertex < g.vertices; vertex++ {
        for _, neighbor := range g.adjList[vertex] {
            inDegree[neighbor]++
        }
    }
    
    // Find vertices with 0 in-degree
    queue := []int{}
    for vertex := 0; vertex < g.vertices; vertex++ {
        if inDegree[vertex] == 0 {
            queue = append(queue, vertex)
        }
    }
    
    result := []int{}
    
    // Process vertices with 0 in-degree
    for len(queue) > 0 {
        vertex := queue[0]
        queue = queue[1:]
        result = append(result, vertex)
        
        // Reduce in-degree of neighbors
        for _, neighbor := range g.adjList[vertex] {
            inDegree[neighbor]--
            if inDegree[neighbor] == 0 {
                queue = append(queue, neighbor)
            }
        }
    }
    
    // Check for cycle
    if len(result) != g.vertices {
        return nil // Cycle detected
    }
    
    return result
}
```

**Complexity**: 
- Time: O(V + E)
- Space: O(V) for queue and in-degree array

## Algorithm Comparison

| Aspect | DFS-Based | Kahn's Algorithm |
|--------|-----------|------------------|
| **Approach** | Recursive, post-order | Iterative, BFS-like |
| **Data Structure** | Stack (recursion) | Queue |
| **Cycle Detection** | Separate function needed | Built-in |
| **Intuition** | Process dependencies first | Remove independent tasks first |
| **Implementation** | More elegant | More intuitive |
| **Space Usage** | Recursion stack | Queue + in-degree array |

## Practical Examples

### Example 1: Course Scheduling

```
Prerequisites:
- Math → Physics
- Math → Chemistry  
- Physics → Advanced Physics
- Chemistry → Biology

Possible topological order: [Math, Physics, Chemistry, Advanced Physics, Biology]
```

### Example 2: Build Dependencies

```
Dependencies:
- utils.go → main.go
- config.go → main.go
- database.go → main.go
- utils.go → database.go

Build order: [utils.go, config.go, database.go, main.go]
```

### Example 3: Task Scheduling

```
Task Dependencies:
- Setup → Design
- Design → Code
- Code → Test
- Test → Deploy
- Design → Documentation

Execution order: [Setup, Design, Code, Documentation, Test, Deploy]
```

## Cycle Detection

Topological sort is only possible for DAGs. If a cycle exists, no valid ordering exists.

### DFS-Based Cycle Detection:
```go
func (g *DirectedGraph) HasCycle() bool {
    visited := make(map[int]bool)
    recStack := make(map[int]bool)  // Recursion stack
    
    for vertex := 0; vertex < g.vertices; vertex++ {
        if !visited[vertex] {
            if g.hasCycleUtil(vertex, visited, recStack) {
                return true
            }
        }
    }
    return false
}
```

### Kahn's Algorithm Cycle Detection:
If the result doesn't contain all vertices, a cycle exists.

## Visual Example

Consider this DAG:
```
    5 ──→ 2 ──→ 3
    ↓           ↓
    0     4 ──→ 1
          ↓     ↗
          └─────┘
```

**DFS Result**: [5, 4, 2, 3, 1, 0]
**Kahn's Result**: [4, 5, 0, 2, 3, 1]

Both are valid topological orderings!

## When to Use Each Algorithm

### Use DFS-Based When:
- You need to detect cycles separately
- Recursive solution is preferred
- Working with sparse graphs
- Need to find strongly connected components

### Use Kahn's Algorithm When:
- You want built-in cycle detection
- Iterative solution is preferred
- Working with dense graphs
- Need to process by dependency levels

## Common Applications

### 1. **Build Systems**
- Makefile dependencies
- Package compilation order
- Webpack module bundling

### 2. **Scheduling**
- Project task ordering
- CPU instruction scheduling
- Database transaction ordering

### 3. **Academic Planning**
- Course prerequisite chains
- Skill development paths
- Learning roadmaps

### 4. **Software Dependencies**
- Package managers (npm, pip, go mod)
- Library loading order
- Plugin initialization

### 5. **Data Processing**
- ETL pipeline stages
- MapReduce job ordering
- Machine learning workflow

## Key Insights

1. **Multiple Solutions**: Usually multiple valid topological orderings exist
2. **Cycle Detection**: Essential prerequisite - no cycles allowed
3. **Applications**: Ubiquitous in computer science and real-world systems
4. **Performance**: Both algorithms are linear time O(V + E)
5. **Choice**: Pick based on whether you need built-in cycle detection (Kahn's) or prefer recursive elegance (DFS)

## Practice Problems

Try implementing topological sort for:
1. **Course Schedule** (LeetCode 207, 210)
2. **Build System** with file dependencies
3. **Task Scheduler** with prerequisites
4. **Package Manager** dependency resolver
5. **Compilation Order** for modules

The key is recognizing when a problem has dependency relationships that form a DAG! 