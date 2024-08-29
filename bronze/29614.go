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

	gradeToScore := map[byte]float64{
		'A': 4.0,
		'B': 3.0,
		'C': 2.0,
		'D': 1.0,
		'F': 0.0,
	}

	s := scanString()

	totalScore := 0.0
	subjectCount := 0
	for i := 0; i < len(s)-1; i++ {
		totalScore += gradeToScore[s[i]]
		subjectCount++

		if s[i+1] == '+' {
			totalScore += 0.5
			i++
		}
	}

	if s[len(s)-1] != '+' {
		totalScore += gradeToScore[s[len(s)-1]]
		subjectCount++
	}

	writer.WriteString(strconv.FormatFloat(totalScore/float64(subjectCount), 'f', -1, 64))
	writer.WriteByte('\n')
}
