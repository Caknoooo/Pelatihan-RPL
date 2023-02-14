package main

import (
	"fmt"
	"math/rand"
)

// Untuk membuat node
type node_t struct{
	data interface{}
	next *node_t
}

// Untuk membuat list atau dengan kata lain menghubungkan node satu dengan node lainnya
type list_t struct{
	_size uint
	head *node_t
}

// initialize list 
func(l *list_t) init(){
	l.head = nil
	l._size = 0
}

// Fungsi untuk mengecek apakah list kosong atau tidak
func list_empty(l *list_t) (bool){
	return l.head == nil
}

// Menambahkan data dari depan. time complexity O(1)
func(l *list_t) push_front(d interface{}){
	// Memory allocation for list
	newNode := &node_t{data: d, next: nil}
	l._size++

	if(list_empty(l)){
		newNode.next = nil
	} else{
		newNode.next = l.head
	}

	l.head = newNode
}

// Menghapus data dari depan
func(l *list_t) pop_front(){
	if(!list_empty(l)){
		temp := l.head
		l.head = l.head.next
		temp = temp.next
		l._size--;
	}
}

// Menambahkan data dari belakang, time complexity O(n)
func (l *list_t) push_back(d interface{}){
	// Memory allocation for list
	newNode := &node_t{data: d, next: nil}
	l._size++

	if(list_empty(l)){
		l.head = newNode
	} else{
		currNode := l.head
		for currNode.next != nil{
			currNode = currNode.next
		}
		currNode.next = newNode
	}
}

// Menghapus data dari belakang
func (l *list_t) pop_back(){
	if(!list_empty(l)){
		nextNode := l.head.next
		currNode := l.head

		// Jika data cuman 1
		if(currNode.next == nil){
			currNode = nil
			l.head = nil
			l._size--
			return
		}

		// Jika data lebih dari 1
		for nextNode.next != nil {
			currNode = nextNode
			nextNode = nextNode.next
		}
		currNode.next = nil
		l._size--
	}
}

// Menambahkan data ke dalam index tertentu
func (l *list_t) insertAt(index uint, data_t interface{}){
	// Jika list kosong atau menambahkan daya di luar size nya maka masukkan ke paling belakang
	if(list_empty(l) || index >= l._size){
		l.push_back(data_t)
		return
	} else if(index <= 0){ // Jika mau menambahkan ke index awal
		l.push_front(data_t)
		return
	} 

	newNode := &node_t{data: data_t, next: nil}
	temp := l.head
	var idx uint = 0
	
	for temp.next !=  nil && idx < index - 1{
		temp = temp.next
		idx++
	}
	newNode.data = data_t
	newNode.next = temp.next
	temp.next = newNode
	l._size++
}

// Menghapus list dalam index tertentu
func(l *list_t) removeAt(index uint){
	if(!list_empty(l)){
		if(index >= l._size){
			l.pop_back()
			return
		} else if(index <= 0){
			l.pop_front()
			return
		}

		temp := l.head
		var idx uint = 0
		for temp.next != nil && idx < index - 1{
			temp = temp.next
			idx++
		}
		nextTo := temp.next.next
		temp.next = nil
		temp.next = nextTo
		l._size--
	}
}

// Menghapus data menggunakan value yang akan dicari
func (l *list_t) remove(data_t interface{}){
	if(!list_empty(l)){
		temp :=  l.head
		var currNode *node_t;

		if(temp.data == data_t ){
			l.pop_front()
			return
		}

		for temp.next != nil && temp.data != data_t{
			currNode = temp
			temp = temp.next
		}

		if(temp == nil){
			return
		}

		currNode.next = temp.next
		temp = nil
		l._size--
	}
}

func(l *list_t) getAt(index uint)(interface{}){
	if(!list_empty(l)){
		temp := l.head
		var idx uint = 0;

		for temp.next != nil && idx < index{
			temp = temp.next
			idx++
		}
		return temp.data
	}
	return 0
}

// Mengeluarkan data paling depan
func show_front(l *list_t) (interface{}) {
	if(!list_empty(l)){
		return l.head.data
	}
	return 0
}

// Mengeluarkan data paling belakang
func show_back(l *list_t) (interface{}){
	if(!list_empty(l)){
		temp := l.head
		for temp.next != nil{
			temp = temp.next
		}
		return temp.data
	}
	return 0
}
 
// Mengeluarkan seluruh data, time complexity O(n)
func show(l *list_t){
	currNode := l.head
	for currNode != nil{
		fmt.Printf("-> %v ", currNode.data)
		currNode = currNode.next
	}
	fmt.Printf("\nSize data : %v", l._size);
}

func main(){
	s1 := list_t{}
	s1.init()


	for i := 0; i < 10; i++{
		s1.push_back(rand.Intn(500))
	}

	s1.push_back(200)
	s1.push_front(100)
	show(&s1)

	fmt.Println()
	s1.insertAt(3, 99)
	show(&s1)

	fmt.Println()
	fmt.Println("Nilai pada posisi ke-2 adalah", s1.getAt(2))
	// s1.removeAt(3)
	// s1.removeAt(1)
	// s1.remove(200)
	show(&s1)

	fmt.Println()
	fmt.Println("Nilai depan : ", show_front(&s1))
	fmt.Println("Nilai belakang : ", show_back(&s1))
	fmt.Println("Nilai pada posisi ke-2 adalah", s1.getAt(2))
	s1.removeAt(2)
	show(&s1)
}