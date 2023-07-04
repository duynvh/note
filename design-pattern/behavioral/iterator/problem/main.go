package main

import "fmt"

type Follower interface {
	Receive(message string)
}

type Profile struct {
	name string
}

func (p Profile) Receive(message string) {
	fmt.Printf("%s has received: %s\n", p.name, message)
}

var arrayOfFollowers = []Follower{
	Profile{name: "A"},
	Profile{name: "B"},
	Profile{name: "C"},
}

type LinkedNode struct {
	val Follower
	next *LinkedNode
}

var linkedListOfFollower = &LinkedNode{
	val: Profile{name: "D"},
	next: &LinkedNode{
		val: Profile{name: "E"},
		next: &LinkedNode{
			val: Profile{name: "F"},
			next: nil,
		},
	},
}

type TreeNode struct {
	val Follower
	children []TreeNode
}

var treeOfFollowers = &TreeNode{
	val: Profile{name: "Peter"},
	children: []TreeNode{
		{
			val: Profile{name: "Tom"},
			children: []TreeNode{
				{val: Profile{name: "Mary"}},
				{val: Profile{name: "Vincent"}},
				{val: Profile{name: "Vicky"}},
			},
		},
		{
			val: Profile{name: "Bob"},
			children: []TreeNode{
				{val: Profile{name: "Alice"}},
			},
		},
	},
}

func sendMessageForArray(msg string) {
	for i := range arrayOfFollowers {
		arrayOfFollowers[i].Receive(msg)
	}
}

func sendMessageForLinkedList(msg string) {
	node := linkedListOfFollower

	for node != nil {
		node.val.Receive(msg)
		node = node.next
	}
}

func sendMessageForTree(node *TreeNode, msg string) {
	if node == nil {
		return 
	}

	node.val.Receive(msg)

	for i := range node.children {
		sendMessageForTree(&node.children[i], msg)
	}
}

func main() {
	message := "hello"

	fmt.Println("Sending for array")
	sendMessageForArray(message)

	fmt.Println("Sending for linked list")
	sendMessageForLinkedList(message)

	fmt.Println("Sending for tree")
	sendMessageForTree(treeOfFollowers, message)
}