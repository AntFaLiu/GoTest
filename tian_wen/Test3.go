package main

import "fmt"

/**
	通过嵌入Animalinto Dog，你得到的唯一东西是Dog.Feed（），它不适用于Dog.Eat（）
	Animaland Dog之间没有任何关系，Animaldon不知道也无法到达狗

	匿名领域就像普通领域一样，没有什么神奇之处。
	嵌入字段和封闭结构之间没有任何关系
	前者不知道后来的存在因此无法达到它
*/

//type Eater interface {
//	Eat()
//}
//type Animal struct {
//	Eater
//}
//
//func (a *Animal) Feed() {
//	// prepare food
//	a.Eater.Eat()
//	// cleanup
//}
//
//type Dog struct {
//	Animal
//}
//
//func (d *Dog) Eat() {
//	fmt.Println("Dog eating...")
//}
//
//func main() {
//	d := Dog{}
//	d.Feed()
//}

type Animal struct {
}

func (a *Animal) Eat() {
	fmt.Println("Eating...")
}
func (a *Animal) Feed() {
	// prepare food
	a.Eat()
	// cleanup
}

type Dog struct {
	Animal
}

func (d *Dog) Eat() {
	fmt.Println("Dog eating...")
}
func main() {
	d := Dog{}
	d.Feed()
}
