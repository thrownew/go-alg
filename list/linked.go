package list

import (
	"errors"
	"fmt"
	"strings"
)

const (
	Left  Side = true
	Right Side = false
)

type (
	List struct {
		first *Node
		last  *Node
		len   int
	}
	Node struct {
		v    string
		next *Node
		prev *Node
	}
	Side bool
)

var (
	ErrInvalidIndex = errors.New("invalid index")
)

func NewList(ss ...string) *List {
	l := &List{}
	for _, s := range ss {
		l.Push(Right, s)
	}
	return l
}

func (l *List) hasIndex(idx int) bool {
	return idx >= 0 && idx <= l.len-1
}

func (l *List) indexNearestSide(idx int) Side {
	s := Left
	if idx > l.len-idx {
		s = Right
	}
	return s
}

func (l *List) loop(s Side, f func(i int, n *Node) bool) {
	var n *Node
	var idx int
	if s == Left {
		n = l.first
	} else {
		n = l.last
		idx = l.len - 1
	}
	for {
		if n == nil {
			return
		}
		if !f(idx, n) {
			return
		}
		if s == Left {
			idx++
			n = n.next
		} else {
			idx--
			n = n.prev
		}
	}
}

func (l *List) indexOf(s Side, v string) (int, bool) {
	var index int
	var found bool
	l.loop(s, func(i int, n *Node) bool {
		if n.v == v {
			index = i
			found = true
			return false
		}
		return true
	})
	return index, found
}

func (l *List) replace(idx int, v string) (*Node, error) {
	if !l.hasIndex(idx) {
		return nil, ErrInvalidIndex
	}
	var p *Node
	l.loop(l.indexNearestSide(idx), func(i int, n *Node) bool {
		if i != idx {
			return true
		}
		p = n
		node := &Node{
			v:    v,
			prev: n.prev,
			next: n.next,
		}
		if n.next != nil {
			n.next.prev = node
		}
		if n.prev != nil {
			n.prev.next = node
		}
		return false
	})
	return p, nil
}

func (l *List) add(s Side, v string) {
	n := &Node{
		v: v,
	}
	if s == Left {
		if l.first != nil {
			l.first.prev = n
			n.next = l.first
		}
		l.first = n
		if l.last == nil {
			l.last = n
		}
	} else {
		if l.last != nil {
			l.last.next = n
			n.prev = l.last
		}
		l.last = n
		if l.first == nil {
			l.first = n
		}
	}
	l.len++
}

func (l *List) remove(idx int) (*Node, error) {
	if !l.hasIndex(idx) {
		return nil, ErrInvalidIndex
	}
	var p *Node
	l.loop(l.indexNearestSide(idx), func(i int, n *Node) bool {
		if i != idx {
			return true
		}
		l.len--
		p = n
		if i == 0 {
			l.first = n.next
		} else if i == l.len-1 {
			l.last = n.prev
		}
		if n.next != nil {
			n.next.prev = n.prev
		}
		if n.prev != nil {
			n.prev.next = n.next
		}
		return false
	})
	return p, nil
}

func (l *List) String() string {
	b := strings.Builder{}
	b.WriteRune('[')
	l.loop(Left, func(i int, n *Node) bool {
		b.WriteString(" \"")
		b.WriteString(n.v)
		b.WriteRune('"')
		if i != l.len-1 {
			b.WriteRune(',')
		}
		return true
	})
	if l.len > 0 {
		b.WriteRune(' ')
	}
	b.WriteRune(']')
	return b.String()
}

func (l *List) Clear() {
	l.first = nil
	l.last = nil
	l.len = 0
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Loop(s Side, f func(i int, v string) bool) {
	l.loop(s, func(i int, n *Node) bool {
		return f(i, n.v)
	})
}

func (l *List) Contain(v string) bool {
	var found bool
	l.loop(Left, func(_ int, n *Node) bool {
		if n.v != v {
			return true
		}
		found = true
		return false
	})
	return found
}

func (l *List) First() (string, bool) {
	if l.first == nil {
		return "", false
	}
	return l.first.v, true
}

func (l *List) Last() (string, bool) {
	if l.last == nil {
		return "", false
	}
	return l.last.v, true
}

func (l *List) IndexOf(s Side, v string) (int, bool) {
	return l.indexOf(s, v)
}

func (l *List) Push(s Side, vs ...string) {
	for _, v := range vs {
		l.add(s, v)
	}
}

func (l *List) Pop(s Side) (string, bool) {
	if l.len == 0 {
		return "", false
	}
	idx := 0
	if s == Right {
		idx = l.len - 1
	}
	if n, err := l.remove(idx); err == nil {
		return n.v, true
	}
	return "", false
}

func (l *List) Replace(idx int, v string) (string, error) {
	n, err := l.replace(idx, v)
	if err != nil {
		return "", fmt.Errorf("replace: %w", err)
	}
	return n.v, nil
}
