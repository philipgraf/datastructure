package list

import "fmt"

type List struct {
	Root *node
}

type node struct {
	Value int
	Next  *node
}

func (l *List) Add(val int) {
	l.addNode(&node{val, nil})
}

func (l *List) addNode(n *node) {
	if l.Root == nil {
		l.Root = n
		return
	}
	tmp := l.Root
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = n
}

func (l *List) Insert(val int, pos int) {
	tmp := l.Root
	if tmp == nil || pos == 0 {
		l.Root = &node{val, tmp}
		return
	}

	for i := 0; i < pos-1 && tmp.Next != nil; i++ {
		tmp = tmp.Next
	}
	newNode := &node{val, tmp.Next}
	tmp.Next = newNode
}

func (l *List) Delete(pos int) {
	tmp := l.Root
	if pos == 0 && tmp != nil {
		l.Root = tmp.Next
	}
	for i := 0; i < pos-1 && tmp.Next != nil; i++ {
		tmp = tmp.Next
	}
	if tmp.Next != nil {
		tmp.Next = tmp.Next.Next
	}
}

func (l *List) Print() {
  fmt.Println("Printing List")
	tmp := l.Root
	for tmp != nil {
		fmt.Println(tmp.Value)
		tmp = tmp.Next
	}
}
