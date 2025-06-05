package main

import (
	"fmt"
)

// ================================
// MORRIS TRAVERSAL (THREADED BINARY TREE)
// ================================

// MorrisTreeNode represents a node in the binary tree
type MorrisTreeNode struct {
	Val   int
	Left  *MorrisTreeNode
	Right *MorrisTreeNode
}

// NewMorrisTreeNode creates a new tree node
func NewMorrisTreeNode(val int) *MorrisTreeNode {
	return &MorrisTreeNode{Val: val}
}

// ================================
// MORRIS INORDER TRAVERSAL IMPLEMENTATION
// ================================

// MorrisInorderTraversal performs inorder traversal using Morris algorithm
func MorrisInorderTraversal(root *MorrisTreeNode) []int {
	result := []int{}
	current := root

	fmt.Println("=== MORRIS INORDER TRAVERSAL ===")
	fmt.Printf("Starting traversal from root\n\n")

	step := 1

	for current != nil {
		fmt.Printf("Step %d: Current node = %d\n", step, current.Val)

		if current.Left == nil {
			// No left subtree, visit current and go right
			fmt.Printf("  No left child, visiting node %d\n", current.Val)
			result = append(result, current.Val)
			current = current.Right
			fmt.Printf("  Moving to right child\n")
		} else {
			// Find inorder predecessor (rightmost node in left subtree)
			predecessor := current.Left
			fmt.Printf("  Has left child, finding inorder predecessor...\n")

			// Find the rightmost node in left subtree or the node that already points to current
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				// First time visiting, create thread and go left
				fmt.Printf("  Predecessor %d found, creating thread to current node %d\n",
					predecessor.Val, current.Val)
				predecessor.Right = current // Create thread
				current = current.Left
				fmt.Printf("  Moving to left child\n")
			} else {
				// Thread already exists, remove it, visit current, and go right
				fmt.Printf("  Thread already exists, removing thread from %d\n", predecessor.Val)
				predecessor.Right = nil // Remove thread
				fmt.Printf("  Visiting node %d\n", current.Val)
				result = append(result, current.Val)
				current = current.Right
				fmt.Printf("  Moving to right child\n")
			}
		}

		fmt.Printf("  Current result: %v\n\n", result)
		step++
	}

	fmt.Printf("Traversal complete! Final result: %v\n\n", result)
	return result
}

// MorrisInorderSimple is a clean version without debug output
func MorrisInorderSimple(root *MorrisTreeNode) []int {
	result := []int{}
	current := root

	for current != nil {
		if current.Left == nil {
			// No left subtree, visit current and go right
			result = append(result, current.Val)
			current = current.Right
		} else {
			// Find inorder predecessor
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				// Create thread and go left
				predecessor.Right = current
				current = current.Left
			} else {
				// Remove thread, visit current, and go right
				predecessor.Right = nil
				result = append(result, current.Val)
				current = current.Right
			}
		}
	}

	return result
}

// ================================
// COMPARISON WITH TRADITIONAL METHODS
// ================================

// RecursiveInorder performs traditional recursive inorder traversal
func RecursiveInorder(root *MorrisTreeNode) []int {
	result := []int{}
	inorderHelper(root, &result)
	return result
}

func inorderHelper(node *MorrisTreeNode, result *[]int) {
	if node != nil {
		inorderHelper(node.Left, result)
		*result = append(*result, node.Val)
		inorderHelper(node.Right, result)
	}
}

// IterativeInorder performs iterative inorder traversal using stack
func IterativeInorder(root *MorrisTreeNode) []int {
	result := []int{}
	stack := []*MorrisTreeNode{}
	current := root

	for current != nil || len(stack) > 0 {
		// Go to leftmost node
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// Pop from stack and visit
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Val)

		// Go to right subtree
		current = current.Right
	}

	return result
}

// ================================
// MORRIS PREORDER TRAVERSAL
// ================================

// MorrisPreorderTraversal performs preorder traversal using Morris algorithm
func MorrisPreorderTraversal(root *MorrisTreeNode) []int {
	result := []int{}
	current := root

	fmt.Println("=== MORRIS PREORDER TRAVERSAL ===")

	for current != nil {
		if current.Left == nil {
			// No left subtree, visit current and go right
			result = append(result, current.Val)
			current = current.Right
		} else {
			// Find inorder predecessor
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				// First time visiting, visit current, create thread, and go left
				result = append(result, current.Val) // Visit before going left (preorder)
				predecessor.Right = current
				current = current.Left
			} else {
				// Remove thread and go right (don't visit again)
				predecessor.Right = nil
				current = current.Right
			}
		}
	}

	fmt.Printf("Preorder result: %v\n\n", result)
	return result
}

// ================================
// TREE CONSTRUCTION AND UTILITIES
// ================================

// BuildSampleTree creates a sample binary tree for demonstration
func BuildSampleTree() *MorrisTreeNode {
	//        4
	//       / \
	//      2   6
	//     / \ / \
	//    1  3 5  7
	root := NewMorrisTreeNode(4)
	root.Left = NewMorrisTreeNode(2)
	root.Right = NewMorrisTreeNode(6)
	root.Left.Left = NewMorrisTreeNode(1)
	root.Left.Right = NewMorrisTreeNode(3)
	root.Right.Left = NewMorrisTreeNode(5)
	root.Right.Right = NewMorrisTreeNode(7)
	return root
}

// BuildComplexTree creates a more complex binary tree
func BuildComplexTree() *MorrisTreeNode {
	//         10
	//        /  \
	//       5    15
	//      / \   / \
	//     3   7 12  20
	//    /   / \     \
	//   1   6   8    25
	root := NewMorrisTreeNode(10)
	root.Left = NewMorrisTreeNode(5)
	root.Right = NewMorrisTreeNode(15)
	root.Left.Left = NewMorrisTreeNode(3)
	root.Left.Right = NewMorrisTreeNode(7)
	root.Right.Left = NewMorrisTreeNode(12)
	root.Right.Right = NewMorrisTreeNode(20)
	root.Left.Left.Left = NewMorrisTreeNode(1)
	root.Left.Right.Left = NewMorrisTreeNode(6)
	root.Left.Right.Right = NewMorrisTreeNode(8)
	root.Right.Right.Right = NewMorrisTreeNode(25)
	return root
}

// BuildLinearTree creates a linear tree (worst case for space complexity)
func BuildLinearTree() *MorrisTreeNode {
	// 1
	//  \
	//   2
	//    \
	//     3
	//      \
	//       4
	//        \
	//         5
	root := NewMorrisTreeNode(1)
	root.Right = NewMorrisTreeNode(2)
	root.Right.Right = NewMorrisTreeNode(3)
	root.Right.Right.Right = NewMorrisTreeNode(4)
	root.Right.Right.Right.Right = NewMorrisTreeNode(5)
	return root
}

// PrintTree displays the tree structure (simple visualization)
func PrintTree(root *MorrisTreeNode, prefix string, isLast bool) {
	if root == nil {
		return
	}

	fmt.Print(prefix)
	if isLast {
		fmt.Print("└── ")
		prefix += "    "
	} else {
		fmt.Print("├── ")
		prefix += "│   "
	}
	fmt.Println(root.Val)

	children := []*MorrisTreeNode{}
	if root.Left != nil {
		children = append(children, root.Left)
	}
	if root.Right != nil {
		children = append(children, root.Right)
	}

	for _, child := range children {
		if child == root.Left {
			fmt.Print(prefix + "├── [L] ")
			fmt.Println(child.Val)
			PrintTree(child.Left, prefix+"│   ", child.Right == nil)
			PrintTree(child.Right, prefix+"│   ", true)
		} else {
			fmt.Print(prefix + "└── [R] ")
			fmt.Println(child.Val)
			PrintTree(child.Left, prefix+"    ", child.Right == nil)
			PrintTree(child.Right, prefix+"    ", true)
		}
	}
}

// VisualizeTree provides a simple tree visualization
func VisualizeTree(root *MorrisTreeNode) {
	if root == nil {
		fmt.Println("Empty tree")
		return
	}

	fmt.Println("Tree structure:")
	levels := getLevels(root)

	for level, nodes := range levels {
		fmt.Printf("Level %d: ", level)
		for _, node := range nodes {
			if node != nil {
				fmt.Printf("%d ", node.Val)
			} else {
				fmt.Printf("null ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// getLevels returns nodes at each level for visualization
func getLevels(root *MorrisTreeNode) map[int][]*MorrisTreeNode {
	levels := make(map[int][]*MorrisTreeNode)
	if root == nil {
		return levels
	}

	queue := []*MorrisTreeNode{root}
	levelQueue := []int{0}

	for len(queue) > 0 {
		node := queue[0]
		level := levelQueue[0]
		queue = queue[1:]
		levelQueue = levelQueue[1:]

		if levels[level] == nil {
			levels[level] = []*MorrisTreeNode{}
		}
		levels[level] = append(levels[level], node)

		if node != nil {
			queue = append(queue, node.Left, node.Right)
			levelQueue = append(levelQueue, level+1, level+1)
		}
	}

	return levels
}

// ================================
// PERFORMANCE ANALYSIS
// ================================

// PerformanceComparison compares different traversal methods
func PerformanceComparison(root *MorrisTreeNode) {
	fmt.Println("=== PERFORMANCE COMPARISON ===")

	// Test all three methods
	fmt.Println("1. Recursive Inorder (uses O(h) space for call stack):")
	recursiveResult := RecursiveInorder(root)
	fmt.Printf("   Result: %v\n", recursiveResult)

	fmt.Println("\n2. Iterative Inorder (uses O(h) space for explicit stack):")
	iterativeResult := IterativeInorder(root)
	fmt.Printf("   Result: %v\n", iterativeResult)

	fmt.Println("\n3. Morris Inorder (uses O(1) space):")
	morrisResult := MorrisInorderSimple(root)
	fmt.Printf("   Result: %v\n", morrisResult)

	// Verify all methods produce same result
	fmt.Printf("\nAll methods produce same result: %v\n\n",
		equalIntSlices(recursiveResult, iterativeResult) &&
			equalIntSlices(iterativeResult, morrisResult))
}

// equalIntSlices checks if two slices are equal
func equalIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// ================================
// ADVANCED APPLICATIONS
// ================================

// MorrisTraversalValidator validates BST property using Morris traversal
func MorrisTraversalValidator(root *MorrisTreeNode) bool {
	if root == nil {
		return true
	}

	current := root
	prev := -1 << 31 // Minimum integer value

	fmt.Println("=== BST VALIDATION USING MORRIS TRAVERSAL ===")

	for current != nil {
		if current.Left == nil {
			// Visit current node
			fmt.Printf("Visiting node %d (previous was %d)\n", current.Val, prev)
			if current.Val <= prev {
				fmt.Printf("BST property violated: %d <= %d\n", current.Val, prev)
				return false
			}
			prev = current.Val
			current = current.Right
		} else {
			// Find predecessor
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				// Create thread
				predecessor.Right = current
				current = current.Left
			} else {
				// Remove thread and visit current
				predecessor.Right = nil
				fmt.Printf("Visiting node %d (previous was %d)\n", current.Val, prev)
				if current.Val <= prev {
					fmt.Printf("BST property violated: %d <= %d\n", current.Val, prev)
					return false
				}
				prev = current.Val
				current = current.Right
			}
		}
	}

	fmt.Println("BST property maintained throughout traversal")
	return true
}

// KthSmallestElementMorris finds kth smallest element using Morris traversal
func KthSmallestElementMorris(root *MorrisTreeNode, k int) int {
	if root == nil || k <= 0 {
		return -1
	}

	current := root
	count := 0

	fmt.Printf("=== FINDING %d-TH SMALLEST ELEMENT ===\n", k)

	for current != nil {
		if current.Left == nil {
			// Visit current node
			count++
			fmt.Printf("Visiting node %d (count = %d)\n", current.Val, count)
			if count == k {
				fmt.Printf("Found %d-th smallest element: %d\n\n", k, current.Val)
				return current.Val
			}
			current = current.Right
		} else {
			// Find predecessor
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				// Create thread
				predecessor.Right = current
				current = current.Left
			} else {
				// Remove thread and visit current
				predecessor.Right = nil
				count++
				fmt.Printf("Visiting node %d (count = %d)\n", current.Val, count)
				if count == k {
					fmt.Printf("Found %d-th smallest element: %d\n\n", k, current.Val)
					return current.Val
				}
				current = current.Right
			}
		}
	}

	fmt.Printf("Tree has fewer than %d elements\n\n", k)
	return -1
}

// ================================
// DEMONSTRATION FUNCTIONS
// ================================

// DemoMorrisTraversal demonstrates Morris traversal with detailed examples
func DemoMorrisTraversal() {
	fmt.Println("=== MORRIS TRAVERSAL ALGORITHM ===\n")

	fmt.Println("Morris Traversal is a tree traversal technique that achieves:")
	fmt.Println("✓ O(n) time complexity")
	fmt.Println("✓ O(1) space complexity (no recursion or stack)")
	fmt.Println("✓ Uses concept of threaded binary trees")
	fmt.Println("✓ Temporarily modifies tree structure during traversal")
	fmt.Println()

	// Example 1: Simple tree
	fmt.Println("=== EXAMPLE 1: Simple Binary Tree ===")
	tree1 := BuildSampleTree()
	VisualizeTree(tree1)

	fmt.Println("Expected inorder: [1, 2, 3, 4, 5, 6, 7]")
	MorrisInorderTraversal(tree1)

	// Example 2: Complex tree
	fmt.Println("=== EXAMPLE 2: Complex Binary Tree ===")
	tree2 := BuildComplexTree()
	VisualizeTree(tree2)

	fmt.Println("Expected inorder: [1, 3, 5, 6, 7, 8, 10, 12, 15, 20, 25]")
	result2 := MorrisInorderSimple(tree2)
	fmt.Printf("Morris result: %v\n\n", result2)

	// Example 3: Linear tree (worst case for recursive)
	fmt.Println("=== EXAMPLE 3: Linear Tree (Space Efficiency Demo) ===")
	tree3 := BuildLinearTree()
	VisualizeTree(tree3)

	fmt.Println("This linear tree would use O(n) space with recursive/iterative methods")
	fmt.Println("but Morris traversal still uses O(1) space!")
	result3 := MorrisInorderSimple(tree3)
	fmt.Printf("Morris result: %v\n\n", result3)

	// Performance comparison
	PerformanceComparison(tree1)
}

// DemoMorrisApplications shows practical applications
func DemoMorrisApplications() {
	fmt.Println("=== PRACTICAL APPLICATIONS ===\n")

	// Application 1: BST Validation
	fmt.Println("1. BST VALIDATION")
	bst := BuildSampleTree() // This is a valid BST
	fmt.Println("Valid BST:")
	VisualizeTree(bst)
	isValidBST := MorrisTraversalValidator(bst)
	fmt.Printf("Is valid BST: %v\n\n", isValidBST)

	// Invalid BST
	invalidBST := BuildSampleTree()
	invalidBST.Left.Right.Val = 10 // Make it invalid (10 > 4 but in left subtree)
	fmt.Println("Invalid BST:")
	VisualizeTree(invalidBST)
	isValid := MorrisTraversalValidator(invalidBST)
	fmt.Printf("Is valid BST: %v\n\n", isValid)

	// Application 2: Kth smallest element
	fmt.Println("2. FINDING KTH SMALLEST ELEMENT")
	tree := BuildComplexTree()
	fmt.Println("Tree for kth smallest search:")
	VisualizeTree(tree)

	k := 5
	kthElement := KthSmallestElementMorris(tree, k)
	fmt.Printf("Result: %d\n", kthElement)

	// Application 3: Different traversal orders
	fmt.Println("3. DIFFERENT TRAVERSAL ORDERS")
	tree4 := BuildSampleTree()
	fmt.Println("Sample tree:")
	VisualizeTree(tree4)

	inorderResult := MorrisInorderSimple(tree4)
	preorderResult := MorrisPreorderTraversal(tree4)

	fmt.Printf("Inorder:  %v\n", inorderResult)
	fmt.Printf("Preorder: %v\n\n", preorderResult)

	// Application 4: Memory-constrained environments
	fmt.Println("4. MEMORY-CONSTRAINED ENVIRONMENTS")
	fmt.Println("Morris Traversal is perfect for:")
	fmt.Println("• Embedded systems with limited memory")
	fmt.Println("• Large trees that don't fit in memory")
	fmt.Println("• Systems where stack overflow is a concern")
	fmt.Println("• Streaming tree processing")
	fmt.Println()

	fmt.Println("=== ALGORITHM CHARACTERISTICS ===")
	fmt.Println("Time Complexity: O(n)")
	fmt.Println("- Each edge is traversed at most 3 times")
	fmt.Println("- Constant work per node")
	fmt.Println()
	fmt.Println("Space Complexity: O(1)")
	fmt.Println("- No recursion stack")
	fmt.Println("- No explicit stack data structure")
	fmt.Println("- Only uses a few pointer variables")
	fmt.Println()
	fmt.Println("Key Insight:")
	fmt.Println("- Uses 'threading' to remember path back to parent")
	fmt.Println("- Right pointer of inorder predecessor points to current node")
	fmt.Println("- Threads are created and removed during traversal")
	fmt.Println("- Tree structure is restored after traversal")
	fmt.Println()
}
