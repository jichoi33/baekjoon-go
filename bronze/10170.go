package main

import (
	"bufio"
	"os"
)

var (
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()

	writer.WriteString(`NFC West       W   L  T
-----------------------
Seattle        13  3  0
San Francisco  12  4  0
Arizona        10  6  0
St. Louis      7   9  0

NFC North      W   L  T
-----------------------
Green Bay      8   7  1
Chicago        8   8  0
Detroit        7   9  0
Minnesota      5  10  1`)
	writer.WriteByte('\n')
}
