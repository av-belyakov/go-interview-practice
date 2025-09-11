package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read input from standard input
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()

		// Call the ReverseString function
		output := ReverseString(input)

		// Print the result
		fmt.Println(output)
	}
}

// ReverseString returns the reversed string of s.
func ReverseString(s string) string {
    num := len(s)
    if num == 0 {
	    return ""
    }

    if num > 1_000 {
        s = s[0:1_000]
    }
    
    strRune := []rune(s)
    newStr := strings.Builder{}
	for i := num-1; i >= 0; i-- {
		newStr.WriteRune(strRune[i])
	}
	
	return newStr.String()
}
