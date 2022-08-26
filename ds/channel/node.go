package channel

type node[Type any] struct {
	ch   chan Type
	next *node[Type]
}

func (n *node[Type]) push(t Type) (success bool) {
	select {
	case n.ch <- t:
		return true
	default:
		return false
	}
}

func (n *node[Type]) pop() (t Type, success bool) {
	select {
	case t = <-n.ch:
		return t, true
	default:
		return t, false
	}
}

func newNode[Type any](prev *node[Type]) *node[Type] {
	n := &node[Type]{
		ch:   make(chan Type, 128),
		next: nil,
	}

	if prev != nil {
		prev.next = n
	}

	return n
}
