package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	for sc.Scan() {
		input := sc.Text()
		if input == "0 0 0" {
			break
		}
		words := strings.Fields(input)

		numbers := make([]int, 3)
		for i := range words {
			numbers[i], _ = strconv.Atoi(words[i])
		}

		sort.Slice(numbers, func(i, j int) bool {
			return numbers[i] < numbers[j]
		})

		if numbers[0]*numbers[0]+numbers[1]*numbers[1] == numbers[2]*numbers[2] {
			fmt.Fprintln(wr, "right")
		} else {
			fmt.Fprintln(wr, "wrong")
		}
	}
}
