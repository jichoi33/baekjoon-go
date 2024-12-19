package main

import (
	"bufio"
	"os"
	"strconv"
)

type coordinate struct {
	y, x int
}

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

func main() {
	defer writer.Flush()

	M, N := scanInt(), scanInt()

	// Initialize the 3D array of tomatoes
	tomatoes := make([][]int, N)
	for i := 0; i < N; i++ {
		tomatoes[i] = make([]int, M)
	}

	// Read input and initialize the queue for BFS
	queue := make([]coordinate, 0, N*M)
	unripeCount := 0

	for y := 0; y < N; y++ {
		for x := 0; x < M; x++ {
			tomatoes[y][x] = scanInt()
			if tomatoes[y][x] == 1 {
				queue = append(queue, coordinate{y, x})
			} else if tomatoes[y][x] == 0 {
				unripeCount++
			}
		}
	}

	// If all tomatoes are already ripe, print 0 and return
	if unripeCount == 0 {
		writer.WriteString("0\n")
		return
	}

	dy := []int{-1, 1, 0, 0}
	dx := []int{0, 0, -1, 1}

	days := 0

	// BFS
	for len(queue) > 0 {
		todaySize := len(queue)
		for i := 0; i < todaySize; i++ {
			cur := queue[0]
			queue = queue[1:]
			for d := 0; d < 4; d++ {
				ny, nx := cur.y+dy[d], cur.x+dx[d]

				// Check boundaries and if the tomato is unripe
				if ny >= 0 && ny < N && nx >= 0 && nx < M && tomatoes[ny][nx] == 0 {
					tomatoes[ny][nx] = 1
					queue = append(queue, coordinate{ny, nx})
					unripeCount--

					// If all have ripened, print and return immediately
					if unripeCount == 0 {
						writer.WriteString(strconv.Itoa(days + 1))
						return
					}
				}
			}
		}
		days++
	}

	// There are still unripe tomatoes
	writer.WriteString("-1\n")
}
