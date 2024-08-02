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

var (
	isPrime = make([]bool, 1001)
)

func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func setIsPrime() {
	for i := 2; i < 1001; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i < 1001; i++ {
		if isPrime[i] {
			for j := i * i; j < 1001; j += i {
				isPrime[j] = false
			}
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := nextInt()

	setIsPrime()
	primeCount := 0

	for i := 0; i < n; i++ {
		if isPrime[nextInt()] {
			primeCount++
		}
	}

	wr.WriteString(strconv.Itoa(primeCount) + "\n")
}
