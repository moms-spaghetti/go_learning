package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

type car struct {
	manufacturer string
	model string
	year int
	isElectric bool
}

func newCar(manufacturer, model string, year int, isElectric bool) *car {
	p := car{
		manufacturer: manufacturer,
		model: model,
		year: year,
		isElectric: isElectric,
	}

	return &p
}

func main() {
	fmt.Println(person{"bob", 20})
	fmt.Println(person{name: "alice", age: 30})
	fmt.Println(person{name: "fred"})
	fmt.Println(&person{name: "ann", age: 40})
	fmt.Println(newPerson("jon"))
	s := person{name: "sean", age: 50}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age)

	vehicle1 := car{"tesla", "model 3", 2018, true}
	fmt.Println(vehicle1)
	fmt.Println(vehicle1.model)
}