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

	n, x, y := scanInt(), scanInt(), scanInt()

	mealCount, discard := 0, 0

	for i := 0; i < n; i++ {
		sandwich := scanInt()

		if sandwich >= x {
			mealCount += sandwich / x

			if sandwich > (sandwich/x)*y {
				discard += sandwich - (sandwich/x)*y
			}
		} else {
			discard += sandwich
		}
	}

	writer.WriteString(strconv.Itoa(mealCount))
	writer.WriteByte('\n')
	writer.WriteString(strconv.Itoa(discard))
	writer.WriteByte('\n')
}
