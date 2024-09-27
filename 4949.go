package main

import (
	"bufio"
	"os"
)

var (
	scanner *bufio.Scanner
	writer  *bufio.Writer
)

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
}

func scanStr() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	defer writer.Flush()

	for {
		s := scanStr()
		if s == "." {
			break
		}

		stack := make([]rune, 0, len(s))
		valid := true

		for _, c := range s {
			switch c {
			case '[', '(':
				stack = append(stack, c)
			case ']':
				if len(stack) == 0 || stack[len(stack)-1] != '[' {
					valid = false
					break
				}
				stack = stack[:len(stack)-1]
			case ')':
				if len(stack) == 0 || stack[len(stack)-1] != '(' {
					valid = false
					break
				}
				stack = stack[:len(stack)-1]
			}
		}

		if valid && len(stack) == 0 {
			writer.WriteString("yes\n")
		} else {
			writer.WriteString("no\n")
		}
	}
}
