package main

import (
	"bufio"
	"os"
	"strconv"
)

type coordinate struct {
	z, y, x int
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

	M, N, H := scanInt(), scanInt(), scanInt()

	// Initialize the 3D array of tomatoes
	tomatoes := make([][][]int, H)
	for i := 0; i < H; i++ {
		tomatoes[i] = make([][]int, N)
		for j := 0; j < N; j++ {
			tomatoes[i][j] = make([]int, M)
		}
	}

	// Read input and initialize the queue for BFS
	queue := make([]coordinate, 0, H*N*M)
	unripeCount := 0

	for z := 0; z < H; z++ {
		for y := 0; y < N; y++ {
			for x := 0; x < M; x++ {
				tomatoes[z][y][x] = scanInt()
				if tomatoes[z][y][x] == 1 {
					queue = append(queue, coordinate{z, y, x})
				} else if tomatoes[z][y][x] == 0 {
					unripeCount++
				}
			}
		}
	}

	// If all tomatoes are already ripe, print 0 and return
	if unripeCount == 0 {
		writer.WriteString("0\n")
		return
	}

	dz := []int{0, 0, 0, 0, -1, 1}
	dy := []int{-1, 1, 0, 0, 0, 0}
	dx := []int{0, 0, -1, 1, 0, 0}

	days := 0

	// BFS
	for len(queue) > 0 {
		todaySize := len(queue)
		for i := 0; i < todaySize; i++ {
			cur := queue[0]
			queue = queue[1:]
			for d := 0; d < 6; d++ {
				nz, ny, nx := cur.z+dz[d], cur.y+dy[d], cur.x+dx[d]

				// Check boundaries and if the tomato is unripe
				if nz >= 0 && nz < H && ny >= 0 && ny < N && nx >= 0 && nx < M && tomatoes[nz][ny][nx] == 0 {
					tomatoes[nz][ny][nx] = 1
					queue = append(queue, coordinate{nz, ny, nx})
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
