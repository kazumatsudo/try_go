package main

import (
	"fmt"
	"time"
)

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

type animal interface {
	bark() string
}

type dog struct {
	name string
}

func (d dog) bark() string {
	return "bow"
}

type cat struct {
	name string
}

func (d cat) bark() string {
	return "mew"
}

func task(second int) {
	time.Sleep(time.Duration(second) * time.Second)
	fmt.Printf("finish: %d s\n", second)
}

func main() {
	{ // 並行処理
		go task(1)
		go task(2)
		go task(3)
		// 並行処理で終わってしまうため、待機
		task(3)
	}

	{ // インタフェース
		animals := []animal{dog{name: "inu"}, cat{name: "neko"}}
		for _, animal := range animals {
			fmt.Printf("%s: %s\n", animal, animal.bark())
		}
	}

	{ // 構造体
		u1 := new(user)
		u1.age = 10
		u1.name = "taro"

		// name: taro, age: 10
		u1.show()

		u1.incrementAge()
		// name: taro, age: 11
		u1.show()
	}
}
