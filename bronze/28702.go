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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	for i := 3; i >= 1; i-- {
		s := scanString()
		if s[0] < 'A' {
			num, _ := strconv.Atoi(s)
			num += i

			if num%3 == 0 && num%5 == 0 {
				writer.WriteString("FizzBuzz")
			} else if num%3 == 0 {
				writer.WriteString("Fizz")
			} else if num%5 == 0 {
				writer.WriteString("Buzz")
			} else {
				writer.WriteString(strconv.Itoa(num))
			}
			writer.WriteByte('\n')
			return
		}
	}
}
