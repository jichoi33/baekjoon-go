package main

import (
	"bufio"
	"os"
	"sort"
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

type coordinate struct {
	x, y int
}

func main() {
	defer writer.Flush()

	n := scanInt()
	coordinates := make([]coordinate, n)

	for i := 0; i < n; i++ {
		coordinates[i] = coordinate{
			x: scanInt(),
			y: scanInt(),
		}
	}

	sort.Slice(coordinates, func(i, j int) bool {
		if coordinates[i].x == coordinates[j].x {
			return coordinates[i].y < coordinates[j].y
		}
		return coordinates[i].x < coordinates[j].x
	})

	for _, c := range coordinates {
		writer.WriteString(strconv.Itoa(c.x))
		writer.WriteByte(' ')
		writer.WriteString(strconv.Itoa(c.y))
		writer.WriteByte('\n')
	}
}
