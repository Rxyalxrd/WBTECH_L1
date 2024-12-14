package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской
структуры Human (аналог наследования).
*/

type human struct {
	name string
}

type action struct {
	human
}

func (msg *human) say_before() {
	fmt.Println("Hey," + " " + msg.name + " " + "is Python developer!")
}

func (msg *action) say_after() {
	fmt.Println("Now," + " " + msg.name + " " + "will be become a Golang developer!")
}

func main() {

	sample := action{
		human{name: "Max"},
	}

	sample.say_before()
	sample.say_after()
}
