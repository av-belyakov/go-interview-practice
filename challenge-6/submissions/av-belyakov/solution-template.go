// Package challenge6 contains the solution for Challenge 6.
package challenge6

import (
    "strings"
    "regexp"
)

// CountWordFrequency takes a string containing multiple words and returns
// a map where each key is a word and the value is the number of times that
// word appears in the string. The comparison is case-insensitive.
//
// Words are defined as sequences of letters and digits.
// All words are converted to lowercase before counting.
// All punctuation, spaces, and other non-alphanumeric characters are ignored.
//
// For example:
// Input: "The quick brown fox jumps over the lazy dog."
// Output: map[string]int{"the": 2, "quick": 1, "brown": 1, "fox": 1, "jumps": 1, "over": 1, "lazy": 1, "dog": 1}
func CountWordFrequency(text string) map[string]int {
    rexOne := regexp.MustCompile(`[\.|,|!|\?|']+`)
    rexTwo := regexp.MustCompile(`[\t|\n|-]+`)
    result := map[string]int{}
    count := len(text)
    
    if count <= 0 || count >= 10_000 {
        return result
    }

    text = rexOne.ReplaceAllString(text, "")
    text = rexTwo.ReplaceAllString(text, " ")

    for v := range strings.SplitSeq(text, " ") {
        if v == "" {
            continue
        }
        
        v = strings.ToLower(v)
        if _, ok := result[v]; !ok {
            result[v] = 1
            
            continue
        }
        
        result[v]++
    }
    
	return result
} 