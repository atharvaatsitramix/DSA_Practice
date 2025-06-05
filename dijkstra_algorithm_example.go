package main

import (
	"container/heap"
	"fmt"
	"math"
)

// ================================
// DIJKSTRA'S ALGORITHM
// ================================

// WeightedEdge represents a weighted edge in the graph
type WeightedEdge struct {
	to     int     // destination vertex
	weight float64 // edge weight
}

// WeightedGraph represents a weighted directed graph
type WeightedGraph struct {
	vertices int
	adjList  [][]WeightedEdge
}

// NewWeightedGraph creates a new weighted graph
func NewWeightedGraph(vertices int) *WeightedGraph {
	return &WeightedGraph{
		vertices: vertices,
		adjList:  make([][]WeightedEdge, vertices),
	}
}

// AddEdge adds a weighted edge to the graph
func (g *WeightedGraph) AddEdge(from, to int, weight float64) {
	g.adjList[from] = append(g.adjList[from], WeightedEdge{to: to, weight: weight})
}

// AddUndirectedEdge adds an undirected weighted edge
func (g *WeightedGraph) AddUndirectedEdge(u, v int, weight float64) {
	g.AddEdge(u, v, weight)
	g.AddEdge(v, u, weight)
}

// PrintGraph displays the graph structure
func (g *WeightedGraph) PrintGraph() {
	fmt.Println("Graph structure:")
	for i := 0; i < g.vertices; i++ {
		fmt.Printf("Vertex %d: ", i)
		for _, edge := range g.adjList[i] {
			fmt.Printf("-> %d(%.1f) ", edge.to, edge.weight)
		}
		fmt.Println()
	}
	fmt.Println()
}

// ================================
// PRIORITY QUEUE IMPLEMENTATION
// ================================

// PQItem represents an item in the priority queue
type PQItem struct {
	vertex   int     // vertex number
	distance float64 // distance from source
	index    int     // index in the heap
}

// PriorityQueue implements a min-heap for Dijkstra's algorithm
type PriorityQueue []*PQItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PQItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// update modifies the distance of an item in the queue
func (pq *PriorityQueue) update(item *PQItem, distance float64) {
	item.distance = distance
	heap.Fix(pq, item.index)
}

// ================================
// DIJKSTRA'S ALGORITHM IMPLEMENTATION
// ================================

// DijkstraResult contains the results of Dijkstra's algorithm
type DijkstraResult struct {
	distances []float64 // shortest distances from source
	previous  []int     // previous vertex in shortest path
	source    int       // source vertex
	visited   []bool    // vertices that have been processed
}

// Dijkstra implements Dijkstra's shortest path algorithm
func (g *WeightedGraph) Dijkstra(source int) *DijkstraResult {
	fmt.Printf("=== DIJKSTRA'S ALGORITHM FROM VERTEX %d ===\n\n", source)

	// Initialize distances and previous vertices
	distances := make([]float64, g.vertices)
	previous := make([]int, g.vertices)
	visited := make([]bool, g.vertices)

	// Initialize all distances to infinity except source
	for i := 0; i < g.vertices; i++ {
		distances[i] = math.Inf(1)
		previous[i] = -1
	}
	distances[source] = 0

	// Create priority queue and add source
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// Track items in priority queue for updates
	items := make([]*PQItem, g.vertices)
	for i := 0; i < g.vertices; i++ {
		item := &PQItem{
			vertex:   i,
			distance: distances[i],
		}
		items[i] = item
		heap.Push(&pq, item)
	}

	fmt.Printf("Initial state:\n")
	fmt.Printf("Distances: %v\n", formatDistances(distances))
	fmt.Printf("Previous:  %v\n\n", previous)

	step := 1

	// Main algorithm loop
	for pq.Len() > 0 {
		// Extract vertex with minimum distance
		current := heap.Pop(&pq).(*PQItem)
		u := current.vertex

		if visited[u] {
			continue
		}

		visited[u] = true
		fmt.Printf("Step %d: Process vertex %d (distance %.1f)\n", step, u, distances[u])

		// If distance is infinity, remaining vertices are unreachable
		if distances[u] == math.Inf(1) {
			fmt.Printf("  All remaining vertices are unreachable\n")
			break
		}

		// Update distances to all adjacent vertices
		fmt.Printf("  Checking neighbors: ")
		hasNeighbors := false
		for _, edge := range g.adjList[u] {
			v := edge.to
			weight := edge.weight

			if !visited[v] {
				hasNeighbors = true
				newDistance := distances[u] + weight
				fmt.Printf("%d(%.1f) ", v, weight)

				if newDistance < distances[v] {
					fmt.Printf("[UPDATED: %.1f->%.1f] ", distances[v], newDistance)
					distances[v] = newDistance
					previous[v] = u

					// Update priority queue
					if items[v].index >= 0 {
						pq.update(items[v], newDistance)
					}
				}
			}
		}

		if !hasNeighbors {
			fmt.Printf("none")
		}
		fmt.Println()

		fmt.Printf("  Updated distances: %v\n", formatDistances(distances))
		fmt.Printf("  Updated previous:  %v\n\n", previous)
		step++
	}

	return &DijkstraResult{
		distances: distances,
		previous:  previous,
		source:    source,
		visited:   visited,
	}
}

// formatDistances formats distances for pretty printing
func formatDistances(distances []float64) []string {
	result := make([]string, len(distances))
	for i, d := range distances {
		if d == math.Inf(1) {
			result[i] = "∞"
		} else {
			result[i] = fmt.Sprintf("%.1f", d)
		}
	}
	return result
}

// GetPath reconstructs the shortest path from source to target
func (result *DijkstraResult) GetPath(target int) []int {
	if result.distances[target] == math.Inf(1) {
		return nil // No path exists
	}

	path := []int{}
	current := target

	// Trace back from target to source
	for current != -1 {
		path = append([]int{current}, path...) // Prepend to path
		current = result.previous[current]
	}

	return path
}

// GetDistance returns the shortest distance to a vertex
func (result *DijkstraResult) GetDistance(vertex int) float64 {
	return result.distances[vertex]
}

// PrintResults displays the complete results
func (result *DijkstraResult) PrintResults() {
	fmt.Printf("=== FINAL RESULTS ===\n")
	fmt.Printf("Shortest distances from vertex %d:\n", result.source)

	for i := 0; i < len(result.distances); i++ {
		if result.distances[i] == math.Inf(1) {
			fmt.Printf("  To vertex %d: unreachable\n", i)
		} else {
			fmt.Printf("  To vertex %d: %.1f\n", i, result.distances[i])
		}
	}
	fmt.Println()

	fmt.Println("Shortest paths:")
	for i := 0; i < len(result.distances); i++ {
		if i != result.source {
			path := result.GetPath(i)
			if path != nil {
				fmt.Printf("  Path to %d: %v (distance: %.1f)\n", i, path, result.distances[i])
			} else {
				fmt.Printf("  Path to %d: no path exists\n", i)
			}
		}
	}
	fmt.Println()
}

// ================================
// PRACTICAL APPLICATIONS
// ================================

// CityMap represents a city road network
type CityMap struct {
	graph     *WeightedGraph
	cityNames []string
}

// NewCityMap creates a new city map
func NewCityMap(cities []string) *CityMap {
	return &CityMap{
		graph:     NewWeightedGraph(len(cities)),
		cityNames: cities,
	}
}

// AddRoad adds a bidirectional road between cities
func (cm *CityMap) AddRoad(city1, city2 string, distance float64) {
	from := cm.findCityIndex(city1)
	to := cm.findCityIndex(city2)
	if from >= 0 && to >= 0 {
		cm.graph.AddUndirectedEdge(from, to, distance)
	}
}

// findCityIndex finds the index of a city
func (cm *CityMap) findCityIndex(city string) int {
	for i, name := range cm.cityNames {
		if name == city {
			return i
		}
	}
	return -1
}

// FindShortestRoute finds the shortest route between two cities
func (cm *CityMap) FindShortestRoute(from, to string) {
	fromIndex := cm.findCityIndex(from)
	toIndex := cm.findCityIndex(to)

	if fromIndex < 0 || toIndex < 0 {
		fmt.Printf("City not found\n")
		return
	}

	fmt.Printf("=== GPS NAVIGATION: %s to %s ===\n\n", from, to)

	// Print city map
	fmt.Println("City Network:")
	for i, city := range cm.cityNames {
		fmt.Printf("%d: %s\n", i, city)
	}
	fmt.Println()

	result := cm.graph.Dijkstra(fromIndex)

	path := result.GetPath(toIndex)
	distance := result.GetDistance(toIndex)

	if path != nil {
		fmt.Printf("Shortest route from %s to %s:\n", from, to)
		for i, cityIndex := range path {
			if i > 0 {
				fmt.Printf(" -> ")
			}
			fmt.Printf("%s", cm.cityNames[cityIndex])
		}
		fmt.Printf("\nTotal distance: %.1f km\n\n", distance)
	} else {
		fmt.Printf("No route found from %s to %s\n\n", from, to)
	}
}

// NetworkRouter simulates network packet routing
type NetworkRouter struct {
	graph     *WeightedGraph
	nodeNames []string
}

// NewNetworkRouter creates a new network router
func NewNetworkRouter(nodes []string) *NetworkRouter {
	return &NetworkRouter{
		graph:     NewWeightedGraph(len(nodes)),
		nodeNames: nodes,
	}
}

// AddConnection adds a network connection with latency
func (nr *NetworkRouter) AddConnection(node1, node2 string, latency float64) {
	from := nr.findNodeIndex(node1)
	to := nr.findNodeIndex(node2)
	if from >= 0 && to >= 0 {
		nr.graph.AddUndirectedEdge(from, to, latency)
	}
}

// findNodeIndex finds the index of a network node
func (nr *NetworkRouter) findNodeIndex(node string) int {
	for i, name := range nr.nodeNames {
		if name == node {
			return i
		}
	}
	return -1
}

// FindOptimalRoute finds the route with minimum latency
func (nr *NetworkRouter) FindOptimalRoute(source, destination string) {
	sourceIndex := nr.findNodeIndex(source)
	destIndex := nr.findNodeIndex(destination)

	if sourceIndex < 0 || destIndex < 0 {
		fmt.Printf("Network node not found\n")
		return
	}

	fmt.Printf("=== NETWORK ROUTING: %s to %s ===\n\n", source, destination)

	result := nr.graph.Dijkstra(sourceIndex)

	path := result.GetPath(destIndex)
	latency := result.GetDistance(destIndex)

	if path != nil {
		fmt.Printf("Optimal route (minimum latency):\n")
		for i, nodeIndex := range path {
			if i > 0 {
				fmt.Printf(" -> ")
			}
			fmt.Printf("%s", nr.nodeNames[nodeIndex])
		}
		fmt.Printf("\nTotal latency: %.1f ms\n\n", latency)
	} else {
		fmt.Printf("No route found from %s to %s\n\n", source, destination)
	}
}

// ================================
// ALGORITHM VARIATIONS
// ================================

// DijkstraWithPath implements Dijkstra with path tracking in single pass
func (g *WeightedGraph) DijkstraWithPath(source, target int) (float64, []int) {
	distances := make([]float64, g.vertices)
	previous := make([]int, g.vertices)
	visited := make([]bool, g.vertices)

	for i := 0; i < g.vertices; i++ {
		distances[i] = math.Inf(1)
		previous[i] = -1
	}
	distances[source] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &PQItem{vertex: source, distance: 0})

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*PQItem)
		u := current.vertex

		if u == target {
			// Found target, reconstruct path
			path := []int{}
			curr := target
			for curr != -1 {
				path = append([]int{curr}, path...)
				curr = previous[curr]
			}
			return distances[target], path
		}

		if visited[u] {
			continue
		}
		visited[u] = true

		for _, edge := range g.adjList[u] {
			v := edge.to
			if !visited[v] {
				newDistance := distances[u] + edge.weight
				if newDistance < distances[v] {
					distances[v] = newDistance
					previous[v] = u
					heap.Push(&pq, &PQItem{vertex: v, distance: newDistance})
				}
			}
		}
	}

	return math.Inf(1), nil // No path found
}

// AllPairsShortestPath computes shortest paths between all pairs of vertices
func (g *WeightedGraph) AllPairsShortestPath() [][]float64 {
	distances := make([][]float64, g.vertices)

	for i := 0; i < g.vertices; i++ {
		result := g.Dijkstra(i)
		distances[i] = make([]float64, g.vertices)
		copy(distances[i], result.distances)
	}

	return distances
}

// ================================
// DEMONSTRATION FUNCTIONS
// ================================

// DemoDijkstra demonstrates Dijkstra's algorithm with examples
func DemoDijkstra() {
	fmt.Println("=== DIJKSTRA'S SHORTEST PATH ALGORITHM ===\n")

	fmt.Println("Dijkstra's algorithm finds the shortest path from a source vertex")
	fmt.Println("to all other vertices in a weighted graph with non-negative edge weights.")
	fmt.Println("It uses a greedy approach with a priority queue for efficiency.")
	fmt.Println()

	// Example 1: Basic graph
	fmt.Println("=== EXAMPLE 1: Simple Weighted Graph ===")
	graph1 := NewWeightedGraph(5)

	// Build a sample graph
	graph1.AddEdge(0, 1, 4.0)
	graph1.AddEdge(0, 2, 2.0)
	graph1.AddEdge(1, 2, 1.0)
	graph1.AddEdge(1, 3, 5.0)
	graph1.AddEdge(2, 3, 8.0)
	graph1.AddEdge(2, 4, 10.0)
	graph1.AddEdge(3, 4, 2.0)

	graph1.PrintGraph()

	result1 := graph1.Dijkstra(0)
	result1.PrintResults()

	// Example 2: Disconnected graph
	fmt.Println("=== EXAMPLE 2: Graph with Unreachable Vertices ===")
	graph2 := NewWeightedGraph(6)

	// Connected component 1: vertices 0, 1, 2
	graph2.AddUndirectedEdge(0, 1, 3.0)
	graph2.AddUndirectedEdge(1, 2, 2.0)

	// Connected component 2: vertices 3, 4
	graph2.AddUndirectedEdge(3, 4, 1.0)

	// Isolated vertex: 5

	graph2.PrintGraph()

	result2 := graph2.Dijkstra(0)
	result2.PrintResults()
}

// DemoDijkstraApplications shows practical applications
func DemoDijkstraApplications() {
	fmt.Println("=== PRACTICAL APPLICATIONS ===\n")

	// Application 1: GPS Navigation
	fmt.Println("1. GPS NAVIGATION SYSTEM")
	cities := []string{"New York", "Boston", "Philadelphia", "Washington DC", "Atlanta", "Miami"}
	cityMap := NewCityMap(cities)

	// Add roads with distances (simplified)
	cityMap.AddRoad("New York", "Boston", 215)
	cityMap.AddRoad("New York", "Philadelphia", 95)
	cityMap.AddRoad("Philadelphia", "Washington DC", 140)
	cityMap.AddRoad("Washington DC", "Atlanta", 440)
	cityMap.AddRoad("Atlanta", "Miami", 650)
	cityMap.AddRoad("Boston", "Philadelphia", 300)
	cityMap.AddRoad("New York", "Washington DC", 225)

	cityMap.FindShortestRoute("New York", "Miami")

	// Application 2: Network Routing
	fmt.Println("2. NETWORK PACKET ROUTING")
	nodes := []string{"Router-A", "Router-B", "Router-C", "Router-D", "Server", "Client"}
	network := NewNetworkRouter(nodes)

	// Add connections with latencies in milliseconds
	network.AddConnection("Client", "Router-A", 5.0)
	network.AddConnection("Router-A", "Router-B", 10.0)
	network.AddConnection("Router-A", "Router-C", 15.0)
	network.AddConnection("Router-B", "Router-D", 12.0)
	network.AddConnection("Router-C", "Router-D", 8.0)
	network.AddConnection("Router-D", "Server", 6.0)
	network.AddConnection("Router-B", "Server", 20.0) // Direct but slower route

	network.FindOptimalRoute("Client", "Server")

	// Application 3: Cost optimization
	fmt.Println("3. FLIGHT ROUTE OPTIMIZATION")
	airports := []string{"JFK", "LAX", "ORD", "DFW", "ATL", "DEN"}
	flightNetwork := NewCityMap(airports)

	// Add flights with costs
	flightNetwork.AddRoad("JFK", "LAX", 350) // Direct flight
	flightNetwork.AddRoad("JFK", "ORD", 180)
	flightNetwork.AddRoad("JFK", "ATL", 200)
	flightNetwork.AddRoad("ORD", "DFW", 160)
	flightNetwork.AddRoad("ORD", "DEN", 140)
	flightNetwork.AddRoad("DFW", "LAX", 180)
	flightNetwork.AddRoad("ATL", "DFW", 150)
	flightNetwork.AddRoad("DEN", "LAX", 120)

	fmt.Println("Finding cheapest flight route:")
	flightNetwork.FindShortestRoute("JFK", "LAX")

	// Application 4: Supply chain optimization
	fmt.Println("4. SUPPLY CHAIN LOGISTICS")
	locations := []string{"Factory", "Warehouse-A", "Warehouse-B", "Distribution-Center", "Retail-Store"}
	supplyChain := NewCityMap(locations)

	// Add routes with transportation costs
	supplyChain.AddRoad("Factory", "Warehouse-A", 50)
	supplyChain.AddRoad("Factory", "Warehouse-B", 70)
	supplyChain.AddRoad("Warehouse-A", "Distribution-Center", 30)
	supplyChain.AddRoad("Warehouse-B", "Distribution-Center", 25)
	supplyChain.AddRoad("Distribution-Center", "Retail-Store", 15)
	supplyChain.AddRoad("Warehouse-A", "Retail-Store", 60) // Direct route

	fmt.Println("Finding most cost-effective supply route:")
	supplyChain.FindShortestRoute("Factory", "Retail-Store")
}

// DemoComplexityAnalysis demonstrates algorithm performance characteristics
func DemoComplexityAnalysis() {
	fmt.Println("=== COMPLEXITY ANALYSIS ===\n")

	fmt.Println("Time Complexity:")
	fmt.Println("- Using Binary Heap (Priority Queue): O((V + E) log V)")
	fmt.Println("  where V = number of vertices, E = number of edges")
	fmt.Println("- Using Fibonacci Heap: O(E + V log V) [theoretical optimum]")
	fmt.Println("- Using Simple Array: O(V²) [suitable for dense graphs]")
	fmt.Println()

	fmt.Println("Space Complexity: O(V)")
	fmt.Println("- Distance array: O(V)")
	fmt.Println("- Previous array: O(V)")
	fmt.Println("- Priority queue: O(V)")
	fmt.Println()

	fmt.Println("Key Properties:")
	fmt.Println("✓ Finds optimal shortest paths")
	fmt.Println("✓ Works with non-negative edge weights")
	fmt.Println("✓ Greedy algorithm - makes locally optimal choices")
	fmt.Println("✓ Single-source shortest path algorithm")
	fmt.Println("✗ Cannot handle negative edge weights")
	fmt.Println("✗ More complex than BFS for unweighted graphs")
	fmt.Println()

	fmt.Println("When to use Dijkstra vs alternatives:")
	fmt.Println("- Dijkstra: Single source, non-negative weights, optimal paths")
	fmt.Println("- Bellman-Ford: Single source, allows negative weights")
	fmt.Println("- Floyd-Warshall: All pairs shortest paths, allows negative weights")
	fmt.Println("- BFS: Unweighted graphs (simpler and faster)")
	fmt.Println("- A*: Single target with heuristic (often faster in practice)")
	fmt.Println()
}
