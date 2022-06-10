package main

import "fmt"

type name struct {
	Name     string
	Lastname string
	Age      int
}

func main() {
	S := name{}
	S.setName("aom")
	S.Lastname = "last"
	S.Age = 23
	fmt.Println(S)
}

func (n *name) setName(newName string) {
	n.Name = newName
}
