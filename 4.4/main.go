package main

import (
	"fmt"
	"strings"
	"strconv"
)

type Node struct {
	key  int
	prev *Node
	next *Node
}

var root = &Node{}

func (node *Node) listSearch(key int) *Node {
	cur := root.next
	for cur != nil && cur.key != key {
		cur = cur.next
	}
	return cur
}

func (node *Node) String() string {
	var rel string
	cur := root.next
	for cur != nil {
		rel = rel + "->" + fmt.Sprintf("%d", cur.key)
		cur = cur.next
	}
	return rel
}

func (node *Node) deleteNode(t *Node) {
	if t == nil {
		return
	}
	t.prev.next = t.next

	//存在后续节点的情况下做处理
	if t.next != nil {
		t.next.prev = t.prev
	}
}

func (node *Node) deleteKey(key int) {
	node.deleteNode(node.listSearch(key))
}

//头插
func (node *Node) insert(key int) {
	new_n := &Node{key: key, prev: root, next: root.next}

	if root.next != nil {
		root.next.prev = new_n
	}
	root.next = new_n

}

func main() {
	orderList := []string{
		"insert 5",
		"insert 2",
		"insert 3",
		"insert 1",
		"delete 3",
		"insert 6",
		"delete 5",
	}

	for _, order := range orderList {
		if strings.HasPrefix(order, "insert") {
			key, _ := strconv.Atoi(strings.SplitAfterN(order, " ", 2)[1])
			root.insert(key)
		}

		if strings.HasPrefix(order, "delete") {
			key, _ := strconv.Atoi(strings.SplitAfterN(order, " ", 2)[1])
			root.deleteKey(key)
		}
	}

	fmt.Println(root)
}
