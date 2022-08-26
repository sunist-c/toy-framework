package channel

import (
	"sync"
)

type Channel[Type any] struct {
	exMu  *sync.Mutex
	reMu  *sync.Mutex
	first *node[Type]
	last  *node[Type]
	buff  chan Type
	idle  chan *node[Type]
}

func (c *Channel[Type]) extend() {
	if len(c.first.ch) != cap(c.first.ch) {
		return
	}

	select {
	case n := <-c.idle:
		c.last.next = n
		c.last = n
	default:
		c.last = newNode[Type](c.last)
	}
}

func (c *Channel[Type]) reduce() {
	if len(c.first.ch) != 0 {
		return
	}

	select {
	case c.idle <- c.first:
	default:
	}
	t := c.first
	c.first = c.first.next
	t.next = nil
}

func (c *Channel[Type]) Push(t Type) {
	if !c.last.push(t) {
		c.exMu.Lock()
		c.extend()
		c.exMu.Unlock()
		c.Push(t)
	}
}

func (c *Channel[Type]) Pop() (t Type, success bool) {
	if t, success = c.first.pop(); success {
		return t, success
	}

	if c.first == c.last {
		return t, false
	}

	c.reMu.Lock()
	c.reduce()
	c.reMu.Unlock()
	return c.Pop()

}

func NewChannel[Type any](maxIdleNodes ...uint) *Channel[Type] {
	ch := &Channel[Type]{
		exMu:  &sync.Mutex{},
		reMu:  &sync.Mutex{},
		first: newNode[Type](nil),
	}
	ch.last = ch.first
	if len(maxIdleNodes) == 0 {
		ch.idle = make(chan *node[Type], 16)
	} else {
		ch.idle = make(chan *node[Type], maxIdleNodes[0])
	}

	return ch
}
