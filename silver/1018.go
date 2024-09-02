package main

import (
	"bufio"
	"math"
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

func scanStr() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	defer writer.Flush()

	n, m := scanInt(), scanInt()
	chessboard := make([][]byte, n)

	for i := 0; i < n; i++ {
		chessboard[i] = []byte(scanStr())
	}

	minCount := math.MaxInt
	for i := 0; i <= n-8; i++ {
		for j := 0; j <= m-8; j++ {
			countWB := 0
			for k := i; k < i+8; k++ {
				for l := j; l < j+8; l++ {
					if (k+l)%2 == 0 {
						if chessboard[k][l] != 'W' {
							countWB++
						}
					} else {
						if chessboard[k][l] != 'B' {
							countWB++
						}
					}
				}
			}

			countBW := 0
			for k := i; k < i+8; k++ {
				for l := j; l < j+8; l++ {
					if (k+l)%2 == 0 {
						if chessboard[k][l] != 'B' {
							countBW++
						}
					} else {
						if chessboard[k][l] != 'W' {
							countBW++
						}
					}
				}
			}

			if countWB < minCount {
				minCount = countWB
			}
			if countBW < minCount {
				minCount = countBW
			}
		}
	}

	writer.WriteString(strconv.Itoa(minCount))
	writer.WriteByte('\n')
}
