package main

import (
	"bufio"
	"os"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	defer writer.Flush()

	for {
		scanner.Scan()
		input := scanner.Text()

		if input == "0" {
			return
		}

		if isPalindrome(input) {
			writer.WriteString("yes\n")
		} else {
			writer.WriteString("no\n")
		}
	}
}
