# Union-Find (Disjoint Set Union) Algorithm in Go

## Overview

**Union-Find** (also called **Disjoint Set Union** or **DSU**) is a data structure that efficiently handles **dynamic connectivity** problems. It supports two main operations:

1. **Union**: Merge two disjoint sets
2. **Find**: Determine which set an element belongs to

### Key Characteristics:
- **Time Complexity**: O(α(n)) ≈ O(1) amortized for both operations (with optimizations)
- **Space Complexity**: O(n)
- **α(n)**: Inverse Ackermann function (grows extremely slowly)

## Core Problem: Dynamic Connectivity

Given a set of N objects, we want to:
- **Connect** two objects
- **Check** if two objects are connected
- **Count** number of connected components

### Example Scenario:
```
Objects: {0, 1, 2, 3, 4, 5}
Operations:
- union(0, 1) → Connect 0 and 1
- union(2, 3) → Connect 2 and 3  
- connected(0, 1)? → true
- connected(0, 2)? → false
- union(1, 3) → Connect components {0,1} and {2,3}
- connected(0, 2)? → true (now in same component)
```

## How Union-Find Works

### Core Idea:
- Each **set** is represented as a **tree**
- Each element points to its **parent**
- **Root** of tree represents the set
- Elements with same root belong to same set

### Basic Structure:
```
Initial state: Each element is its own parent
parent[0] = 0, parent[1] = 1, parent[2] = 2, ...

After union(0, 1):
parent[1] = 0  (1 points to 0)
Tree: 0 ← 1

After union(2, 3):
parent[3] = 2  (3 points to 2)
Trees: 0 ← 1    2 ← 3

After union(1, 3):
parent[2] = 0  (attach tree rooted at 2 to tree rooted at 0)
Tree: 0 ← 1
      ↑
      2 ← 3
```

## Step-by-Step Example

Let's trace through connecting elements {0, 1, 2, 3, 4, 5}:

### Initial State:
```
parent: [0, 1, 2, 3, 4, 5]
Sets: {0}, {1}, {2}, {3}, {4}, {5}
Count: 6
```

### Step 1: union(0, 1)
```
Find(0) → 0 (root)
Find(1) → 1 (root)
Different roots, so merge: parent[1] = 0
parent: [0, 0, 2, 3, 4, 5]
Sets: {0,1}, {2}, {3}, {4}, {5}
Count: 5
```

### Step 2: union(2, 3)
```
Find(2) → 2, Find(3) → 3
Merge: parent[3] = 2
parent: [0, 0, 2, 2, 4, 5]
Sets: {0,1}, {2,3}, {4}, {5}
Count: 4
```

### Step 3: union(1, 3)
```
Find(1) → Follow 1→0, root is 0
Find(3) → Follow 3→2, root is 2
Merge: parent[2] = 0
parent: [0, 0, 0, 2, 4, 5]
Sets: {0,1,2,3}, {4}, {5}
Count: 3
```

## Basic Implementation

### 1. Simple Union-Find

```go
type UnionFind struct {
    parent []int
    count  int
}

func NewUnionFind(n int) *UnionFind {
    parent := make([]int, n)
    for i := 0; i < n; i++ {
        parent[i] = i  // Each element is its own parent
    }
    return &UnionFind{parent: parent, count: n}
}

func (uf *UnionFind) Find(x int) int {
    for uf.parent[x] != x {
        x = uf.parent[x]  // Follow parent pointers
    }
    return x
}

func (uf *UnionFind) Union(x, y int) {
    rootX := uf.Find(x)
    rootY := uf.Find(y)
    
    if rootX != rootY {
        uf.parent[rootX] = rootY  // Make one root point to other
        uf.count--
    }
}
```

**Problem**: This can create very deep trees, making Find() slow in worst case.

## Key Optimizations

### 1. Path Compression

**Problem**: Following long chains of parent pointers is slow.

**Solution**: During Find(), make every node point directly to the root.

```go
func (uf *UnionFind) Find(x int) int {
    if uf.parent[x] != x {
        uf.parent[x] = uf.Find(uf.parent[x])  // Recursive path compression
    }
    return uf.parent[x]
}
```

**Benefit**: Flattens the tree, making future Find() operations faster.

### 2. Union by Rank

**Problem**: Always attaching first tree to second can create unbalanced trees.

**Solution**: Attach smaller tree under the root of larger tree.

```go
type UnionFind struct {
    parent []int
    rank   []int  // Approximate depth of tree
    count  int
}

func (uf *UnionFind) Union(x, y int) bool {
    rootX := uf.Find(x)
    rootY := uf.Find(y)
    
    if rootX == rootY {
        return false
    }
    
    // Union by rank: attach smaller tree under larger
    if uf.rank[rootX] < uf.rank[rootY] {
        uf.parent[rootX] = rootY
    } else if uf.rank[rootX] > uf.rank[rootY] {
        uf.parent[rootY] = rootX
    } else {
        uf.parent[rootY] = rootX
        uf.rank[rootX]++  // Same rank: increase rank of new root
    }
    
    uf.count--
    return true
}
```

### 3. Union by Size (Alternative)

Instead of rank, track actual size of each tree:

```go
type WeightedUnionFind struct {
    parent []int
    size   []int  // Size of tree rooted at i
    count  int
}

func (wuf *WeightedUnionFind) Union(x, y int) bool {
    rootX := wuf.Find(x)
    rootY := wuf.Find(y)
    
    if rootX == rootY {
        return false
    }
    
    // Attach smaller tree to larger tree
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
```

## Complexity Analysis

### Without Optimizations:
- **Find**: O(n) worst case
- **Union**: O(n) worst case

### With Path Compression Only:
- **Find**: O(log n) amortized
- **Union**: O(log n) amortized

### With Union by Rank/Size Only:
- **Find**: O(log n) worst case
- **Union**: O(log n) worst case

### With Both Optimizations:
- **Find**: O(α(n)) amortized
- **Union**: O(α(n)) amortized

Where **α(n)** is the inverse Ackermann function:
- α(n) ≤ 4 for all practical values of n
- α(2^65536) = 5
- Essentially **constant time** for real-world problems

## Real-World Applications

### 1. **Network Connectivity**
- **Problem**: Track which computers are connected in a network
- **Solution**: Union-Find to maintain connected components
- **Operations**: Connect two computers, check if two computers can communicate

```go
// Connect two computers
networkUF.Union(computer1, computer2)

// Check if computers can communicate
canCommunicate := networkUF.Connected(computer1, computer2)
```

### 2. **Kruskal's Minimum Spanning Tree**
- **Problem**: Find minimum cost to connect all cities with roads
- **Solution**: Use Union-Find to detect cycles while adding edges

```go
func KruskalMST(n int, edges []Edge) []Edge {
    sort.Slice(edges, func(i, j int) bool {
        return edges[i].Weight < edges[j].Weight
    })
    
    uf := NewUnionFind(n)
    mst := []Edge{}
    
    for _, edge := range edges {
        if uf.Union(edge.From, edge.To) {  // No cycle created
            mst = append(mst, edge)
        }
    }
    
    return mst
}
```

### 3. **Image Processing - Connected Components**
- **Problem**: Find connected regions in an image
- **Solution**: Union-Find to group adjacent pixels of same color

```go
func NumberOfIslands(grid [][]byte) int {
    uf := NewUnionFind(rows * cols)
    
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == '1' {
                // Check 4 adjacent cells
                for _, dir := range directions {
                    ni, nj := i+dir[0], j+dir[1]
                    if isValid(ni, nj) && grid[ni][nj] == '1' {
                        uf.Union(i*cols+j, ni*cols+nj)
                    }
                }
            }
        }
    }
    
    return countComponents(uf, grid)
}
```

### 4. **Social Networks - Friend Circles**
- **Problem**: Count number of friend groups
- **Solution**: Union friends and count connected components

```go
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
```

### 5. **Percolation Theory**
- **Problem**: Determine if water can flow from top to bottom of a grid
- **Solution**: Union-Find with virtual top and bottom nodes

### 6. **Database Systems**
- **Problem**: Track equivalent rows after joins and updates
- **Solution**: Union-Find to maintain equivalence classes

### 7. **Compiler Design**
- **Problem**: Variable aliasing analysis
- **Solution**: Union-Find to track which variables refer to same memory

### 8. **Game Development**
- **Problem**: Group management in multiplayer games
- **Solution**: Union-Find for guild/team management

## Advanced Applications

### 1. **Account Merging**
Multiple accounts with overlapping emails need to be merged:

```go
func AccountsMerge(accounts [][]string) [][]string {
    // Map each email to unique index
    emailToIndex := make(map[string]int)
    emailToName := make(map[string]string)
    index := 0
    
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
    
    // Union emails in same account
    for _, account := range accounts {
        if len(account) > 1 {
            firstEmailIndex := emailToIndex[account[1]]
            for i := 2; i < len(account); i++ {
                uf.Union(firstEmailIndex, emailToIndex[account[i]])
            }
        }
    }
    
    // Group emails by component and build result
    // ... (implementation details)
}
```

### 2. **Dynamic Connectivity with Rollback**
Some applications need to undo Union operations:

```go
type RollbackUnionFind struct {
    parent []int
    rank   []int
    history []Operation  // Store operations for rollback
}

type Operation struct {
    Type   string  // "union" or "rollback"
    X, Y   int
    OldParent int
    OldRank   int
}
```

### 3. **Weighted Union-Find**
Track additional information about relationships:

```go
type WeightedUF struct {
    parent []int
    weight []int  // Weight of edge from i to parent[i]
}

func (wuf *WeightedUF) Find(x int) (int, int) {
    if wuf.parent[x] != x {
        root, w := wuf.Find(wuf.parent[x])
        wuf.parent[x] = root
        wuf.weight[x] += w  // Path compression with weight update
        return root, wuf.weight[x]
    }
    return x, 0
}
```

## When to Use Union-Find

### ✅ **Perfect for:**
- **Dynamic connectivity** queries
- **Cycle detection** in undirected graphs
- **Connected components** counting
- **Equivalence relations** tracking
- **Kruskal's MST** algorithm
- **Percolation** problems

### ❌ **Not suitable for:**
- **Shortest path** queries (use BFS/Dijkstra)
- **Directed graphs** (edges have direction)
- **Deleting connections** (Union-Find doesn't support efficient split)
- **Exact tree structure** needed (Union-Find flattens trees)

## Comparison with Other Approaches

| Problem | Union-Find | DFS/BFS | Other |
|---------|------------|---------|-------|
| **Connected Components** | O(α(n)) per query | O(V+E) per query | - |
| **Cycle Detection** | O(E⋅α(V)) | O(V+E) | - |
| **MST (Kruskal's)** | O(E⋅α(V)) | Not applicable | Prim's: O(E log V) |
| **Shortest Path** | Not applicable | O(V+E) BFS | Dijkstra: O(E log V) |

## Implementation Tips

### 1. **0-indexed vs 1-indexed**
```go
// Be consistent with indexing
func NewUnionFind(n int) *UnionFind {
    // Creates UF for elements 0 to n-1
}
```

### 2. **Input Validation**
```go
func (uf *UnionFind) Find(x int) int {
    if x < 0 || x >= len(uf.parent) {
        panic("index out of bounds")
    }
    // ... rest of implementation
}
```

### 3. **Coordinate Mapping for 2D Problems**
```go
// Map 2D coordinates to 1D index
func coord2D(row, col, cols int) int {
    return row*cols + col
}

// Usage in grid problems
id1 := coord2D(i, j, cols)
id2 := coord2D(ni, nj, cols)
uf.Union(id1, id2)
```

### 4. **Virtual Nodes**
```go
// For percolation: add virtual top and bottom nodes
func NewPercolationUF(n int) *UnionFind {
    // n*n grid + 2 virtual nodes
    uf := NewUnionFind(n*n + 2)
    virtualTop := n * n
    virtualBottom := n*n + 1
    
    // Connect top row to virtual top
    for j := 0; j < n; j++ {
        uf.Union(j, virtualTop)
    }
    
    return uf
}
```

## Common Mistakes

### 1. **Forgetting Path Compression**
```go
// Wrong: Linear chain traversal
func (uf *UnionFind) Find(x int) int {
    for uf.parent[x] != x {
        x = uf.parent[x]
    }
    return x
}

// Correct: Path compression
func (uf *UnionFind) Find(x int) int {
    if uf.parent[x] != x {
        uf.parent[x] = uf.Find(uf.parent[x])
    }
    return uf.parent[x]
}
```

### 2. **Not Using Union by Rank/Size**
```go
// Wrong: Can create unbalanced trees
func (uf *UnionFind) Union(x, y int) {
    rootX := uf.Find(x)
    rootY := uf.Find(y)
    uf.parent[rootX] = rootY  // Always attach first to second
}

// Correct: Union by rank
func (uf *UnionFind) Union(x, y int) {
    // ... implement union by rank logic
}
```

### 3. **Coordinate Mapping Errors**
```go
// Wrong: Inconsistent mapping
id1 := i*rows + j    // Should be i*cols + j
id2 := ni*cols + nj

// Correct: Consistent row-major mapping
id1 := i*cols + j
id2 := ni*cols + nj
```

## Practice Problems

### Beginner:
1. **Number of Connected Components** in undirected graph
2. **Friend Circles** counting
3. **Find if Path Exists** between two nodes

### Intermediate:
4. **Number of Islands** in 2D grid
5. **Accounts Merge** problem
6. **Redundant Connection** (cycle detection)

### Advanced:
7. **Optimize Water Distribution** (MST variant)
8. **Satisfiability of Equality Equations**
9. **Smallest String With Swaps**

## Key Insights

1. **Tree Representation**: Union-Find represents each set as a tree, with efficient operations due to tree properties

2. **Optimizations Matter**: Path compression and union by rank transform O(n) operations into effectively O(1)

3. **Amortized Analysis**: Individual operations might be slow, but average over many operations is very fast

4. **Inverse Ackermann**: The α(n) function grows so slowly it's essentially constant for all practical purposes

5. **Dynamic vs Static**: Union-Find excels at dynamic connectivity but doesn't support edge deletion efficiently

6. **Graph Problems**: Many graph problems that seem to need DFS/BFS can be solved more efficiently with Union-Find

The Union-Find data structure is a perfect example of how the right optimizations can transform a simple idea into an incredibly efficient tool for connectivity problems! 🚀 