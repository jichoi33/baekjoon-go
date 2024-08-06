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

const (
	Mod = 1_234_567_891
	R   = 31
)

func nextInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	l := nextInt()

	scanner.Scan()
	input := scanner.Text()

	hashValue := 0

	for i := 0; i < l; i++ {
		tmp := int(input[i] - 'a' + 1)
		for j := 0; j < i; j++ {
			tmp *= R
			tmp %= Mod
		}
		hashValue += tmp
	}

	hashValue %= Mod

	writer.WriteString(strconv.Itoa(hashValue))
	writer.WriteByte('\n')
}
