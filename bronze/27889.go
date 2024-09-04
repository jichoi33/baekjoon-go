package main

import (
	"bufio"
	"os"
)

var (
	scanner *bufio.Scanner
	writer  *bufio.Writer
)

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer = bufio.NewWriter(os.Stdout)
}

func scanStr() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	defer writer.Flush()

	abbrToName := map[string]string{
		"NLCS": "North London Collegiate School",
		"BHA":  "Branksome Hall Asia",
		"KIS":  "Korea International School",
		"SJA":  "St. Johnsbury Academy",
	}

	abbr := scanStr()

	writer.WriteString(abbrToName[abbr])
	writer.WriteByte('\n')
}
