package main

import (
	"bufio"
	"os"
	"strconv"
)

type MaxHeap struct {
	elements []int
	size     int
}

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func (h *MaxHeap) push(e int) {
	h.elements = append(h.elements, e)
	h.size++
	h.bubbleUp(h.size - 1)
}

func (h *MaxHeap) pop() int {
	root := h.elements[0]
	h.elements[0] = h.elements[h.size-1]
	h.elements = h.elements[:h.size-1]
	h.size--

	if h.size > 1 {
		h.bubbleDown(0)
	}

	return root
}

func (h *MaxHeap) bubbleUp(i int) {
	if i > 0 && h.elements[i] > h.elements[(i-1)/2] {
		h.swap(i, (i-1)/2)
		h.bubbleUp((i - 1) / 2)
	}
}

func (h *MaxHeap) bubbleDown(i int) {
	maxIndex := i
	leftIndex := i*2 + 1
	rightIndex := i*2 + 2

	if leftIndex < h.size && h.elements[leftIndex] > h.elements[maxIndex] {
		maxIndex = leftIndex
	}
	if rightIndex < h.size && h.elements[rightIndex] > h.elements[maxIndex] {
		maxIndex = rightIndex
	}

	if i != maxIndex {
		h.swap(i, maxIndex)
		h.bubbleDown(maxIndex)
	}
}

func (h *MaxHeap) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	h := &MaxHeap{
		elements: make([]int, 0, 100_000),
		size:     0,
	}

	n := nextInt()

	for i := 0; i < n; i++ {
		x := nextInt()
		if x == 0 {
			if h.size == 0 {
				wr.WriteString("0\n")
				continue
			}
			wr.WriteString(strconv.Itoa(h.pop()) + "\n")
		} else {
			h.push(x)
		}
	}
}
