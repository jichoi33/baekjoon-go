package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	n := scanInt()

	num := 666
	count := 0

	for {
		if strings.Contains(strconv.Itoa(num), "666") {
			count++
			if count == n {
				writer.WriteString(strconv.Itoa(num))
				writer.WriteByte('\n')
				return
			}
		}
		num++
	}
}
