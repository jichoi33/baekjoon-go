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

func scanFloat() float64 {
	scanner.Scan()
	num, _ := strconv.ParseFloat(scanner.Text(), 64)
	return num
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	n := scanInt()

	scores := make([]float64, n)
	maxScore := -1.0

	for i := 0; i < n; i++ {
		scores[i] = scanFloat()
		if scores[i] > maxScore {
			maxScore = scores[i]
		}
	}

	newAvg := 0.0

	for _, score := range scores {
		newAvg += score / maxScore * 100
	}

	newAvg /= float64(n)

	writer.WriteString(strconv.FormatFloat(newAvg, 'f', -1, 64))
	writer.WriteByte('\n')
}
