package main

import (
	"bufio"
	"os"
	"strconv"
)

type cow struct {
	startTime int
	endTime   int
	buckets   int
}

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	n := scanInt()

	cows := make([]cow, n)

	for i := 0; i < n; i++ {
		cows[i] = cow{
			startTime: scanInt(),
			endTime:   scanInt(),
			buckets:   scanInt(),
		}
	}

	maxBuckets := 0

	for curTime := 1; curTime <= 1000; curTime++ {
		curBuckets := 0
		for _, cow := range cows {
			if cow.startTime <= curTime && curTime <= cow.endTime {
				curBuckets += cow.buckets
			}
		}
		if curBuckets > maxBuckets {
			maxBuckets = curBuckets
		}
	}

	writer.WriteString(strconv.Itoa(maxBuckets))
	writer.WriteByte('\n')
}
