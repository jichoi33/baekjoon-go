package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func sumPlusExpression(s string) int {
	splitPlus := strings.Split(s, "+")
	sum := 0
	for i := 0; i < len(splitPlus); i++ {
		tmp, _ := strconv.Atoi(splitPlus[i])
		sum += tmp
	}
	return sum
}

func main() {
	defer wr.Flush()

	sc.Scan()
	expression := sc.Text()

	splitMinus := strings.Split(expression, "-")

	sum := 0
	sum += sumPlusExpression(splitMinus[0])
	for i := 1; i < len(splitMinus); i++ {
		sum -= sumPlusExpression(splitMinus[i])
	}

	wr.WriteString(strconv.Itoa(sum) + "\n")
}
