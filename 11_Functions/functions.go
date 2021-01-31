package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func printName(name string) string {
	return "hello " + name
}

func main() {
	res := plus(1, 2)
	fmt.Println(res)

	res = plusPlus(1, 2, 3)
	fmt.Println(res)

	greeting := printName("moms spaghetti")
	fmt.Println(greeting)

}