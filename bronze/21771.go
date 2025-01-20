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

func scanBytes() []byte {
	scanner.Scan()
	bytes := make([]byte, len(scanner.Bytes()))
	copy(bytes, scanner.Bytes())
	return bytes
}

// ==============================
// Problem Solving Logic
// ==============================

func main() {
	defer writer.Flush()

	roomR, roomC := scanInt(), scanInt()
	scanner.Scan()
	scanner.Scan()
	pillowR, pillowC := scanInt(), scanInt()

	room := make([][]byte, roomR)
	for r := 0; r < roomR; r++ {
		room[r] = scanBytes()
	}

	pillowArea := 0

	for r := 0; r < roomR; r++ {
		for c := 0; c < roomC; c++ {
			if room[r][c] == 'P' {
				pillowArea++
			}
		}
	}

	if pillowArea < pillowR*pillowC {
		writer.WriteString("1\n")
	} else {
		writer.WriteString("0\n")
	}
}
