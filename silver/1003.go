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

	zero := make([]int, 41)
	one := make([]int, 41)

	zero[0] = 1
	one[1] = 1
	for i := 2; i < 41; i++ {
		zero[i] = zero[i-1] + zero[i-2]
		one[i] = one[i-1] + one[i-2]
	}

	t := nextInt()
	for i := 0; i < t; i++ {
		n := nextInt()
		wr.WriteString(strconv.Itoa(zero[n]) + " " + strconv.Itoa(one[n]) + "\n")
	}
}
