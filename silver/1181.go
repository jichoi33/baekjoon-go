package main

import (
	"bufio"
	"os"
	"sort"
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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	n := scanInt()
	set := make(map[string]bool, n)

	for i := 0; i < n; i++ {
		s := scanString()
		if !set[s] {
			set[s] = true
		}
	}

	words := make([]string, 0, n)

	for s := range set {
		words = append(words, s)
	}

	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			return words[i] < words[j]
		}
		return len(words[i]) < len(words[j])
	})

	for _, word := range words {
		writer.WriteString(word)
		writer.WriteByte('\n')
	}
}
