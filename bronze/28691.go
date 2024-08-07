package main

import (
	"bufio"
	"os"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	firstLetterToClubName := map[string]string{
		"M": "MatKor",
		"W": "WiCys",
		"C": "CyKor",
		"A": "AlKor",
		"$": "$clear",
	}

	scanner.Scan()
	firstLetter := scanner.Text()

	writer.WriteString(firstLetterToClubName[firstLetter])
	writer.WriteByte('\n')
}
