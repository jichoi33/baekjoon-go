package main

import (
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for i := 0; i < t; i++ {
		categoryToCount := make(map[string]int)
		n := nextInt()
		for j := 0; j < n; j++ {
			sc.Scan()
			sc.Scan()
			categoryToCount[sc.Text()] += 1
		}

		res := 1
		for _, c := range categoryToCount {
			res *= (c + 1)
		}
		wr.WriteString(strconv.Itoa(res-1) + "\n")
	}

}
