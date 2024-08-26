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

func isPalindrome(s string) bool {
	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	T := scanInt()
	for t := 0; t < T; t++ {
		n := scanInt()

		reversedN := 0
		tmp := n
		for tmp > 0 {
			reversedN *= 10
			reversedN += tmp % 10
			tmp /= 10
		}

		sum := n + reversedN

		if isPalindrome(strconv.Itoa(sum)) {
			writer.WriteString("YES\n")
		} else {
			writer.WriteString("NO\n")
		}
	}
}
