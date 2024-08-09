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

func gcd(x, y int) int {
	if x > y {
		x, y = y, x
	}
	for x%y != 0 {
		r := x % y
		x, y = y, r
	}
	return y
}

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	num1, num2 := scanInt(), scanInt()
	gcdValue := gcd(num1, num2)
	lcmValue := num1 * num2 / gcdValue

	writer.WriteString(strconv.Itoa(gcdValue))
	writer.WriteByte('\n')
	writer.WriteString(strconv.Itoa(lcmValue))
	writer.WriteByte('\n')
}
