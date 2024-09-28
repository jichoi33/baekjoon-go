package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	scanner *bufio.Scanner
	writer  *bufio.Writer
)

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
}

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func scanStr() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	defer writer.Flush()

	T := scanInt()
	for t := 0; t < T; t++ {
		s := scanStr()
		stack := make([]rune, 0, len(s))
		valid := true

		for _, r := range s {
			switch r {
			case '(':
				stack = append(stack, r)
			case ')':
				if len(stack) == 0 {
					valid = false
					break
				}

				stack = stack[:len(stack)-1]
			}
		}

		if valid && len(stack) == 0 {
			writer.WriteString("YES\n")
		} else {
			writer.WriteString("NO\n")
		}
	}
}
