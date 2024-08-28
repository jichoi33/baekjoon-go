package main

import (
	"bufio"
	"os"
	"strconv"
)

type person struct {
	weight int
	height int
}

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

	n := scanInt()
	people := make([]person, n)

	for i := 0; i < n; i++ {
		people[i] = person{
			weight: scanInt(),
			height: scanInt(),
		}
	}

	for i := 0; i < n; i++ {
		rank := 1
		for j := 0; j < n; j++ {
			if people[j].weight > people[i].weight && people[j].height > people[i].height {
				rank++
			}
		}

		writer.WriteString(strconv.Itoa(rank))
		writer.WriteByte(' ')
	}
	writer.WriteByte('\n')
}
