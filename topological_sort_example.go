package main

import (
	"fmt"
)

// DirectedGraph represents a directed graph using adjacency list
type DirectedGraph struct {
	vertices int
	adjList  map[int][]int
}

// NewDirectedGraph creates a new directed graph
func NewDirectedGraph(vertices int) *DirectedGraph {
	return &DirectedGraph{
		vertices: vertices,
		adjList:  make(map[int][]int),
	}
}

// AddEdge adds a directed edge from u to v
func (g *DirectedGraph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
}

// ================================
// TOPOLOGICAL SORT USING DFS
// ================================

// TopologicalSortDFS performs topological sorting using DFS
// Time Complexity: O(V + E)
// Space Complexity: O(V)
func (g *DirectedGraph) TopologicalSortDFS() []int {
	visited := make(map[int]bool)
	stack := []int{}

	// Call DFS for all unvisited vertices
	for vertex := 0; vertex < g.vertices; vertex++ {
		if !visited[vertex] {
			g.topologicalSortUtil(vertex, visited, &stack)
		}
	}

	// Reverse the stack to get topological order
	result := make([]int, len(stack))
	for i := 0; i < len(stack); i++ {
		result[i] = stack[len(stack)-1-i]
	}

	return result
}

// topologicalSortUtil is a recursive utility function for DFS-based topological sort
func (g *DirectedGraph) topologicalSortUtil(vertex int, visited map[int]bool, stack *[]int) {
	visited[vertex] = true

	// Recursively visit all adjacent vertices
	for _, neighbor := range g.adjList[vertex] {
		if !visited[neighbor] {
			g.topologicalSortUtil(neighbor, visited, stack)
		}
	}

	// Push current vertex to stack after visiting all neighbors
	*stack = append(*stack, vertex)
}

// ================================
// TOPOLOGICAL SORT USING KAHN'S ALGORITHM (BFS)
// ================================

// TopologicalSortKahn performs topological sorting using Kahn's algorithm
// Time Complexity: O(V + E)
// Space Complexity: O(V)
func (g *DirectedGraph) TopologicalSortKahn() []int {
	// Calculate in-degree of all vertices
	inDegree := make([]int, g.vertices)
	for vertex := 0; vertex < g.vertices; vertex++ {
		for _, neighbor := range g.adjList[vertex] {
			inDegree[neighbor]++
		}
	}

	// Find all vertices with 0 in-degree
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

		// Reduce in-degree of all adjacent vertices
		for _, neighbor := range g.adjList[vertex] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Check for cycle (if result doesn't contain all vertices)
	if len(result) != g.vertices {
		fmt.Println("Graph contains a cycle! Topological sort not possible.")
		return nil
	}

	return result
}

// ================================
// CYCLE DETECTION
// ================================

// HasCycle detects if the directed graph has a cycle using DFS
func (g *DirectedGraph) HasCycle() bool {
	visited := make(map[int]bool)
	recStack := make(map[int]bool)

	for vertex := 0; vertex < g.vertices; vertex++ {
		if !visited[vertex] {
			if g.hasCycleUtil(vertex, visited, recStack) {
				return true
			}
		}
	}
	return false
}

func (g *DirectedGraph) hasCycleUtil(vertex int, visited, recStack map[int]bool) bool {
	visited[vertex] = true
	recStack[vertex] = true

	for _, neighbor := range g.adjList[vertex] {
		if !visited[neighbor] && g.hasCycleUtil(neighbor, visited, recStack) {
			return true
		} else if recStack[neighbor] {
			return true
		}
	}

	recStack[vertex] = false
	return false
}

// ================================
// PRACTICAL APPLICATIONS
// ================================

// CourseSchedule represents a course scheduling problem
type CourseSchedule struct {
	courses []string
	graph   *DirectedGraph
}

// NewCourseSchedule creates a new course scheduling system
func NewCourseSchedule(courses []string) *CourseSchedule {
	return &CourseSchedule{
		courses: courses,
		graph:   NewDirectedGraph(len(courses)),
	}
}

// AddPrerequisite adds a prerequisite relationship (prerequisite -> course)
func (cs *CourseSchedule) AddPrerequisite(prerequisite, course string) {
	prereqIndex := cs.findCourseIndex(prerequisite)
	courseIndex := cs.findCourseIndex(course)

	if prereqIndex != -1 && courseIndex != -1 {
		cs.graph.AddEdge(prereqIndex, courseIndex)
	}
}

func (cs *CourseSchedule) findCourseIndex(course string) int {
	for i, c := range cs.courses {
		if c == course {
			return i
		}
	}
	return -1
}

// GetOptimalOrder returns the optimal order to take courses
func (cs *CourseSchedule) GetOptimalOrder() []string {
	if cs.graph.HasCycle() {
		fmt.Println("Circular dependency detected! Cannot schedule courses.")
		return nil
	}

	order := cs.graph.TopologicalSortDFS()
	result := make([]string, len(order))

	for i, courseIndex := range order {
		result[i] = cs.courses[courseIndex]
	}

	return result
}

// ================================
// TASK SCHEDULING EXAMPLE
// ================================

// TaskScheduler represents a task scheduling system
type TaskScheduler struct {
	tasks []string
	graph *DirectedGraph
}

// NewTaskScheduler creates a new task scheduler
func NewTaskScheduler(tasks []string) *TaskScheduler {
	return &TaskScheduler{
		tasks: tasks,
		graph: NewDirectedGraph(len(tasks)),
	}
}

// AddDependency adds a task dependency (dependency -> task)
func (ts *TaskScheduler) AddDependency(dependency, task string) {
	depIndex := ts.findTaskIndex(dependency)
	taskIndex := ts.findTaskIndex(task)

	if depIndex != -1 && taskIndex != -1 {
		ts.graph.AddEdge(depIndex, taskIndex)
	}
}

func (ts *TaskScheduler) findTaskIndex(task string) int {
	for i, t := range ts.tasks {
		if t == task {
			return i
		}
	}
	return -1
}

// GetExecutionOrder returns the optimal order to execute tasks
func (ts *TaskScheduler) GetExecutionOrder() []string {
	if ts.graph.HasCycle() {
		fmt.Println("Circular dependency detected! Cannot schedule tasks.")
		return nil
	}

	order := ts.graph.TopologicalSortKahn()
	if order == nil {
		return nil
	}

	result := make([]string, len(order))
	for i, taskIndex := range order {
		result[i] = ts.tasks[taskIndex]
	}

	return result
}

// ================================
// DEMO FUNCTIONS
// ================================

func DemoTopologicalSort() {
	fmt.Println("=== TOPOLOGICAL SORT EXPLANATION ===\n")

	fmt.Println("Topological Sort is a linear ordering of vertices in a Directed Acyclic Graph (DAG)")
	fmt.Println("such that for every directed edge (u,v), vertex u comes before v in the ordering.\n")

	// Example 1: Simple DAG
	fmt.Println("=== EXAMPLE 1: Simple DAG ===")
	fmt.Println("Graph: 0 → 1 → 3")
	fmt.Println("       ↓   ↗")
	fmt.Println("       2 ───")

	graph1 := NewDirectedGraph(4)
	graph1.AddEdge(0, 1)
	graph1.AddEdge(0, 2)
	graph1.AddEdge(1, 3)
	graph1.AddEdge(2, 3)

	fmt.Println("\nAdjacency List:")
	for i := 0; i < graph1.vertices; i++ {
		fmt.Printf("Vertex %d: %v\n", i, graph1.adjList[i])
	}

	dfsResult := graph1.TopologicalSortDFS()
	kahnResult := graph1.TopologicalSortKahn()

	fmt.Printf("\nTopological Sort (DFS):   %v\n", dfsResult)
	fmt.Printf("Topological Sort (Kahn): %v\n", kahnResult)
	fmt.Printf("Has Cycle: %v\n\n", graph1.HasCycle())

	// Example 2: Course Scheduling
	fmt.Println("=== EXAMPLE 2: Course Scheduling Problem ===")
	courses := []string{"Math", "Physics", "Chemistry", "Biology", "Advanced Physics"}

	cs := NewCourseSchedule(courses)
	cs.AddPrerequisite("Math", "Physics")
	cs.AddPrerequisite("Math", "Chemistry")
	cs.AddPrerequisite("Physics", "Advanced Physics")
	cs.AddPrerequisite("Chemistry", "Biology")

	fmt.Println("Prerequisites:")
	fmt.Println("- Math → Physics")
	fmt.Println("- Math → Chemistry")
	fmt.Println("- Physics → Advanced Physics")
	fmt.Println("- Chemistry → Biology")

	optimalOrder := cs.GetOptimalOrder()
	fmt.Printf("\nOptimal Course Order: %v\n\n", optimalOrder)

	// Example 3: Task Scheduling
	fmt.Println("=== EXAMPLE 3: Task Scheduling ===")
	tasks := []string{"Setup", "Design", "Code", "Test", "Deploy", "Documentation"}

	ts := NewTaskScheduler(tasks)
	ts.AddDependency("Setup", "Design")
	ts.AddDependency("Design", "Code")
	ts.AddDependency("Code", "Test")
	ts.AddDependency("Test", "Deploy")
	ts.AddDependency("Design", "Documentation")

	fmt.Println("Task Dependencies:")
	fmt.Println("- Setup → Design")
	fmt.Println("- Design → Code")
	fmt.Println("- Code → Test")
	fmt.Println("- Test → Deploy")
	fmt.Println("- Design → Documentation")

	executionOrder := ts.GetExecutionOrder()
	fmt.Printf("\nOptimal Task Execution Order: %v\n\n", executionOrder)

	// Example 4: Graph with Cycle
	fmt.Println("=== EXAMPLE 4: Graph with Cycle ===")
	cyclicGraph := NewDirectedGraph(3)
	cyclicGraph.AddEdge(0, 1)
	cyclicGraph.AddEdge(1, 2)
	cyclicGraph.AddEdge(2, 0) // Creates a cycle

	fmt.Println("Graph: 0 → 1 → 2 → 0 (cycle)")
	fmt.Printf("Has Cycle: %v\n", cyclicGraph.HasCycle())

	fmt.Println("\nTrying topological sort on cyclic graph:")
	cyclicResult := cyclicGraph.TopologicalSortKahn()
	if cyclicResult == nil {
		fmt.Println("Topological sort failed due to cycle detection.\n")
	}

	// Example 5: Complex DAG
	fmt.Println("=== EXAMPLE 5: Complex DAG ===")
	complexGraph := NewDirectedGraph(6)
	complexGraph.AddEdge(5, 2)
	complexGraph.AddEdge(5, 0)
	complexGraph.AddEdge(4, 0)
	complexGraph.AddEdge(4, 1)
	complexGraph.AddEdge(2, 3)
	complexGraph.AddEdge(3, 1)

	fmt.Println("Graph: 5 → 2 → 3 → 1")
	fmt.Println("       ↓       ↗")
	fmt.Println("       0   4 ───")
	fmt.Println("           ↓")
	fmt.Println("           1")

	complexDFS := complexGraph.TopologicalSortDFS()
	complexKahn := complexGraph.TopologicalSortKahn()

	fmt.Printf("\nTopological Sort (DFS):   %v\n", complexDFS)
	fmt.Printf("Topological Sort (Kahn): %v\n", complexKahn)

	fmt.Println("\n=== ALGORITHM COMPARISON ===")
	fmt.Println("DFS-based Topological Sort:")
	fmt.Println("- Uses recursion and stack")
	fmt.Println("- Post-order traversal")
	fmt.Println("- Good for detecting cycles")
	fmt.Println("- Time: O(V + E), Space: O(V)")

	fmt.Println("\nKahn's Algorithm (BFS-based):")
	fmt.Println("- Uses queue and in-degree calculation")
	fmt.Println("- Processes vertices with 0 in-degree first")
	fmt.Println("- Natural cycle detection")
	fmt.Println("- Time: O(V + E), Space: O(V)")
	fmt.Println("- More intuitive for beginners")
}

// // Application-specific demo
// func DemoApplications() {
// 	fmt.Println("\n=== REAL-WORLD APPLICATIONS ===")

// 	fmt.Println("1. Build Systems:")
// 	fmt.Println("   - Determining compilation order")
// 	fmt.Println("   - Makefile dependencies")

// 	fmt.Println("\n2. Course Prerequisites:")
// 	fmt.Println("   - Academic curriculum planning")
// 	fmt.Println("   - Skill development paths")

// 	fmt.Println("\n3. Task Scheduling:")
// 	fmt.Println("   - Project management")
// 	fmt.Println("   - Manufacturing processes")

// 	fmt.Println("\n4. Package Management:")/
// 	fmt.Println("   - Software dependency resolution")
// 	fmt.Println("   - Library installation order")

// 	fmt.Println("\n5. Data Processing Pipelines:")
// 	fmt.Println("   - ETL operations")
// 	fmt.Println("   - Machine learning workflows")
// }
