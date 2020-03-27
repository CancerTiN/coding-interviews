package main

import "fmt"

type stack struct {
	slice []interface{}
}

func NewStack() (s *stack) {
	s = &stack{slice: make([]interface{}, 0)}
	return
}

func (s *stack) Push(a interface{}) {
	s.slice = append(s.slice, a)
}

func (s *stack) Pop() (a interface{}) {
	a, s.slice = s.slice[len(s.slice)-1], s.slice[:len(s.slice)-1]
	return
}

func (s *stack) Size() (size int) {
	size = len(s.slice)
	return
}

func (s *stack) Top() (a interface{}) {
	a = s.slice[len(s.slice)-1]
	return
}

type queue struct {
	stack1 *stack
	stack2 *stack
}

func (q *queue) appendTail(a interface{}) {
	q.stack1.Push(a)
}

func (q *queue) deleteHead() (a interface{}) {
	if q.stack2.Size() != 0 {
		a = q.stack2.Pop()
	} else {
		if q.stack1.Size() != 0 {
			for q.stack1.Size() > 0 {
				q.stack2.Push(q.stack1.Pop())
			}
			a = q.stack2.Pop()
		} else {
			panic("queue is empty")
		}
	}
	return
}

func NewQueue() (q *queue) {
	q = &queue{stack1: NewStack(), stack2: NewStack()}
	return
}

func main() {
	q := NewQueue()
	fmt.Printf("queue: %#v\n", q)
	for _, c := range "abc" {
		q.appendTail(c)
	}
	fmt.Printf("queue: %#v\n", q)
	fmt.Printf("queue pop: %#v\n", q.deleteHead())
	fmt.Printf("queue: %#v\n", q)
	fmt.Printf("queue pop: %#v\n", q.deleteHead())
	fmt.Printf("queue: %#v\n", q)
	q.appendTail('d')
	fmt.Printf("queue: %#v\n", q)
	fmt.Printf("queue pop: %#v\n", q.deleteHead())
	fmt.Printf("queue: %#v\n", q)
}
