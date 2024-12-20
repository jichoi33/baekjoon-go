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
	scanner.Split(bufio.ScanWords)
	writer = bufio.NewWriter(os.Stdout)
}

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	defer writer.Flush()

	myMap := make(map[int]bool, 42)

	for i := 0; i < 10; i++ {
		myMap[scanInt()%42] = true
	}

	writer.WriteString(strconv.Itoa(len(myMap)))
	writer.WriteByte('\n')
}
