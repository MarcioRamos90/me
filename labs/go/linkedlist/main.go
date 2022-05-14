package main

import "fmt"

type node struct {
    data int
    next *node
}

type linkedList struct {
    head *node
    length int
}

func (l *linkedList) prepend(n *node) {
    second := l.head
    l.head = n
    l.head.next = second
    l.length++
}

func (l *linkedList) deleteWithValue(value int) {
    if l.head == nil {
        return
    }

    if l.head.data == value {
        l.head = l.head.next
        l.length--
        return
    }

    previousToDelete := l.head

    for previousToDelete.next.data != value && previousToDelete.next != nil {
        previousToDelete = previousToDelete.next
    }

    if previousToDelete.next != nil {
        previousToDelete.next = previousToDelete.next.next
        l.length--
    }
}

func (l *linkedList) posppend(n *node) {
    currentNode := l.head
    for currentNode.next != nil {
        currentNode = currentNode.next
    }
    currentNode.next = n
    l.length++
}

func (l linkedList) printListData() {
    length := l.length
    toPrint := l.head
    fmt.Printf("[ ")
    for l.length != 0 {
        fmt.Printf("%d ", toPrint.data)
        toPrint = toPrint.next
        l.length--
    }
    fmt.Printf("] %d\n", length)
}

func main() {
    mylist := linkedList{}
    node1 := &node{data:34}
    node2 := &node{data:23}
    node3 := &node{data:12}
    node4 := &node{data:77}

    mylist.prepend(node1)
    mylist.prepend(node2)
    mylist.prepend(node3)
    mylist.prepend(node4)

    fmt.Println(":", mylist)

    mylist.printListData()
    mylist.posppend(node3)
    mylist.printListData()
    mylist.deleteWithValue(77)
    mylist.printListData()
}

