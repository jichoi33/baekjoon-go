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

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	T := scanInt()

	for t := 0; t < T; t++ {
		a, b := scanString(), scanString()
		count := 0

		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				count++
			}
		}

		writer.WriteString("Hamming distance is ")
		writer.WriteString(strconv.Itoa(count))
		writer.WriteString(".\n")
	}
}
