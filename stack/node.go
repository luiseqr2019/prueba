package stack

type Node[T comparable] struct {
	value T
	next  *Node[T]
}
