package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func names(nameA, nameB string) (string, string) {
	return nameA, nameB 
}

func calcTwice(num1, num2 int) (int, int) {
	add := num1 + num2
	subtract := num1 - num2
	return add, subtract
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	d, e := names("dave", "mike")
	fmt.Println(d)
	fmt.Println(e)

	added, subtracted := calcTwice(5, 10)
	fmt.Println("added:", added)
	fmt.Println("subtacted:", subtracted)
}