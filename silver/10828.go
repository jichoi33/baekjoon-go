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

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(data T) {
	s.items = append(s.items, data)
}

func (s *Stack[T]) Pop() T {
	x := s.items[s.Len()-1]
	s.items = s.items[:s.Len()-1]
	return x
}

func (s *Stack[T]) Peek() T {
	return s.items[s.Len()-1]
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

func main() {
	defer writer.Flush()

	N := scanInt()
	stack := &Stack[int]{
		items: make([]int, 0, N),
	}

	for i := 0; i < N; i++ {
		cmd := scanString()

		switch cmd {
		case "push":
			stack.Push(scanInt())
		case "pop":
			if stack.Len() == 0 {
				writer.WriteString("-1\n")
			} else {
				writer.WriteString(strconv.Itoa(stack.Pop()))
				writer.WriteByte('\n')
			}
		case "size":
			writer.WriteString(strconv.Itoa(stack.Len()))
			writer.WriteByte('\n')
		case "empty":
			if stack.Len() == 0 {
				writer.WriteString("1\n")
			} else {
				writer.WriteString("0\n")
			}
		case "top":
			if stack.Len() == 0 {
				writer.WriteString("-1\n")
			} else {
				writer.WriteString(strconv.Itoa(stack.Peek()))
				writer.WriteByte('\n')
			}
		}
	}
}
