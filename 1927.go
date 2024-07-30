package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

type minHeap struct {
	elements []int
}

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func (h *minHeap) push(e int) {
	h.elements = append(h.elements, e)
}

func (h *minHeap) pop() int {
	sort.Ints(h.elements)
	tmp := h.elements[0]
	h.elements = h.elements[1:]
	return tmp
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	heap := &minHeap{make([]int, 0)}

	n := nextInt()
	for i := 0; i < n; i++ {
		x := nextInt()
		switch x {
		case 0:
			if len(heap.elements) == 0 {
				wr.WriteString("0\n")
				continue
			}
			wr.WriteString(strconv.Itoa(heap.pop()) + "\n")
		default:
			heap.push(x)
		}
	}
}
