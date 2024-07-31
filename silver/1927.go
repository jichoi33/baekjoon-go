package main

import (
	"bufio"
	"os"
	"strconv"
)

type MinHeap struct {
	elements []int
	size     int
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

func (h *MinHeap) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func (h *MinHeap) bubbleUp(i int) {
	if i > 0 && h.elements[i] < h.elements[(i-1)/2] {
		h.swap(i, (i-1)/2)
		h.bubbleUp((i - 1) / 2)
	}
}

func (h *MinHeap) bubbleDown(i int) {
	smallest := i
	left := i*2 + 1
	right := i*2 + 2

	if left < h.size && h.elements[left] < h.elements[smallest] {
		smallest = left
	}
	if right < h.size && h.elements[right] < h.elements[smallest] {
		smallest = right
	}

	if smallest != i {
		h.swap(i, smallest)
		h.bubbleDown(smallest)
	}
}

func (h *MinHeap) Push(e int) {
	h.elements = append(h.elements, e)
	h.size++
	h.bubbleUp(h.size - 1)
}

func (h *MinHeap) Pop() int {
	root := h.elements[0]
	h.elements[0] = h.elements[h.size-1]
	h.elements = h.elements[:h.size-1]
	h.size--
	h.bubbleDown(0)

	return root
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	h := &MinHeap{make([]int, 0, 100_000), 0}

	n := nextInt()
	for i := 0; i < n; i++ {
		x := nextInt()
		switch x {
		case 0:
			if h.size == 0 {
				wr.WriteString("0\n")
				continue
			}
			wr.WriteString(strconv.Itoa(h.Pop()) + "\n")
		default:
			h.Push(x)
		}
	}
}
