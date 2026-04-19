package linkedlist

/*
	Очередь на двунаправленном LinkedList
	интерфейс: put, take
*/

// type Node struct {
// 	value string
// 	next  *Node
// 	prev  *Node
// }
// type LinkedList struct {
// 	Head *Node
// 	Tail *Node
// 	cnt  int
// }

// func (l *LinkedList) Print() {
// 	node := l.Head
// 	fmt.Print("HEAD: ")
// 	for node != nil {
// 		fmt.Printf("%v ", node.value)
// 		node = node.next
// 	}
// 	fmt.Println()
// }
// func (l *LinkedList) Put(elems ...string) {
// 	for _, elem := range elems {
// 		newNode := &Node{value: elem}
// 		// Вставка первого элемента
// 		if l.cnt == 0 {
// 			l.Head, l.Tail = newNode, newNode
// 		} else {
// 			l.Tail.next = newNode
// 			newNode.prev = l.Tail
// 			l.Tail = newNode
// 		}
// 		l.cnt++
// 	}
// }
// func (l *LinkedList) Take() *Node {
// 	node := l.Head
// 	if l.cnt > 1 {
// 		l.Head = l.Head.next
// 		l.Head.prev = nil
// 		node.next = nil
// 		l.cnt--
// 	} else if l.cnt == 1 {
// 		l.Head, l.Tail = nil, nil
// 		l.cnt--
// 	}
// 	return node
// }
