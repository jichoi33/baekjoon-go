package main

import (
	"bufio"
	"os"
	"strconv"
)

// ===========================
// I/O Setup
// ===========================

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

func scanString() string {
	scanner.Scan()
	return scanner.Text()
}

// ==============================
// Problem Solving Logic
// ==============================

type Queue[T any] struct {
	items []T
}

func (s *Queue[T]) Push(data T) {
	s.items = append(s.items, data)
}

func (s *Queue[T]) Pop() T {
	x := s.items[0]
	s.items = s.items[1:]
	return x
}

func (s *Queue[T]) PeekFront() T {
	return s.items[0]
}

func (s *Queue[T]) PeekBack() T {
	return s.items[s.Len()-1]
}

func (s *Queue[T]) Len() int {
	return len(s.items)
}

func main() {
	defer writer.Flush()

	N := scanInt()
	queue := &Queue[int]{
		items: make([]int, 0, N),
	}

	for i := 0; i < N; i++ {
		cmd := scanString()

		switch cmd {
		case "push":
			queue.Push(scanInt())
		case "pop":
			if queue.Len() == 0 {
				writer.WriteString("-1\n")
			} else {
				writer.WriteString(strconv.Itoa(queue.Pop()))
				writer.WriteByte('\n')
			}
		case "size":
			writer.WriteString(strconv.Itoa(queue.Len()))
			writer.WriteByte('\n')
		case "empty":
			if queue.Len() == 0 {
				writer.WriteString("1\n")
			} else {
				writer.WriteString("0\n")
			}
		case "front":
			if queue.Len() == 0 {
				writer.WriteString("-1\n")
			} else {
				writer.WriteString(strconv.Itoa(queue.PeekFront()))
				writer.WriteByte('\n')
			}
		case "back":
			if queue.Len() == 0 {
				writer.WriteString("-1\n")
			} else {
				writer.WriteString(strconv.Itoa(queue.PeekBack()))
				writer.WriteByte('\n')
			}
		}
	}
}
