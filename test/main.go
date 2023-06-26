package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

// 值类型接收者
func (p Person) SayHi() {
	fmt.Printf("Hi, my name is %s and I'm %d years old.\n", p.name, p.age)
}

// 指针类型接收者
func (p *Person) SetAge(age int) {
	p.age = age
}
func main() {
	p := Person{name: "Alice", age: 30}
	p.SayHi() // 输出 "Hi, my name is Alice and I'm 30 years old."
	p.SetAge(31)
	fmt.Println(p.age) // 输出31
	p2 := &Person{}
	p2.age = 10
	p2.name = "Apple"
	p2.SayHi()
	fmt.Println(*p2)
}
