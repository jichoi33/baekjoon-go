package main

import (
	"bufio"
	"os"
	"strconv"
)

// ===========================
// I/O Setup
// ===========================
var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func init() {
	scanner.Split(bufio.ScanWords)
}

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

// ==============================
// Problem Solving Logic
// ==============================

func main() {
	defer writer.Flush()

	numTestCases := scanInt()
	for t := 0; t < numTestCases; t++ {
		targetWord := scanString()

		numWords := scanInt()
		words := make([]string, numWords)

		for i := 0; i < numWords; i++ {
			words[i] = scanString()
		}

		bestMatchingWord := ""
		maxMatchingCount := -1
		for _, word := range words {
			mathingCount := 0
			for i := 0; i < len(targetWord); i++ {
				if targetWord[i] == word[i] {
					mathingCount++
				}
			}
			if mathingCount > maxMatchingCount {
				bestMatchingWord = word
				maxMatchingCount = mathingCount
			}
		}

		writer.WriteString(bestMatchingWord)
		writer.WriteByte('\n')
	}
}
