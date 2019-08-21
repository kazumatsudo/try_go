package main

import "fmt"

type user struct {
	name string
	age  int
}

func (u user) show() {
	fmt.Printf("name: %s, age: %d\n", u.name, u.age)
}

func (u *user) incrementAge() {
	u.age++
}

func main() {
	u1 := new(user)
	u1.age = 10
	u1.name = "taro"

	// name: taro, age: 10
	u1.show()

	u1.incrementAge()
	// name: taro, age: 11
	u1.show()
}
