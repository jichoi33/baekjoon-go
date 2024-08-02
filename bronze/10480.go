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

	n := nextInt()

	for i := 0; i < n; i++ {
		x := nextInt()

		if x%2 == 0 {
			writer.WriteString(strconv.Itoa(x) + " is even\n")
		} else {
			writer.WriteString(strconv.Itoa(x) + " is odd\n")
		}
	}
}
