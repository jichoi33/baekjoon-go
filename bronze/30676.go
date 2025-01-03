package main

import (
	"bufio"
	"os"
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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

// ==============================
// Problem Solving Logic
// ==============================

func main() {
	defer writer.Flush()

	wavelength := scanInt()

	if wavelength >= 620 {
		writer.WriteString("Red\n")
	} else if wavelength >= 590 {
		writer.WriteString("Orange\n")
	} else if wavelength >= 570 {
		writer.WriteString("Yellow\n")
	} else if wavelength >= 495 {
		writer.WriteString("Green\n")
	} else if wavelength >= 450 {
		writer.WriteString("Blue\n")
	} else if wavelength >= 425 {
		writer.WriteString("Indigo\n")
	} else {
		writer.WriteString("Violet\n")
	}
}
