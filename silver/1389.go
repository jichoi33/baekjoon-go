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

const INF = int(1e9)

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	n, m := scanInt(), scanInt()

	distances := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		distances[i] = make([]int, n+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i != j {
				distances[i][j] = INF
			}
		}
	}

	for i := 0; i < m; i++ {
		a, b := scanInt(), scanInt()
		distances[a][b] = 1
		distances[b][a] = 1
	}

	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if distances[i][k]+distances[k][j] < distances[i][j] {
					distances[i][j] = distances[i][k] + distances[k][j]
				}
			}
		}
	}

	minSum := math.MaxInt64
	minUser := 0

	for i := 1; i <= n; i++ {
		sum := 0
		for j := 1; j <= n; j++ {
			sum += distances[i][j]
		}
		if sum < minSum {
			minSum = sum
			minUser = i
		}
	}

	writer.WriteString(strconv.Itoa(minUser))
	writer.WriteByte('\n')
}
