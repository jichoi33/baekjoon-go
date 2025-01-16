package main

import (
	"bufio"
	"os"
	"strconv"
)

// ===========================
// I/O Setup
// ===========================
var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func init() {
	scanner.Split(bufio.ScanWords)
}

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

type item struct {
	pos, count int
}

func main() {
	defer writer.Flush()

	move := make([]int, 101)
	for i := 1; i <= 100; i++ {
		move[i] = i
	}

	N, M := scanInt(), scanInt()

	for i := 0; i < N; i++ {
		x, y := scanInt(), scanInt()
		move[x] = y
	}
	for i := 0; i < M; i++ {
		u, v := scanInt(), scanInt()
		move[u] = v
	}

	queue := []item{{1, 0}}
	visited := make([]bool, 101)
	visited[1] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.pos == 100 {
			writer.WriteString(strconv.Itoa(cur.count))
			writer.WriteByte('\n')
			return
		}

		for dice := 1; dice <= 6; dice++ {
			nextPos := cur.pos + dice
			if nextPos > 100 {
				break
			}

			finalPos := move[nextPos]

			if !visited[finalPos] {
				visited[finalPos] = true
				queue = append(queue, item{finalPos, cur.count + 1})
			}
		}
	}
}
