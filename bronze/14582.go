package main

import (
	"bufio"
	"os"
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

func main() {
	defer writer.Flush()

	woollimScores := make([]int, 9)
	starklinkScores := make([]int, 9)

	for i := 0; i < 9; i++ {
		woollimScores[i] = scanInt()
	}
	for i := 0; i < 9; i++ {
		starklinkScores[i] = scanInt()
	}

	scoreDiff := 0
	wasLeading := false

	for i := 0; i < 9; i++ {
		scoreDiff += woollimScores[i]
		if scoreDiff > 0 {
			wasLeading = true
		}
		scoreDiff -= starklinkScores[i]
	}

	if wasLeading && scoreDiff < 0 {
		writer.WriteString("Yes\n")
	} else {
		writer.WriteString("No\n")
	}
}
