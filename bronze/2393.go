package main

import (
	"bufio"
	"os"
)

var wr = bufio.NewWriter(os.Stdout)

func main() {
	defer wr.Flush()

	wr.WriteString("  ___  ___  ___\n  | |__| |__| |\n  |           |\n   \\_________/\n    \\_______/\n     |     |\n     |     |\n     |     |\n     |     |\n     |_____|\n  __/       \\__\n /             \\\n/_______________\\\n")
}
