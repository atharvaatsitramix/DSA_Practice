package main

import (
	"fmt"
	"strings"
)

// ================================
// TRIE (PREFIX TREE) DATA STRUCTURE
// ================================

// TrieNode represents a node in the Trie
type TrieNode struct {
	children map[rune]*TrieNode // Map of character to child node
	isEnd    bool               // Marks end of a word
	count    int                // Number of words ending at this node
}

// NewTrieNode creates a new Trie node
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
		count:    0,
	}
}

// Trie represents the Trie data structure
type Trie struct {
	root *TrieNode
	size int // Total number of words in the Trie
}

// NewTrie creates a new Trie
func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
		size: 0,
	}
}

// ================================
// CORE OPERATIONS: INSERT & SEARCH
// ================================

// Insert adds a word to the Trie with detailed tracing
func (t *Trie) Insert(word string) {
	fmt.Printf("=== INSERTING WORD: '%s' ===\n", word)

	current := t.root
	fmt.Printf("Starting at root node\n")

	for i, char := range word {
		fmt.Printf("Step %d: Processing character '%c'\n", i+1, char)

		if current.children[char] == nil {
			fmt.Printf("  Character '%c' not found, creating new node\n", char)
			current.children[char] = NewTrieNode()
		} else {
			fmt.Printf("  Character '%c' already exists, following existing path\n", char)
		}

		current = current.children[char]
		fmt.Printf("  Moved to node for character '%c'\n", char)
	}

	if !current.isEnd {
		fmt.Printf("Marking end of word '%s'\n", word)
		current.isEnd = true
		current.count = 1
		t.size++
		fmt.Printf("New word added! Total words in Trie: %d\n", t.size)
	} else {
		fmt.Printf("Word '%s' already exists, incrementing count\n", word)
		current.count++
	}

	fmt.Printf("Insert complete!\n\n")
}

// InsertSimple adds a word to the Trie without tracing
func (t *Trie) InsertSimple(word string) {
	current := t.root

	for _, char := range word {
		if current.children[char] == nil {
			current.children[char] = NewTrieNode()
		}
		current = current.children[char]
	}

	if !current.isEnd {
		current.isEnd = true
		current.count = 1
		t.size++
	} else {
		current.count++
	}
}

// Search looks for a word in the Trie with detailed tracing
func (t *Trie) Search(word string) bool {
	fmt.Printf("=== SEARCHING FOR WORD: '%s' ===\n", word)

	current := t.root
	fmt.Printf("Starting search at root node\n")

	for i, char := range word {
		fmt.Printf("Step %d: Looking for character '%c'\n", i+1, char)

		if current.children[char] == nil {
			fmt.Printf("  Character '%c' not found! Word does not exist.\n", char)
			fmt.Printf("Search result: FALSE\n\n")
			return false
		}

		fmt.Printf("  Character '%c' found, moving to next node\n", char)
		current = current.children[char]
	}

	if current.isEnd {
		fmt.Printf("Reached end of word '%s' and isEnd = true\n", word)
		fmt.Printf("Word count: %d\n", current.count)
		fmt.Printf("Search result: TRUE\n\n")
		return true
	}

	fmt.Printf("Reached end of traversal but isEnd = false\n")
	fmt.Printf("'%s' is a prefix but not a complete word\n", word)
	fmt.Printf("Search result: FALSE\n\n")
	return false
}

// SearchSimple looks for a word in the Trie without tracing
func (t *Trie) SearchSimple(word string) bool {
	current := t.root

	for _, char := range word {
		if current.children[char] == nil {
			return false
		}
		current = current.children[char]
	}

	return current.isEnd
}

// ================================
// ADDITIONAL OPERATIONS
// ================================

// StartsWith checks if any word in the Trie starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	fmt.Printf("=== CHECKING PREFIX: '%s' ===\n", prefix)

	current := t.root

	for i, char := range prefix {
		fmt.Printf("Step %d: Looking for character '%c'\n", i+1, char)

		if current.children[char] == nil {
			fmt.Printf("  Character '%c' not found! No words with this prefix.\n", char)
			fmt.Printf("Prefix check result: FALSE\n\n")
			return false
		}

		fmt.Printf("  Character '%c' found, continuing...\n", char)
		current = current.children[char]
	}

	fmt.Printf("All characters of prefix '%s' found in Trie\n", prefix)
	fmt.Printf("Prefix check result: TRUE\n\n")
	return true
}

// GetWordsWithPrefix returns all words that start with the given prefix
func (t *Trie) GetWordsWithPrefix(prefix string) []string {
	fmt.Printf("=== FINDING WORDS WITH PREFIX: '%s' ===\n", prefix)

	// First, navigate to the prefix
	current := t.root
	for _, char := range prefix {
		if current.children[char] == nil {
			fmt.Printf("Prefix '%s' not found in Trie\n\n", prefix)
			return []string{}
		}
		current = current.children[char]
	}

	// Collect all words starting from this node
	var words []string
	t.collectWords(current, prefix, &words)

	fmt.Printf("Found %d words with prefix '%s': %v\n\n", len(words), prefix, words)
	return words
}

// collectWords is a helper function for DFS traversal
func (t *Trie) collectWords(node *TrieNode, currentWord string, words *[]string) {
	if node.isEnd {
		for i := 0; i < node.count; i++ {
			*words = append(*words, currentWord)
		}
	}

	for char, child := range node.children {
		t.collectWords(child, currentWord+string(char), words)
	}
}

// Delete removes a word from the Trie
func (t *Trie) Delete(word string) bool {
	fmt.Printf("=== DELETING WORD: '%s' ===\n", word)

	return t.deleteHelper(t.root, word, 0)
}

// deleteHelper is a recursive helper for deletion
func (t *Trie) deleteHelper(node *TrieNode, word string, index int) bool {
	if index == len(word) {
		// Reached end of word
		if !node.isEnd {
			fmt.Printf("Word '%s' not found in Trie\n\n", word)
			return false
		}

		if node.count > 1 {
			fmt.Printf("Word '%s' has count > 1, decrementing count\n", word)
			node.count--
			return false // Don't delete node
		}

		node.isEnd = false
		node.count = 0
		t.size--
		fmt.Printf("Word '%s' deleted! Remaining words: %d\n\n", word, t.size)

		// Return true if node has no children (can be deleted)
		return len(node.children) == 0
	}

	char := rune(word[index])
	child := node.children[char]

	if child == nil {
		fmt.Printf("Word '%s' not found in Trie\n\n", word)
		return false
	}

	shouldDeleteChild := t.deleteHelper(child, word, index+1)

	if shouldDeleteChild {
		delete(node.children, char)
		// Return true if current node can also be deleted
		return !node.isEnd && len(node.children) == 0
	}

	return false
}

// ================================
// VISUALIZATION AND UTILITY
// ================================

// PrintTrie displays the Trie structure
func (t *Trie) PrintTrie() {
	fmt.Println("=== TRIE STRUCTURE ===")
	fmt.Printf("Total words: %d\n", t.size)
	fmt.Println("Structure:")
	t.printTrieHelper(t.root, "", "")
	fmt.Println()
}

// printTrieHelper is a recursive helper for printing
func (t *Trie) printTrieHelper(node *TrieNode, prefix, indent string) {
	if node.isEnd {
		fmt.Printf("%s'%s' (count: %d) ✓\n", indent, prefix, node.count)
	}

	chars := make([]rune, 0, len(node.children))
	for char := range node.children {
		chars = append(chars, char)
	}

	for i, char := range chars {
		isLast := i == len(chars)-1
		newIndent := indent
		if isLast {
			fmt.Printf("%s└── %c\n", indent, char)
			newIndent += "    "
		} else {
			fmt.Printf("%s├── %c\n", indent, char)
			newIndent += "│   "
		}
		t.printTrieHelper(node.children[char], prefix+string(char), newIndent)
	}
}

// GetAllWords returns all words in the Trie
func (t *Trie) GetAllWords() []string {
	var words []string
	t.collectWords(t.root, "", &words)
	return words
}

// Size returns the number of words in the Trie
func (t *Trie) Size() int {
	return t.size
}

// IsEmpty checks if the Trie is empty
func (t *Trie) IsEmpty() bool {
	return t.size == 0
}

// ================================
// ADVANCED APPLICATIONS
// ================================

// AutoComplete provides word suggestions based on prefix
type AutoComplete struct {
	trie           *Trie
	maxSuggestions int
}

// NewAutoComplete creates a new autocomplete system
func NewAutoComplete(maxSuggestions int) *AutoComplete {
	return &AutoComplete{
		trie:           NewTrie(),
		maxSuggestions: maxSuggestions,
	}
}

// AddWord adds a word to the autocomplete dictionary
func (ac *AutoComplete) AddWord(word string) {
	ac.trie.InsertSimple(strings.ToLower(word))
}

// GetSuggestions returns word suggestions for a prefix
func (ac *AutoComplete) GetSuggestions(prefix string) []string {
	prefix = strings.ToLower(prefix)
	words := ac.trie.GetWordsWithPrefix(prefix)

	// Limit suggestions
	if len(words) > ac.maxSuggestions {
		words = words[:ac.maxSuggestions]
	}

	return words
}

// SpellChecker provides spell checking functionality
type SpellChecker struct {
	trie *Trie
}

// NewSpellChecker creates a new spell checker
func NewSpellChecker() *SpellChecker {
	return &SpellChecker{
		trie: NewTrie(),
	}
}

// AddToDictionary adds a word to the spell checker dictionary
func (sc *SpellChecker) AddToDictionary(word string) {
	sc.trie.InsertSimple(strings.ToLower(word))
}

// CheckSpelling checks if a word is spelled correctly
func (sc *SpellChecker) CheckSpelling(word string) bool {
	return sc.trie.SearchSimple(strings.ToLower(word))
}

// GetSuggestions provides spelling suggestions (simplified)
func (sc *SpellChecker) GetSuggestions(word string) []string {
	word = strings.ToLower(word)

	// Try removing one character
	suggestions := []string{}

	for i := 0; i < len(word); i++ {
		candidate := word[:i] + word[i+1:]
		if candidate != "" && sc.trie.SearchSimple(candidate) {
			suggestions = append(suggestions, candidate)
		}
	}

	// Try prefix matching
	if len(suggestions) < 5 {
		prefixSuggestions := sc.trie.GetWordsWithPrefix(word[:len(word)/2])
		for _, suggestion := range prefixSuggestions {
			if len(suggestions) >= 5 {
				break
			}
			suggestions = append(suggestions, suggestion)
		}
	}

	return suggestions
}

// ================================
// DEMONSTRATION FUNCTIONS
// ================================

// DemoTrieBasics demonstrates basic Trie operations
func DemoTrieBasics() {
	fmt.Println("=== TRIE DATA STRUCTURE BASICS ===\n")

	fmt.Println("A Trie (Prefix Tree) is a tree-like data structure that:")
	fmt.Println("✓ Stores strings efficiently")
	fmt.Println("✓ Enables fast prefix-based searches")
	fmt.Println("✓ Supports autocomplete and spell checking")
	fmt.Println("✓ Has O(m) time complexity for search/insert (m = string length)")
	fmt.Println()

	// Create a new Trie
	trie := NewTrie()

	// Example 1: Basic Insert and Search
	fmt.Println("=== EXAMPLE 1: Basic Operations ===")

	words := []string{"cat", "cats", "dog", "doggy", "car", "card", "care", "careful"}

	fmt.Println("Inserting words into Trie:")
	for _, word := range words {
		trie.Insert(word)
	}

	trie.PrintTrie()

	// Search examples
	fmt.Println("=== SEARCH EXAMPLES ===")
	searchWords := []string{"cat", "car", "care", "caring", "dog", "do"}

	for _, word := range searchWords {
		found := trie.Search(word)
		fmt.Printf("'%s' found: %v\n", word, found)
	}
	fmt.Println()
}

// DemoTrieAdvanced demonstrates advanced Trie operations
func DemoTrieAdvanced() {
	fmt.Println("=== ADVANCED TRIE OPERATIONS ===\n")

	trie := NewTrie()

	// Build dictionary
	dictionary := []string{
		"apple", "app", "application", "apply", "appreciate",
		"banana", "band", "bandana", "ban",
		"cat", "cats", "caterpillar", "catch",
		"dog", "doggy", "dogs", "dogma",
	}

	fmt.Println("Building dictionary...")
	for _, word := range dictionary {
		trie.InsertSimple(word)
	}

	fmt.Printf("Dictionary loaded with %d words\n\n", trie.Size())

	// Prefix operations
	fmt.Println("=== PREFIX OPERATIONS ===")
	prefixes := []string{"app", "cat", "dog", "xyz"}

	for _, prefix := range prefixes {
		fmt.Printf("Prefix '%s':\n", prefix)
		fmt.Printf("  Exists: %v\n", trie.StartsWith(prefix))
		words := trie.GetWordsWithPrefix(prefix)
		fmt.Printf("  Words: %v\n\n", words)
	}

	// Deletion examples
	fmt.Println("=== DELETION EXAMPLES ===")
	deleteWords := []string{"app", "cats", "nonexistent"}

	for _, word := range deleteWords {
		fmt.Printf("Deleting '%s':\n", word)
		deleted := trie.Delete(word)
		fmt.Printf("  Success: %v\n", deleted)
		fmt.Printf("  Remaining size: %d\n\n", trie.Size())
	}
}

// DemoAutoComplete demonstrates autocomplete functionality
func DemoAutoComplete() {
	fmt.Println("=== AUTOCOMPLETE SYSTEM ===\n")

	ac := NewAutoComplete(5) // Maximum 5 suggestions

	// Load common words
	commonWords := []string{
		"hello", "help", "helpful", "hero", "health",
		"world", "work", "word", "worry", "worth",
		"programming", "program", "progress", "project", "problem",
		"computer", "compute", "company", "complete", "compare",
	}

	fmt.Println("Loading autocomplete dictionary...")
	for _, word := range commonWords {
		ac.AddWord(word)
	}

	fmt.Printf("Dictionary loaded with %d unique words\n\n", len(commonWords))

	// Test autocomplete
	testPrefixes := []string{"he", "wo", "pro", "com", "xyz"}

	for _, prefix := range testPrefixes {
		suggestions := ac.GetSuggestions(prefix)
		fmt.Printf("Autocomplete for '%s': %v\n", prefix, suggestions)
	}
	fmt.Println()
}

// DemoSpellChecker demonstrates spell checking functionality
func DemoSpellChecker() {
	fmt.Println("=== SPELL CHECKER SYSTEM ===\n")

	sc := NewSpellChecker()

	// Load dictionary
	dictionary := []string{
		"hello", "world", "computer", "programming", "algorithm",
		"structure", "search", "insert", "delete", "traverse",
		"efficiency", "complexity", "optimization", "performance",
	}

	fmt.Println("Loading spell checker dictionary...")
	for _, word := range dictionary {
		sc.AddToDictionary(word)
	}

	fmt.Printf("Dictionary loaded with %d words\n\n", len(dictionary))

	// Test spell checking
	testWords := []string{
		"hello",      // correct
		"wrold",      // misspelled (world)
		"algoritm",   // misspelled (algorithm)
		"computer",   // correct
		"programing", // misspelled (programming)
		"xyz",        // not in dictionary
	}

	for _, word := range testWords {
		isCorrect := sc.CheckSpelling(word)
		fmt.Printf("Word: '%s'\n", word)
		fmt.Printf("  Correct spelling: %v\n", isCorrect)

		if !isCorrect {
			suggestions := sc.GetSuggestions(word)
			fmt.Printf("  Suggestions: %v\n", suggestions)
		}
		fmt.Println()
	}
}

// DemoTrieComplexity demonstrates Trie complexity characteristics
func DemoTrieComplexity() {
	fmt.Println("=== COMPLEXITY ANALYSIS ===\n")

	fmt.Println("Time Complexity:")
	fmt.Println("- Insert: O(m) where m = length of word")
	fmt.Println("- Search: O(m) where m = length of word")
	fmt.Println("- Delete: O(m) where m = length of word")
	fmt.Println("- Prefix search: O(p + n) where p = prefix length, n = results")
	fmt.Println()

	fmt.Println("Space Complexity:")
	fmt.Println("- O(ALPHABET_SIZE * N * M) in worst case")
	fmt.Println("- Where N = number of words, M = average length")
	fmt.Println("- Much more efficient when words share prefixes")
	fmt.Println()

	fmt.Println("Advantages:")
	fmt.Println("✓ Fast prefix-based operations")
	fmt.Println("✓ No hash collisions")
	fmt.Println("✓ Lexicographically sorted output")
	fmt.Println("✓ Excellent for autocomplete/spell check")
	fmt.Println()

	fmt.Println("Disadvantages:")
	fmt.Println("✗ High memory usage for sparse datasets")
	fmt.Println("✗ Cache performance issues with deep trees")
	fmt.Println("✗ More complex than hash tables for simple lookups")
	fmt.Println()

	// Demonstrate with example
	fmt.Println("=== SPACE EFFICIENCY EXAMPLE ===")

	trie := NewTrie()

	// Words with common prefixes (efficient)
	efficientWords := []string{
		"programming", "program", "programmer", "programs",
		"application", "apply", "apple", "applicable",
	}

	fmt.Println("Inserting words with common prefixes:")
	for _, word := range efficientWords {
		fmt.Printf("  %s\n", word)
		trie.InsertSimple(word)
	}

	fmt.Println("\nTrie structure (notice shared prefixes):")
	trie.PrintTrie()
}
