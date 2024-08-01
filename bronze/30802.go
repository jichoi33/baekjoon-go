package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := nextInt()

	tShirts := make([]int, 6)
	for i := 0; i < 6; i++ {
		tShirts[i] = nextInt()
	}

	t, p := nextInt(), nextInt()

	var ansT int
	for i := 0; i < 6; i++ {
		ansT += tShirts[i] / t
		if tShirts[i]%t != 0 {
			ansT++
		}
	}

	wr.WriteString(strconv.Itoa(ansT) + "\n")
	wr.WriteString(strconv.Itoa(n/p) + " " + strconv.Itoa(n%p) + "\n")
}
