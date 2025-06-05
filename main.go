package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to DSA Practice!")

	// Run Union-Find demonstration
	DemoUnionFind()
	DemoAdvancedApplications()

	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")

	// Run KMP Algorithm demonstration
	DemoKMP()

	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")

	DemoKMPApplications()

	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")

	// Run Dijkstra's Algorithm demonstration
	DemoDijkstra()

	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")

	DemoDijkstraApplications()

	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")

	// Run Morris Traversal demonstration
	DemoMorrisTraversal()

	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")

	DemoMorrisApplications()
}
