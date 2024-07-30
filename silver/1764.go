package main

import (
	"bufio"
	"os"
	"slices"
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
	unheard := make(map[string]bool, n)
	var unheardAndUnseen []string

	for i := 0; i < n; i++ {
		sc.Scan()
		unheard[sc.Text()] = true
	}

	for i := 0; i < m; i++ {
		sc.Scan()
		name := sc.Text()
		_, ok := unheard[name]
		if ok {
			unheardAndUnseen = append(unheardAndUnseen, name)
		}
	}

	slices.Sort(unheardAndUnseen)

	wr.WriteString(strconv.Itoa(len(unheardAndUnseen)) + "\n")
	for _, name := range unheardAndUnseen {
		wr.WriteString(name + "\n")
	}
}
