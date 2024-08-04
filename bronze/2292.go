package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func nextInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	n := nextInt()

	room := 1
	for i := 0; ; i++ {
		room += i * 6
		if n <= room {
			writer.WriteString(strconv.Itoa(i+1) + "\n")
			return
		}
	}
}
