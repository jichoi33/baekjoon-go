package main

import (
	"bufio"
	"os"
	"slices"
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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

// ==============================
// Problem Solving Logic
// ==============================

func main() {
	defer writer.Flush()

	parentColors := make([]string, 4)

	for i := 0; i < 4; i++ {
		parentColors[i] = scanString()
	}

	childColorSet := make(map[string]bool, 16)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			childColor := parentColors[i] + " " + parentColors[j]
			childColorSet[childColor] = true
		}
	}

	childColors := make([]string, 0, len(childColorSet))

	for childColor := range childColorSet {
		childColors = append(childColors, childColor)
	}

	slices.Sort(childColors)

	for i := range childColors {
		writer.WriteString(childColors[i])
		writer.WriteByte('\n')
	}
}
