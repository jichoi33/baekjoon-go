package main

import (
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func nextInt() (r int) {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n, m := nextInt(), nextInt()
	nameToIndex := make(map[string]int, n)
	indexToName := make(map[int]string, n)

	for i := 1; i <= n; i++ {
		sc.Scan()
		name := sc.Text()

		nameToIndex[name] = i
		indexToName[i] = name
	}

	for i := 0; i < m; i++ {
		sc.Scan()
		prob := sc.Text()
		number, ok := nameToIndex[prob]
		if ok {
			wr.WriteString(strconv.Itoa(number) + "\n")
		} else {
			name, _ := strconv.Atoi(prob)
			wr.WriteString(indexToName[name] + "\n")
		}
	}
}
