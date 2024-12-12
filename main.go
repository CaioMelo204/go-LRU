package main

import "fmt"

const SIZE = 10

type Node struct {
	left  *Node
	value string
	right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

func newCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := new(Node)
	tail := new(Node)
	head.right = tail
	tail.left = head
	return Queue{Head: head, Tail: tail}
}

func (c *Cache) Put(value string) {
	node := &Node{value: value}

	if val, ok := c.Hash[value]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{value: value}
	}
	c.Hash[value] = node
	c.Unshift(node)
}

func (c *Cache) Unshift(n *Node) {
	fmt.Println("add", n.value)

	tmp := c.Queue.Head.right

	c.Queue.Head.right = n
	n.left = c.Queue.Head
	n.right = tmp
	tmp.left = n

	c.Queue.Length++
	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.left)
	}
}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Println("Remove ", n.value)

	left := n.left
	right := n.right

	left.right = right
	right.left = left
	c.Queue.Length -= 1
	delete(c.Hash, n.value)
	return n
}

func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	node := q.Head.right
	fmt.Printf("%d - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.value)
		if i < q.Length-1 {
			fmt.Printf("<-->")
		}
		node = node.right
	}
	fmt.Println("]")
}

type Hash map[string]*Node

func main() {
	fmt.Println("Start Cache")
	cache := newCache()
	for _, word := range []string{"a", "b", "c", "d", "e", "f", "g", "b", "g"} {
		cache.Put(word)
		cache.Display()
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
