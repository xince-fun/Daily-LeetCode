package main

type AllOne struct {
	head, tail *Node
	m          map[string]*Node
}

type Node struct {
	Count      int
	Set        map[string]bool
	Prev, Next *Node
}

func Constructor() AllOne {
	head, tail := &Node{Set: make(map[string]bool)}, &Node{Set: make(map[string]bool)}
	head.Next = tail
	tail.Prev = head
	return AllOne{
		m:    make(map[string]*Node),
		head: head,
		tail: tail,
	}
}

func (t *AllOne) clear(node *Node) {
	if len(node.Set) == 0 {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}
}

func (t *AllOne) Inc(key string) {
	if node, ok := t.m[key]; ok {
		// node.Set[key] = false
		delete(node.Set, key)
		count := node.Count
		next := &Node{}
		if node.Next.Count == count+1 {
			next = node.Next
		} else {
			next = &Node{Count: count + 1, Set: make(map[string]bool)}
			next.Next = node.Next
			next.Prev = node
			node.Next.Prev = next
			node.Next = next
		}
		next.Set[key] = true
		t.m[key] = next
		t.clear(node)
	} else {
		node := &Node{}
		if t.head.Next.Count == 1 {
			node = t.head.Next
		} else {
			node = &Node{Count: 1, Set: make(map[string]bool)}
			node.Next = t.head.Next
			node.Prev = t.head
			t.head.Next.Prev = node
			t.head.Next = node
		}
		node.Set[key] = true
		t.m[key] = node
	}
}

func (t *AllOne) Dec(key string) {
	node := t.m[key]
	delete(node.Set, key)
	// node.Set[key] = false
	count := node.Count
	if count == 1 {
		delete(t.m, key)
	} else {
		prev := &Node{}
		if node.Prev.Count == count-1 {
			prev = node.Prev
		} else {
			prev = &Node{Count: count - 1, Set: make(map[string]bool)}
			prev.Next = node
			prev.Prev = node.Prev
			node.Prev.Next = prev
			node.Prev = prev
		}
		prev.Set[key] = true
		t.m[key] = prev
	}
	t.clear(node)
}

func (t *AllOne) GetMaxKey() string {
	node := t.tail.Prev
	for k := range node.Set {
		return k
	}
	return ""
}

func (t *AllOne) GetMinKey() string {
	node := t.head.Next
	for k := range node.Set {
		return k
	}
	return ""
}
