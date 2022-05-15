package main

import "fmt"

// the size of the hash table array
const ArraySize = 10

// Hashtable structure
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket is a linked list
type bucket struct {
	head *bucketNode
}

// bucketNode structure
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// insert
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("Already exists!")
	}
}

// search
func (b bucket) search(k string) bool {
	current := b.head
	for current != nil {
		if current.key == k {
			return true
		}
		current = current.next
	}

	return false
}

// delete
func (b *bucket) delete(k string) {
	currentNode := b.head
	if currentNode.key == k {
		b.head = b.head.next
		return
	}
	for currentNode.next != nil {
		if currentNode.next.key == k {
			currentNode.next = currentNode.next.next
			return
		}
		currentNode = currentNode.next
	}
}

// hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// init will create a bucket in each slot
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	testTable := Init()
	fmt.Println(testTable)
	fmt.Println(hash("marcio"), hash("RANDY"))

	testBucket := &bucket{}
	testBucket.insert("RANDY")
	testBucket.insert("MARCIO")

	fmt.Println(testBucket.search("RAND"))
	fmt.Println(testBucket.search("MARCIO"))

	testBucket.insert("RANDY")

	testBucket.delete("RANDY")
	testBucket.delete("MARCIO")
	fmt.Println(testBucket.search("MARCIO"))
	fmt.Println(testBucket.search("RANDY"))

}
