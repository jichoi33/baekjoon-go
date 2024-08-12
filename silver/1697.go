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

	n, k := scanInt(), scanInt()

	if n >= k {
		writer.WriteString(strconv.Itoa(n - k))
		writer.WriteByte('\n')
		return
	}

	const maxPosition = 100000
	visited := make([]bool, maxPosition+1)
	queue := []int{n}
	visited[n] = true
	steps := 0

	for len(queue) > 0 {
		nextQueue := []int{}
		for _, current := range queue {
			if current == k {
				writer.WriteString(strconv.Itoa(steps))
				writer.WriteByte('\n')
				return
			}
			if current-1 >= 0 && !visited[current-1] {
				visited[current-1] = true
				nextQueue = append(nextQueue, current-1)
			}
			if current+1 <= maxPosition && !visited[current+1] {
				visited[current+1] = true
				nextQueue = append(nextQueue, current+1)
			}
			if current*2 <= maxPosition && !visited[current*2] {
				visited[current*2] = true
				nextQueue = append(nextQueue, current*2)
			}
		}
		queue = nextQueue
		steps++
	}
}
