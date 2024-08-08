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

	n, r, c := scanInt(), scanInt(), scanInt()

	ans := 0

	for n > 1 {
		half := 1 << (n - 1)
		quadrant := 1 << ((n - 1) * 2)

		if r >= half && c >= half {
			ans += quadrant * 3
			r -= half
			c -= half
		} else if r >= half && c < half {
			ans += quadrant * 2
			r -= half
		} else if r < half && c >= half {
			ans += quadrant * 1
			c -= half
		}
		n -= 1
	}

	if r == 0 && c == 1 {
		ans += 1
	} else if r == 1 && c == 0 {
		ans += 2
	} else if r == 1 && c == 1 {
		ans += 3
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}
