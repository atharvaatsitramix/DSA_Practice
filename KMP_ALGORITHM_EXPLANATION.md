# KMP (Knuth-Morris-Pratt) Algorithm in Go

## Overview

**KMP (Knuth-Morris-Pratt)** is an efficient string pattern matching algorithm that searches for occurrences of a pattern within a text string. It was developed by Donald Knuth, James H. Morris, and Vaughan Pratt in 1977.

### Key Characteristics:
- **Time Complexity**: O(n + m) where n = text length, m = pattern length
- **Space Complexity**: O(m) for the LPS (failure function) table
- **Key Advantage**: Never re-examines characters in the text
- **Applications**: Text search, DNA analysis, network security, data compression

## The Problem: String Pattern Matching

Given:
- **Text** T of length n
- **Pattern** P of length m

**Goal**: Find all occurrences of pattern P in text T.

### Example:
```
Text:    "ABABDABACDABABCABAB"
Pattern: "ABABCABAB"
Goal:    Find where pattern appears in text
```

## Naive Approach vs KMP

### Naive (Brute Force) Approach:
```
For each position i in text:
    Compare pattern with text starting at position i
    If mismatch occurs, move to position i+1 and start over
```

**Problem**: In worst case, compares every character multiple times.
- **Time Complexity**: O(nm)
- **Example**: Pattern "AAAB" in text "AAAAAAB" requires many re-examinations

### KMP Approach:
```
1. Preprocess pattern to build LPS table
2. Use LPS table to skip unnecessary comparisons
3. Never go backward in the text
```

**Advantage**: Each character in text is examined at most once.
- **Time Complexity**: O(n + m)

## Core Concept: LPS Table

**LPS** = **Longest Proper Prefix which is also Suffix**

### What is LPS?
For each position i in the pattern, LPS[i] is the length of the longest proper prefix of pattern[0...i] which is also a proper suffix of pattern[0...i].

- **Proper prefix**: A prefix that is not equal to the string itself
- **Proper suffix**: A suffix that is not equal to the string itself

### LPS Table Examples:

#### Pattern: "ABABCABAB"
```
Index:  0 1 2 3 4 5 6 7 8
Pattern:A B A B C A B A B
LPS:    0 0 1 2 0 1 2 3 4
```

**Explanation**:
- LPS[0] = 0: Single character has no proper prefix/suffix
- LPS[1] = 0: "AB" has no matching prefix/suffix
- LPS[2] = 1: "ABA" → prefix "A" = suffix "A"
- LPS[3] = 2: "ABAB" → prefix "AB" = suffix "AB"
- LPS[4] = 0: "ABABC" has no matching prefix/suffix
- LPS[5] = 1: "ABABCA" → prefix "A" = suffix "A"
- LPS[6] = 2: "ABABCAB" → prefix "AB" = suffix "AB"
- LPS[7] = 3: "ABABCABA" → prefix "ABA" = suffix "ABA"
- LPS[8] = 4: "ABABCABAB" → prefix "ABAB" = suffix "ABAB"

#### Pattern: "AAACAAAA"
```
Index:  0 1 2 3 4 5 6 7
Pattern:A A A C A A A A
LPS:    0 1 2 0 1 2 3 3
```

## How KMP Works: Step-by-Step

### Phase 1: Build LPS Table

```go
func buildLPS(pattern string) []int {
    lps := make([]int, len(pattern))
    length := 0  // Length of previous longest prefix suffix
    i := 1
    
    lps[0] = 0  // First element is always 0
    
    while i < len(pattern) {
        if pattern[i] == pattern[length] {
            length++
            lps[i] = length
            i++
        } else {
            if length != 0 {
                length = lps[length-1]  // Key insight: backtrack
            } else {
                lps[i] = 0
                i++
            }
        }
    }
    
    return lps
}
```

### Phase 2: Search Using LPS Table

```go
func KMPSearch(text, pattern string) []int {
    lps := buildLPS(pattern)
    matches := []int{}
    
    i := 0  // Index for text
    j := 0  // Index for pattern
    
    while i < len(text) {
        if text[i] == pattern[j] {
            i++
            j++
        }
        
        if j == len(pattern) {
            matches = append(matches, i-j)  // Found match
            j = lps[j-1]                    // Continue searching
        } else if i < len(text) && text[i] != pattern[j] {
            if j != 0 {
                j = lps[j-1]  // Use LPS to avoid re-checking
            } else {
                i++
            }
        }
    }
    
    return matches
}
```

## Detailed Example Walkthrough

### Pattern: "ABABCABAB", Text: "ABABDABACDABABCABAB"

#### Step 1: Build LPS Table
```
Pattern: A B A B C A B A B
Index:   0 1 2 3 4 5 6 7 8
LPS:     0 0 1 2 0 1 2 3 4
```

#### Step 2: Search Process
```
Text:    A B A B D A B A C D A B A B C A B A B
Pattern: A B A B C A B A B
         ^
i=0,j=0: A=A ✓, advance both → i=1,j=1

Text:    A B A B D A B A C D A B A B C A B A B
Pattern: A B A B C A B A B
           ^
i=1,j=1: B=B ✓, advance both → i=2,j=2

Text:    A B A B D A B A C D A B A B C A B A B
Pattern: A B A B C A B A B
             ^
i=2,j=2: A=A ✓, advance both → i=3,j=3

Text:    A B A B D A B A C D A B A B C A B A B
Pattern: A B A B C A B A B
               ^
i=3,j=3: B=B ✓, advance both → i=4,j=4

Text:    A B A B D A B A C D A B A B C A B A B
Pattern: A B A B C A B A B
                 ^
i=4,j=4: D≠C ✗, j=LPS[4-1]=LPS[3]=2

Text:    A B A B D A B A C D A B A B C A B A B
Pattern:     A B A B C A B A B
                 ^
i=4,j=2: D≠A ✗, j=LPS[2-1]=LPS[1]=0

Text:    A B A B D A B A C D A B A B C A B A B
Pattern:         A B A B C A B A B
                 ^
i=4,j=0: D≠A ✗, j=0 so advance i → i=5,j=0

... (continue until match found at position 10)
```

## Why LPS Table Works

### Key Insight:
When a mismatch occurs at position j, we know:
1. Characters pattern[0...j-1] matched text[i-j...i-1]
2. We can use the LPS value to determine how many characters we can skip

### Example:
```
Text:    ...A B A B C...
Pattern:   A B A B D...
           0 1 2 3 4
```

When mismatch occurs at j=4:
- We know "ABAB" matched
- LPS[3] = 2 means first 2 chars ("AB") equal last 2 chars ("AB")
- So we can restart matching from j=2, not j=0

## Complexity Analysis

### Time Complexity:

#### LPS Table Construction: O(m)
- Each character in pattern is processed at most twice
- Inner while loop moves `length` backward, but `length` can only increase m times total

#### Searching: O(n)
- Each character in text is examined at most once
- `i` never decreases, only increases
- `j` can decrease via LPS, but total decreases ≤ total increases ≤ n

#### Total: O(n + m)

### Space Complexity: O(m)
- Only need LPS table of size m

### Comparison:

| Algorithm | Time Complexity | Space Complexity | Re-examinations |
|-----------|----------------|------------------|-----------------|
| **Naive** | O(nm) | O(1) | Yes |
| **KMP** | O(n + m) | O(m) | No |

## Real-World Applications

### 1. **Text Editors**
- **Find/Replace** functionality
- **Syntax highlighting** (finding keywords)
- **Auto-completion** suggestions

```go
// Find all occurrences of "TODO" in code
matcher := NewKMPMatcher("TODO")
positions := matcher.Search(sourceCode)
```

### 2. **Bioinformatics**
- **DNA sequence analysis**
- **Protein pattern matching**
- **Gene finding**

```go
// Find restriction enzyme sites in DNA
dnaSequence := "ATCGATCGATCGTAGCTAGCT"
restrictionSite := "GAATTC"  // EcoRI site
matcher := NewKMPMatcher(restrictionSite)
sites := matcher.Search(dnaSequence)
```

### 3. **Network Security**
- **Intrusion detection** systems
- **Virus pattern matching**
- **Network packet inspection**

```go
// Detect malicious patterns in network traffic
malwareSignatures := []string{"VIRUS", "MALWARE", "EXPLOIT"}
for _, signature := range malwareSignatures {
    matcher := NewKMPMatcher(signature)
    if len(matcher.Search(packetData)) > 0 {
        alert("Malware detected: " + signature)
    }
}
```

### 4. **Data Compression**
- **Lempel-Ziv** algorithms use pattern matching
- **Dictionary-based** compression
- **Redundancy detection**

### 5. **Web Search**
- **Full-text search** engines
- **Log file analysis**
- **Document processing**

### 6. **Programming Tools**
- **Code analysis** tools
- **Refactoring** utilities
- **Plagiarism detection**

## Advanced Variations

### 1. **Multiple Pattern Search**
```go
type MultiPatternKMP struct {
    matchers []*KMPMatcher
    patterns []string
}

func (mp *MultiPatternKMP) SearchAll(text string) map[string][]int {
    results := make(map[string][]int)
    for i, matcher := range mp.matchers {
        results[mp.patterns[i]] = matcher.Search(text)
    }
    return results
}
```

### 2. **Case-Insensitive Search**
```go
func CaseInsensitiveKMP(text, pattern string) []int {
    return KMPSearchSimple(strings.ToLower(text), strings.ToLower(pattern))
}
```

### 3. **Approximate Matching**
- **Boyer-Moore-Horspool** algorithm
- **Rabin-Karp** with rolling hash
- **Aho-Corasick** for multiple patterns

## Implementation Tips

### 1. **Boundary Checks**
```go
if len(pattern) == 0 {
    return []int{}  // Handle empty pattern
}
if len(text) < len(pattern) {
    return []int{}  // Pattern longer than text
}
```

### 2. **Unicode Support**
```go
// Convert to runes for proper Unicode handling
textRunes := []rune(text)
patternRunes := []rune(pattern)
```

### 3. **Memory Optimization**
```go
// For one-time searches, don't store LPS table
func KMPSearchOnce(text, pattern string) int {
    lps := buildLPS(pattern)
    // ... search logic
    // LPS table is automatically garbage collected
}
```

## Common Mistakes

### 1. **Incorrect LPS Construction**
```go
// Wrong: Not handling the backtrack properly
if pattern[i] != pattern[length] {
    lps[i] = 0  // Should use LPS[length-1] if length > 0
    i++
}

// Correct:
if pattern[i] != pattern[length] {
    if length != 0 {
        length = lps[length-1]  // Backtrack
    } else {
        lps[i] = 0
        i++
    }
}
```

### 2. **Index Out of Bounds**
```go
// Wrong: Not checking bounds
if text[i] == pattern[j] {
    // What if i >= len(text)?
}

// Correct:
if i < len(text) && text[i] == pattern[j] {
    // Safe comparison
}
```

### 3. **Infinite Loops**
```go
// Wrong: Can cause infinite loop
while i < len(text) {
    if text[i] != pattern[j] && j > 0 {
        j = lps[j-1]
        // Not incrementing i can cause infinite loop
    }
}

// Correct: Ensure progress
while i < len(text) {
    if text[i] != pattern[j] {
        if j > 0 {
            j = lps[j-1]
        } else {
            i++  // Ensure i advances
        }
    }
}
```

## When to Use KMP

### ✅ **Perfect for:**
- **Large text** with **small to medium patterns**
- **Multiple searches** with same pattern
- **No preprocessing** budget constraints
- **Streaming data** (never go backward)

### ❌ **Consider alternatives for:**
- **Very short texts** (overhead not worth it)
- **Single-character patterns** (simple loop is faster)
- **Multiple patterns** (Aho-Corasick is better)
- **Approximate matching** (use fuzzy string algorithms)

## Comparison with Other Algorithms

| Algorithm | Preprocessing | Search | Best For |
|-----------|--------------|--------|----------|
| **Naive** | O(1) | O(nm) | Very short patterns |
| **KMP** | O(m) | O(n) | General purpose |
| **Boyer-Moore** | O(m + σ) | O(n/m) avg | Large alphabets |
| **Rabin-Karp** | O(m) | O(n) avg | Multiple patterns |
| **Aho-Corasick** | O(Σm) | O(n + z) | Many patterns |

Where:
- n = text length, m = pattern length
- σ = alphabet size, z = number of matches
- Σm = sum of all pattern lengths

## Practice Problems

### Beginner:
1. **Implement basic KMP** search
2. **Find first occurrence** of pattern
3. **Count total occurrences** of pattern

### Intermediate:
4. **Case-insensitive search**
5. **Search in circular array**
6. **Find all palindromic substrings** (using KMP variant)

### Advanced:
7. **Multiple pattern search**
8. **Compressed pattern matching**
9. **2D pattern matching** in grids

## Optimizations and Variants

### 1. **Space Optimization**
```go
// Don't store entire LPS table for one-time use
func KMPSearchStreaming(text, pattern string) []int {
    // Compute LPS on-the-fly or use less memory
}
```

### 2. **Early Termination**
```go
// Stop after finding first match
func KMPFindFirst(text, pattern string) int {
    // Return immediately when first match found
}
```

### 3. **Parallel Processing**
```go
// Search multiple patterns in parallel
func ParallelKMPSearch(text string, patterns []string) map[string][]int {
    // Use goroutines for concurrent pattern matching
}
```

## Key Insights

1. **Linear Time**: KMP achieves optimal O(n + m) time complexity for single pattern search

2. **No Backtracking**: The algorithm never re-examines characters in the text, making it suitable for streaming data

3. **Failure Function**: The LPS table encodes information about the pattern's internal structure

4. **Practical Efficiency**: Despite theoretical optimality, simpler algorithms might be faster for small inputs due to lower constant factors

5. **Foundation for Advanced Algorithms**: KMP concepts are used in many other string algorithms like Aho-Corasick

The KMP algorithm is a beautiful example of how preprocessing can dramatically improve algorithmic efficiency! 🚀 