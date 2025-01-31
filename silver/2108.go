package main

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
)

// ===========================
// I/O Setup
// ===========================
var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func init() {
	scanner.Split(bufio.ScanWords)
}

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

// ==============================
// Problem Solving Logic
// ==============================

func main() {
	defer writer.Flush()

	N := scanInt()
	numbers := make([]int, N)

	for i := range numbers {
		numbers[i] = scanInt()
	}

	sum := 0
	for _, num := range numbers {
		sum += num
	}
	mean := int(math.Round(float64(sum) / float64(N)))

	sort.Ints(numbers)
	median := numbers[N/2]

	var mode int
	frequency := make(map[int]int)
	for _, num := range numbers {
		frequency[num]++
	}

	maxFreq := 0
	for _, freq := range frequency {
		if freq > maxFreq {
			maxFreq = freq
		}
	}

	var modes []int
	for num, freq := range frequency {
		if freq == maxFreq {
			modes = append(modes, num)
		}
	}

	sort.Ints(modes)

	if len(modes) >= 2 {
		mode = modes[1]
	} else {
		mode = modes[0]
	}

	minMaxDiff := numbers[N-1] - numbers[0]

	writer.WriteString(strconv.Itoa(mean))
	writer.WriteByte('\n')
	writer.WriteString(strconv.Itoa(median))
	writer.WriteByte('\n')
	writer.WriteString(strconv.Itoa(mode))
	writer.WriteByte('\n')
	writer.WriteString(strconv.Itoa(minMaxDiff))
	writer.WriteByte('\n')
}
