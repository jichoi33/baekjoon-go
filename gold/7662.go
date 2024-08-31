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
	scanner.Split(bufio.ScanWords)
	writer = bufio.NewWriter(os.Stdout)
}

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func scanStr() string {
	scanner.Scan()
	return scanner.Text()
}

type Comparator[T any] func(a, b T) bool

type PriorityQueue[T any] struct {
	items      []T
	comparator Comparator[T]
}

func (p *PriorityQueue[T]) Len() int {
	return len(p.items)
}

func (p *PriorityQueue[T]) Less(i int, j int) bool {
	return p.comparator(p.items[i], p.items[j])
}

func (p *PriorityQueue[T]) Pop() any {
	x := p.items[p.Len()-1]
	p.items = p.items[:p.Len()-1]
	return x
}

func (p *PriorityQueue[T]) Push(x any) {
	p.items = append(p.items, x.(T))
}

func (p *PriorityQueue[T]) Swap(i int, j int) {
	p.items[i], p.items[j] = p.items[j], p.items[i]
}

func main() {
	defer writer.Flush()

	T := scanInt()

	for t := 0; t < T; t++ {
		k := scanInt()

		maxPQ := &PriorityQueue[int]{
			items: make([]int, 0, k),
			comparator: func(a, b int) bool {
				return a > b
			},
		}
		minPQ := &PriorityQueue[int]{
			items: make([]int, 0, k),
			comparator: func(a, b int) bool {
				return a < b
			},
		}
		hashMap := make(map[int]int, k)

		for i := 0; i < k; i++ {
			cmd := scanStr()
			n := scanInt()

			if cmd == "I" {
				hashMap[n]++
				heap.Push(maxPQ, n)
				heap.Push(minPQ, n)
			} else {
				if len(hashMap) > 0 {
					if n == 1 {
						for {
							x := heap.Pop(maxPQ).(int)
							if hashMap[x] >= 1 {
								hashMap[x]--
								if hashMap[x] == 0 {
									delete(hashMap, x)
								}
								break
							}
						}
					} else {
						for {
							x := heap.Pop(minPQ).(int)
							if hashMap[x] >= 1 {
								hashMap[x]--
								if hashMap[x] == 0 {
									delete(hashMap, x)
								}
								break
							}
						}
					}
				}
			}
		}

		if len(hashMap) == 0 {
			writer.WriteString("EMPTY\n")
		} else {
			for {
				x := heap.Pop(maxPQ).(int)
				if hashMap[x] > 0 {
					writer.WriteString(strconv.Itoa(x))
					writer.WriteByte(' ')
					break
				}
			}
			for {
				x := heap.Pop(minPQ).(int)
				if hashMap[x] > 0 {
					writer.WriteString(strconv.Itoa(x))
					writer.WriteByte('\n')
					break
				}
			}
		}
	}
}
