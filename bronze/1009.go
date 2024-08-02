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

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	T := nextInt()

	for t := 0; t < T; t++ {
		a, b := nextInt(), nextInt()
		var ans int

		c := b % 4
		if c == 0 {
			c = 4
		}
		ans = powInt(a%10, c) % 10
		if ans == 0 {
			ans = 10
		}

		writer.WriteString(strconv.Itoa(ans) + "\n")
	}
}
