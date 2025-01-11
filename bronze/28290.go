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

	typeMap := make(map[string]string, 6)

	typeMap["fdsajkl;"] = "in-out"
	typeMap["jkl;fdsa"] = "in-out"
	typeMap["asdf;lkj"] = "out-in"
	typeMap[";lkjasdf"] = "out-in"
	typeMap["asdfjkl;"] = "stairs"
	typeMap[";lkjfdsa"] = "reverse"

	input := scanString()

	output, ok := typeMap[input]

	if ok {
		writer.WriteString(output)
	} else {
		writer.WriteString("molu")
	}
	writer.WriteByte('\n')
}
