package main

import "fmt"

func main(){
	fmt.Println("Hello world")
	var helloWorld string = "Caknoo"
	fmt.Print(helloWorld + "!");

	var hellogan = "Caknii"

	fmt.Println(hellogan);

	var angka int = 10;

	if (angka >= 1) && (angka < 12){
		fmt.Print(angka);
	}

	for i := 0; i < 10; i++{
		fmt.Println("angkat", i);
	}
}