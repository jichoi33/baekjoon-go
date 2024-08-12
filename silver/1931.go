package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

type meeting struct {
	startTime int
	endTime   int
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
	meetings := make([]meeting, n)

	for i := 0; i < n; i++ {
		meetings[i] = meeting{
			startTime: scanInt(),
			endTime:   scanInt(),
		}
	}

	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i].endTime == meetings[j].endTime {
			return meetings[i].startTime < meetings[j].startTime
		}
		return meetings[i].endTime < meetings[j].endTime
	})

	currentTime := 0
	count := 0

	for _, meeting := range meetings {
		if meeting.startTime >= currentTime {
			currentTime = meeting.endTime
			count++
		}
	}

	writer.WriteString(strconv.Itoa(count))
	writer.WriteByte('\n')
}
