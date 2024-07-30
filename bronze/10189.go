package main

import (
	"bufio"
	"os"
)

var wr = bufio.NewWriter(os.Stdout)

func main() {
	defer wr.Flush()

	wr.WriteString("#  # #### #### #  #\n#### #  # #  # # #\n#### #  # #  # # #\n#  # #### #### #  #\n")
}
