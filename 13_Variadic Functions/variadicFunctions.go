package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func greet(names ...string) {
	fmt.Println(names)
	for i := 0; i < len(names); i++ {
		fmt.Println("hello", names[i])
	}
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

	names := []string{"dave", "mike", "stephen"}
	greet(names...)
}




