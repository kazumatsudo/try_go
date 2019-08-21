package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	u1 := new(user)
	u1.age = 10
	u1.name = "taro"

	fmt.Println(u1)
}
