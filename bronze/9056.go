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

	T := scanInt()

loop:
	for t := 0; t < T; t++ {
		s1 := scanBytes()
		s2 := scanBytes()

		alp1 := make([]bool, 26)
		for _, c := range s1 {
			alp1[c-'A'] = true
		}

		alp2 := make([]bool, 26)
		for _, c := range s2 {
			alp2[c-'A'] = true
		}

		for i := 0; i < 26; i++ {
			if alp1[i] != alp2[i] {
				writer.WriteString("NO\n")
				continue loop
			}
		}
		writer.WriteString("YES\n")
	}
}
