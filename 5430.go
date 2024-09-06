package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	scanner *bufio.Scanner
	writer  *bufio.Writer
)

const MaxBuf = 400_001

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, MaxBuf), MaxBuf)
	writer = bufio.NewWriter(os.Stdout)
}

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func scanStr() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	defer writer.Flush()

	T := scanInt()

loop:
	for t := 0; t < T; t++ {
		p := scanStr()
		scanner.Scan()
		input := scanStr()
		numbers := strings.Split(input[1:len(input)-1], ",")

		if input == "[]" {
			numbers = []string{}
		}

		reversed := false
		for i := 0; i < len(p); i++ {
			switch p[i] {
			case 'R':
				reversed = !reversed
			case 'D':
				if len(numbers) == 0 {
					writer.WriteString("error\n")
					continue loop
				}
				if !reversed {
					numbers = numbers[1:]
				} else {
					numbers = numbers[:len(numbers)-1]
				}
			}
		}

		writer.WriteByte('[')
		if len(numbers) != 0 {
			if !reversed {
				writer.WriteString(strings.Join(numbers, ","))
			} else {
				for i := len(numbers) - 1; i > 0; i-- {
					writer.WriteString(numbers[i])
					writer.WriteByte(',')
				}
				writer.WriteString(numbers[0])
			}
		}
		writer.WriteString("]\n")
	}
}
