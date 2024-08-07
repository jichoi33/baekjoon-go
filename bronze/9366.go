package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func nextInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	T := nextInt()

	for t := 1; t <= T; t++ {
		a, b, c := nextInt(), nextInt(), nextInt()

		if a > b {
			a, b = b, a
		}
		if b > c {
			b, c = c, b
		}

		writer.WriteString("Case #")
		writer.WriteString(strconv.Itoa(t))

		if a+b <= c {
			writer.WriteString(": invalid!\n")
			continue
		}

		if a == b && b == c {
			writer.WriteString(": equilateral\n")
		} else if a == b || b == c || c == a {
			writer.WriteString(": isosceles\n")
		} else {
			writer.WriteString(": scalene\n")
		}
	}
}
