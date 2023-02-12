package main

import (
	"fmt"
	"math/rand"
)

type node_t struct{
	data interface{}
	next *node_t
}

type list_t struct{
	head *node_t
}

func (l *list_t) Insert(d interface{}){
	list := &node_t{data: d, next: nil}

	if(l.head == nil){
		l.head = list
	} else{
		p := l.head
		for p.next != nil{
			p = p.next
		}
		p.next = list
	}
}

func Show(l *list_t){
	p := l.head
	for p != nil{
		fmt.Printf("-> %v", p.data)
		p = p.next
	}
}

func main(){
	s1 := list_t{}
	for i := 0; i < 5; i++{
		s1.Insert(rand.Intn(100))
	}
	s1.Insert(200)
	Show(&s1)
}