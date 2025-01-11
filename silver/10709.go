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
	// const MaxBuf int = 1_000_001
	// scanner.Buffer(make([]byte, 0, MaxBuf), MaxBuf)
}

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

// ==============================
// Problem Solving Logic
// ==============================

func main() {
	defer writer.Flush()

	H, W := scanInt(), scanInt()
	cloudPresence := make([][]bool, H)
	for h := 0; h < H; h++ {
		cloudPresence[h] = make([]bool, W)
	}

	for h := 0; h < H; h++ {
		s := scanString()
		for w := 0; w < W; w++ {
			if s[w] == 'c' {
				cloudPresence[h][w] = true
			}
		}
	}

	timeToCloud := make([][]int, H)
	for h := 0; h < H; h++ {
		timeToCloud[h] = make([]int, W)
		for w := 0; w < W; w++ {
			timeToCloud[h][w] = -1
		}
	}

	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			if cloudPresence[h][w] {
				timeToCloud[h][w] = 0
			} else if w > 0 && timeToCloud[h][w-1] >= 0 {
				timeToCloud[h][w] = timeToCloud[h][w-1] + 1
			}
		}
	}

	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			writer.WriteString(strconv.Itoa(timeToCloud[h][w]))
			writer.WriteByte(' ')
		}
		writer.WriteByte('\n')
	}
}
