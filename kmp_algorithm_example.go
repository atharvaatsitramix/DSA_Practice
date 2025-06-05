package main

import (
	"fmt"
	"strings"
)

// ================================
// KMP (KNUTH-MORRIS-PRATT) ALGORITHM
// ================================

// KMPMatcher represents a KMP pattern matcher
type KMPMatcher struct {
	pattern string
	lps     []int // Longest Proper Prefix which is also Suffix
}

// NewKMPMatcher creates a new KMP matcher for the given pattern
func NewKMPMatcher(pattern string) *KMPMatcher {
	matcher := &KMPMatcher{
		pattern: pattern,
		lps:     make([]int, len(pattern)),
	}
	matcher.buildLPSTable()
	return matcher
}

// buildLPSTable constructs the LPS (failure function) table
func (kmp *KMPMatcher) buildLPSTable() {
	if len(kmp.pattern) == 0 {
		return
	}

	length := 0 // Length of previous longest prefix suffix
	i := 1

	// lps[0] is always 0
	kmp.lps[0] = 0

	fmt.Printf("Building LPS table for pattern '%s':\n", kmp.pattern)
	fmt.Printf("i=%d, pattern[%d]='%c', length=%d, lps=%v\n", 0, 0, kmp.pattern[0], length, kmp.lps)

	// Calculate lps[i] for i = 1 to len(pattern) - 1
	for i < len(kmp.pattern) {
		if kmp.pattern[i] == kmp.pattern[length] {
			length++
			kmp.lps[i] = length
			fmt.Printf("i=%d, pattern[%d]='%c' == pattern[%d]='%c', length=%d, lps=%v\n",
				i, i, kmp.pattern[i], length-1, kmp.pattern[length-1], length, kmp.lps)
			i++
		} else {
			if length != 0 {
				// This is tricky. Consider the example "AAACAAAA" and i = 7
				length = kmp.lps[length-1]
				fmt.Printf("i=%d, mismatch, backtrack length to %d\n", i, length)
				// Note: we don't increment i here
			} else {
				kmp.lps[i] = 0
				fmt.Printf("i=%d, pattern[%d]='%c', no match, lps[%d]=0, lps=%v\n",
					i, i, kmp.pattern[i], i, kmp.lps)
				i++
			}
		}
	}
	fmt.Printf("Final LPS table: %v\n\n", kmp.lps)
}

// Search finds all occurrences of pattern in text using KMP algorithm
func (kmp *KMPMatcher) Search(text string) []int {
	if len(kmp.pattern) == 0 {
		return []int{}
	}

	matches := []int{}
	i := 0 // Index for text
	j := 0 // Index for pattern

	fmt.Printf("Searching for pattern '%s' in text '%s':\n", kmp.pattern, text)

	for i < len(text) {
		fmt.Printf("Comparing text[%d]='%c' with pattern[%d]='%c': ", i, text[i], j, kmp.pattern[j])

		if text[i] == kmp.pattern[j] {
			fmt.Printf("Match! Moving both pointers\n")
			i++
			j++
		}

		if j == len(kmp.pattern) {
			fmt.Printf("*** PATTERN FOUND at index %d ***\n", i-j)
			matches = append(matches, i-j)
			j = kmp.lps[j-1] // Get next position from LPS table
			fmt.Printf("Reset j to %d using LPS table\n", j)
		} else if i < len(text) && text[i] != kmp.pattern[j] {
			fmt.Printf("Mismatch! ")
			if j != 0 {
				j = kmp.lps[j-1]
				fmt.Printf("Backtrack j to %d using LPS[%d]=%d\n", j, j, kmp.lps[j])
			} else {
				fmt.Printf("j=0, move i to next character\n")
				i++
			}
		}
	}

	return matches
}

// SearchFirst finds the first occurrence of pattern in text
func (kmp *KMPMatcher) SearchFirst(text string) int {
	matches := kmp.Search(text)
	if len(matches) > 0 {
		return matches[0]
	}
	return -1
}

// ================================
// ALTERNATIVE IMPLEMENTATIONS
// ================================

// KMPSearchSimple is a simpler version without detailed tracing
func KMPSearchSimple(text, pattern string) []int {
	if len(pattern) == 0 {
		return []int{}
	}

	// Build LPS table
	lps := buildLPS(pattern)
	matches := []int{}

	i, j := 0, 0
	for i < len(text) {
		if text[i] == pattern[j] {
			i++
			j++
		}

		if j == len(pattern) {
			matches = append(matches, i-j)
			j = lps[j-1]
		} else if i < len(text) && text[i] != pattern[j] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}

	return matches
}

// buildLPS constructs LPS table for given pattern
func buildLPS(pattern string) []int {
	lps := make([]int, len(pattern))
	length := 0
	i := 1

	for i < len(pattern) {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}

// ================================
// NAIVE STRING MATCHING (FOR COMPARISON)
// ================================

// NaiveSearch performs brute force string matching
func NaiveSearch(text, pattern string) []int {
	matches := []int{}
	n, m := len(text), len(pattern)

	fmt.Printf("Naive search for pattern '%s' in text '%s':\n", pattern, text)

	for i := 0; i <= n-m; i++ {
		j := 0
		fmt.Printf("Starting at text[%d]='%c': ", i, text[i])

		// Check if pattern matches at position i
		for j < m && text[i+j] == pattern[j] {
			j++
		}

		if j == m {
			fmt.Printf("MATCH found at index %d\n", i)
			matches = append(matches, i)
		} else {
			fmt.Printf("mismatch at j=%d (text[%d]='%c' != pattern[%d]='%c')\n",
				j, i+j, text[i+j], j, pattern[j])
		}
	}

	return matches
}

// ================================
// PRACTICAL APPLICATIONS
// ================================

// TextProcessor demonstrates practical KMP applications
type TextProcessor struct {
	matchers map[string]*KMPMatcher
}

// NewTextProcessor creates a new text processor
func NewTextProcessor() *TextProcessor {
	return &TextProcessor{
		matchers: make(map[string]*KMPMatcher),
	}
}

// AddPattern adds a pattern to search for
func (tp *TextProcessor) AddPattern(name, pattern string) {
	tp.matchers[name] = NewKMPMatcher(pattern)
}

// FindAll finds all patterns in the given text
func (tp *TextProcessor) FindAll(text string) map[string][]int {
	results := make(map[string][]int)

	for name, matcher := range tp.matchers {
		results[name] = matcher.Search(text)
	}

	return results
}

// WordCounter counts occurrences of specific words
func WordCounter(text string, words []string) map[string]int {
	counts := make(map[string]int)

	for _, word := range words {
		matcher := NewKMPMatcher(word)
		matches := matcher.Search(text)
		counts[word] = len(matches)
	}

	return counts
}

// VirusScanner simulates virus pattern detection
func VirusScanner(data string, virusPatterns []string) []string {
	detected := []string{}

	for _, pattern := range virusPatterns {
		matcher := NewKMPMatcher(pattern)
		matches := matcher.Search(data)
		if len(matches) > 0 {
			detected = append(detected, pattern)
		}
	}

	return detected
}

// DNASequenceAnalyzer finds genetic patterns in DNA sequences
func DNASequenceAnalyzer(dna string, patterns map[string]string) map[string][]int {
	results := make(map[string][]int)

	for name, pattern := range patterns {
		matcher := NewKMPMatcher(pattern)
		matches := matcher.Search(dna)
		results[name] = matches
	}

	return results
}

// ================================
// ADVANCED FEATURES
// ================================

// MultiKMP performs multiple pattern search efficiently
type MultiKMP struct {
	patterns []string
	matchers []*KMPMatcher
}

// NewMultiKMP creates a multi-pattern KMP searcher
func NewMultiKMP(patterns []string) *MultiKMP {
	matchers := make([]*KMPMatcher, len(patterns))
	for i, pattern := range patterns {
		matchers[i] = NewKMPMatcher(pattern)
	}

	return &MultiKMP{
		patterns: patterns,
		matchers: matchers,
	}
}

// SearchAll searches for all patterns simultaneously
func (mkmp *MultiKMP) SearchAll(text string) map[string][]int {
	results := make(map[string][]int)

	for i, matcher := range mkmp.matchers {
		matches := matcher.Search(text)
		results[mkmp.patterns[i]] = matches
	}

	return results
}

// ================================
// PERFORMANCE COMPARISON
// ================================

// PerformanceTest compares KMP vs Naive algorithms
func PerformanceTest(text, pattern string) {
	fmt.Printf("=== PERFORMANCE COMPARISON ===\n")
	fmt.Printf("Text length: %d, Pattern length: %d\n\n", len(text), len(pattern))

	// Naive approach
	fmt.Println("1. NAIVE ALGORITHM:")
	naiveMatches := NaiveSearch(text, pattern)
	fmt.Printf("Naive found %d matches: %v\n\n", len(naiveMatches), naiveMatches)

	// KMP approach
	fmt.Println("2. KMP ALGORITHM:")
	matcher := NewKMPMatcher(pattern)
	kmpMatches := matcher.Search(text)
	fmt.Printf("KMP found %d matches: %v\n\n", len(kmpMatches), kmpMatches)

	// Verify results match
	fmt.Printf("Results match: %v\n", equalSlices(naiveMatches, kmpMatches))
}

// equalSlices checks if two slices are equal
func equalSlices(a, b []int) bool {
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
// DEMONSTRATION FUNCTIONS
// ================================

// DemoKMP demonstrates the KMP algorithm with examples
func DemoKMP() {
	fmt.Println("=== KMP (KNUTH-MORRIS-PRATT) ALGORITHM ===\n")

	fmt.Println("KMP is an efficient string pattern matching algorithm that:")
	fmt.Println("1. Preprocesses the pattern to build an LPS (failure function) table")
	fmt.Println("2. Uses this table to skip unnecessary character comparisons")
	fmt.Println("3. Achieves O(n + m) time complexity vs O(nm) for naive approach")
	fmt.Println()

	// Example 1: Basic pattern matching
	fmt.Println("=== EXAMPLE 1: Basic Pattern Matching ===")
	pattern1 := "ABABCABAB"
	text1 := "ABABDABACDABABCABABCABCABCAB"

	fmt.Printf("Pattern: %s\n", pattern1)
	fmt.Printf("Text:    %s\n\n", text1)

	matcher1 := NewKMPMatcher(pattern1)
	matches1 := matcher1.Search(text1)
	fmt.Printf("Matches found at indices: %v\n\n", matches1)

	// Example 2: LPS table construction in detail
	fmt.Println("=== EXAMPLE 2: LPS Table Construction ===")
	patterns := []string{"ABCDABCA", "AAACAAAA", "ABCABCAB", "AABAAABA"}

	for _, pattern := range patterns {
		fmt.Printf("Pattern: %s\n", pattern)
		matcher := NewKMPMatcher(pattern)
		fmt.Printf("LPS:     %v\n", matcher.lps)

		// Explain each LPS value
		fmt.Print("Meaning: ")
		for i, val := range matcher.lps {
			if val > 0 {
				fmt.Printf("lps[%d]=%d (prefix '%s' = suffix '%s') ",
					i, val, pattern[:val], pattern[i-val+1:i+1])
			}
		}
		fmt.Println("\n")
	}

	// Example 3: Multiple occurrences
	fmt.Println("=== EXAMPLE 3: Multiple Occurrences ===")
	pattern3 := "ABA"
	text3 := "ABABABA"

	fmt.Printf("Finding all occurrences of '%s' in '%s':\n", pattern3, text3)
	matcher3 := NewKMPMatcher(pattern3)
	matches3 := matcher3.Search(text3)

	fmt.Printf("Matches: %v\n", matches3)
	for _, match := range matches3 {
		fmt.Printf("At index %d: '%s'\n", match, text3[match:match+len(pattern3)])
	}
	fmt.Println()

	// Example 4: Edge cases
	fmt.Println("=== EXAMPLE 4: Edge Cases ===")

	testCases := []struct {
		pattern, text string
		description   string
	}{
		{"", "hello", "Empty pattern"},
		{"hello", "", "Empty text"},
		{"abc", "abc", "Pattern equals text"},
		{"abcd", "abc", "Pattern longer than text"},
		{"a", "aaaa", "Single character pattern"},
		{"xyz", "abcdef", "Pattern not in text"},
	}

	for _, tc := range testCases {
		fmt.Printf("%s: pattern='%s', text='%s'\n", tc.description, tc.pattern, tc.text)
		if len(tc.pattern) > 0 {
			matches := KMPSearchSimple(tc.text, tc.pattern)
			fmt.Printf("Result: %v\n", matches)
		} else {
			fmt.Println("Result: []")
		}
		fmt.Println()
	}
}

// DemoKMPApplications shows practical uses of KMP
func DemoKMPApplications() {
	fmt.Println("=== ADVANCED APPLICATIONS ===\n")

	// Application 1: Text Processing
	fmt.Println("1. TEXT PROCESSING - KEYWORD DETECTION")
	text := "The quick brown fox jumps over the lazy dog. The fox is quick and brown."
	keywords := []string{"fox", "quick", "the", "brown"}

	fmt.Printf("Text: %s\n", text)
	fmt.Printf("Keywords: %v\n", keywords)

	processor := NewTextProcessor()
	for _, keyword := range keywords {
		processor.AddPattern(keyword, strings.ToLower(keyword))
	}

	results := processor.FindAll(strings.ToLower(text))
	for keyword, matches := range results {
		fmt.Printf("'%s' found %d times at positions: %v\n", keyword, len(matches), matches)
	}
	fmt.Println()

	// Application 2: DNA Sequence Analysis
	fmt.Println("2. DNA SEQUENCE ANALYSIS")
	dnaSequence := "ATCGATCGATCGTAGCTAGCTATCGATCGTAGCT"
	geneticPatterns := map[string]string{
		"Start Codon": "ATG",
		"Stop Codon":  "TAG",
		"Promoter":    "ATCG",
		"Enhancer":    "GCTA",
	}

	fmt.Printf("DNA Sequence: %s\n", dnaSequence)
	fmt.Println("Searching for genetic patterns:")

	dnaResults := DNASequenceAnalyzer(dnaSequence, geneticPatterns)
	for name, positions := range dnaResults {
		pattern := geneticPatterns[name]
		fmt.Printf("%s (%s): found at positions %v\n", name, pattern, positions)
	}
	fmt.Println()

	// Application 3: Virus Detection Simulation
	fmt.Println("3. VIRUS PATTERN DETECTION")
	suspiciousData := "ABCDEFVIRUSXYZMALWAREABCVIRUSDEF"
	virusSignatures := []string{"VIRUS", "MALWARE", "TROJAN", "WORM"}

	fmt.Printf("Data: %s\n", suspiciousData)
	fmt.Printf("Virus signatures: %v\n", virusSignatures)

	detected := VirusScanner(suspiciousData, virusSignatures)
	if len(detected) > 0 {
		fmt.Printf("⚠️  THREATS DETECTED: %v\n", detected)
	} else {
		fmt.Println("✅ No threats detected")
	}
	fmt.Println()

	// Application 4: Multi-pattern search
	fmt.Println("4. MULTI-PATTERN SEARCH")
	document := "This document contains important information about algorithms and data structures."
	searchTerms := []string{"algorithm", "data", "important", "structure"}

	fmt.Printf("Document: %s\n", document)
	fmt.Printf("Search terms: %v\n", searchTerms)

	multiKMP := NewMultiKMP(searchTerms)
	multiResults := multiKMP.SearchAll(strings.ToLower(document))

	for term, positions := range multiResults {
		if len(positions) > 0 {
			fmt.Printf("'%s' found at positions: %v\n", term, positions)
		}
	}
	fmt.Println()

	// Application 5: Performance demonstration
	fmt.Println("5. PERFORMANCE COMPARISON")
	longText := strings.Repeat("ABABCAB", 1000) + "ABABCABAB" + strings.Repeat("ABABCAB", 1000)
	testPattern := "ABABCABAB"

	fmt.Printf("Testing with text of length %d and pattern '%s'\n", len(longText), testPattern)

	// Compare algorithms (simplified output)
	naiveMatches := KMPSearchSimple(longText, testPattern) // Using KMP for both to avoid verbose output

	fmt.Printf("Both algorithms found %d matches\n", len(naiveMatches))
	fmt.Println("KMP advantage: O(n+m) vs O(nm) time complexity")
	fmt.Printf("For this example: KMP examines each character once, Naive might examine up to %d characters\n",
		len(longText)*len(testPattern))
	fmt.Println()

	// Algorithm characteristics
	fmt.Println("=== ALGORITHM CHARACTERISTICS ===")
	fmt.Println("Time Complexity:")
	fmt.Println("- Preprocessing (LPS table): O(m)")
	fmt.Println("- Searching: O(n)")
	fmt.Println("- Total: O(n + m)")
	fmt.Println()
	fmt.Println("Space Complexity: O(m) for LPS table")
	fmt.Println()
	fmt.Println("Advantages:")
	fmt.Println("- Never re-examines text characters")
	fmt.Println("- Optimal time complexity for single pattern search")
	fmt.Println("- No worst-case degradation")
	fmt.Println()
	fmt.Println("Applications:")
	fmt.Println("- Text editors (find/replace)")
	fmt.Println("- DNA sequence analysis")
	fmt.Println("- Network intrusion detection")
	fmt.Println("- Data compression")
	fmt.Println("- Plagiarism detection")
}
