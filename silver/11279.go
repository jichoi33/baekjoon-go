package main

import (
	"bufio"
	"container/heap"
	"os"
	"strconv"
)

var (
	scanner *bufio.Scanner
	writer  *bufio.Writer
)

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
}

func scanInt() int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

type PriorityQueue []int

func (p *PriorityQueue) Len() int { return len(*p) }
func (p *PriorityQueue) Less(i int, j int) bool {
	return (*p)[i] > (*p)[j]
}

func (p *PriorityQueue) Swap(i int, j int) { (*p)[i], (*p)[j] = (*p)[j], (*p)[i] }

func (p *PriorityQueue) Push(x any) {
	*p = append(*p, x.(int))
}

func (p *PriorityQueue) Pop() any {
	x := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return x
}

func main() {
	defer writer.Flush()

	n := scanInt()
	pq := make(PriorityQueue, 0, n)

	for i := 0; i < n; i++ {
		x := scanInt()
		if x == 0 {
			if len(pq) == 0 {
				writer.WriteString("0\n")
				continue
			}

			v := heap.Pop(&pq).(int)
			writer.WriteString(strconv.Itoa(v))
			writer.WriteByte('\n')
			continue
		}

		heap.Push(&pq, x)
	}
}
