package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

type member struct {
	order int
	age   int
	name  string
}

var (
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func scanStr() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	scanner.Split(bufio.ScanWords)
	defer writer.Flush()

	n := scanInt()

	members := make([]member, n)

	for i := 0; i < n; i++ {
		members[i] = member{
			order: i,
			age:   scanInt(),
			name:  scanStr(),
		}
	}

	sort.Slice(members, func(i, j int) bool {
		if members[i].age == members[j].age {
			return members[i].order < members[j].order
		}
		return members[i].age < members[j].age
	})

	for _, m := range members {
		writer.WriteString(strconv.Itoa(m.age))
		writer.WriteByte(' ')
		writer.WriteString(m.name)
		writer.WriteByte('\n')
	}
}
