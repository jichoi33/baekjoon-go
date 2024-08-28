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

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	T := scanInt()

	for t := 0; t < T; t++ {
		m, n, x, y := scanInt(), scanInt(), scanInt(), scanInt()

		curX, curY := 1, 1
		count := 1

		for {
			if curX > m || curY > n {
				writer.WriteString("-1\n")
				break
			}
			if curX-curY != x-y {
				if m-curX < n-curY {
					count += m - curX + 1
					curY += m - curX + 1
					curX = 1
				} else {
					count += n - curY + 1
					curX += n - curY + 1
					curY = 1
				}
			} else {
				writer.WriteString(strconv.Itoa(count + x - curX))
				writer.WriteByte('\n')
				break
			}
		}
	}
}
