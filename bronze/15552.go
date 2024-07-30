package main

import (
	"bufio"
	"os"
	"strconv"
)

func nextInt() (r int) {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	sc.Scan()
	t, _ := strconv.Atoi(sc.Text())

	for i := 0; i < t; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		b, _ := strconv.Atoi(sc.Text())

		wr.WriteString(strconv.Itoa(a+b) + "\n")
	}
}
