package main

import (
	"bufio"
	"math"
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

	sum := 0

	for i := 0; i < 5; i++ {
		sum += int(math.Pow(float64(nextInt()), 2))
	}

	writer.WriteString(strconv.Itoa(sum % 10))
	writer.WriteByte('\n')
}
