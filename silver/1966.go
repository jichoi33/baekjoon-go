package main

import (
	"bufio"
	"container/heap"
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

// ==============================
// Problem Solving Logic
// ==============================

type Document struct {
	priority int
	index    int
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

	T := scanInt()
	for t := 0; t < T; t++ {
		N, M := scanInt(), scanInt()
		documents := make([]Document, N)
		priorities := &PriorityQueue[int]{
			items: make([]int, 0, 100),
			comparator: func(a, b int) bool {
				return a > b
			},
		}
		for i := 0; i < N; i++ {
			priority := scanInt()
			documents[i] = Document{priority, i}
			heap.Push(priorities, priority)
		}

		count := 1

		for documents[0].index != M || documents[0].priority != priorities.items[0] {
			if documents[0].priority == priorities.items[0] {
				documents = documents[1:]
				heap.Pop(priorities)
				count++
			} else {
				documents = append(documents[1:], documents[0])
			}
		}

		writer.WriteString(strconv.Itoa(count))
		writer.WriteByte('\n')
	}
}
