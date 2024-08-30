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

type Comparator[T any] func(a, b T) bool

type PriorityQueue[T any] struct {
	items      []T
	comparator Comparator[T]
}

func (p *PriorityQueue[T]) Len() int { return len(p.items) }
func (p *PriorityQueue[T]) Less(i, j int) bool {
	return p.comparator(p.items[i], p.items[j])
}
func (p *PriorityQueue[T]) Swap(i, j int) { p.items[i], p.items[j] = p.items[j], p.items[i] }

func (p *PriorityQueue[T]) Push(x any) {
	p.items = append(p.items, x.(T))
}

func (p *PriorityQueue[T]) Pop() any {
	x := p.items[p.Len()-1]
	p.items = p.items[:p.Len()-1]
	return x
}

func main() {
	defer writer.Flush()

	n := scanInt()
	pq := &PriorityQueue[int]{
		items: make([]int, 0, n),
		comparator: func(a, b int) bool {
			return a > b
		},
	}

	for i := 0; i < n; i++ {
		x := scanInt()
		if x == 0 {
			if pq.Len() == 0 {
				writer.WriteString("0\n")
				continue
			}

			v := heap.Pop(pq).(int)
			writer.WriteString(strconv.Itoa(v))
			writer.WriteByte('\n')
			continue
		}

		heap.Push(pq, x)
	}
}
