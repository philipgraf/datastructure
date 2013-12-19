package list

type DLink struct {
	Root *node
}

type node struct {
	Next  *node
	Prev  *node
	Value int
}
