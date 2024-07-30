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

	set := make(map[int]bool, 20)

	m := nextInt()
	for i := 0; i < m; i++ {
		sc.Scan()
		cmd := sc.Text()
		switch cmd {
		case "add":
			x := nextInt()
			set[x] = true
		case "remove":
			x := nextInt()
			set[x] = false
		case "check":
			x := nextInt()
			if set[x] {
				wr.WriteString("1\n")
			} else {
				wr.WriteString("0\n")
			}
		case "toggle":
			x := nextInt()
			set[x] = !set[x]
		case "all":
			for i := 1; i < 21; i++ {
				set[i] = true
			}
		case "empty":
			for i := 1; i < 21; i++ {
				set[i] = false
			}
		}
	}
}
