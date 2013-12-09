package datastructure

import "fmt"

type List struct {
	Root *Node
}

type Node struct {
	Value int
	Next  *Node
}

func (l *List) add(val int) {
	l.addNode(&Node{val, nil})
}

func (l *List) addNode(n *Node) {
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

func (l *List) insert(val int, pos int) {
	tmp := l.Root
	if tmp == nil || pos == 0 {
		l.Root = &Node{val, tmp}
		return
	}

	for i := 0; i < pos-1 && tmp.Next != nil; i++ {
		tmp = tmp.Next
	}
	newNode := &Node{val, tmp.Next}
	tmp.Next = newNode
}

func (l *List) delete(pos int) {
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

func (l *List) print() {
	tmp := l.Root
	for tmp != nil {
		fmt.Println(tmp.Value)
		tmp = tmp.Next
	}
}

func Test() {
	list := &List{nil}
	list.add(2)
	list.add(5)
	list.insert(1, 0)
	list.insert(6, 100)
	list.delete(1)
	list.print()
}
