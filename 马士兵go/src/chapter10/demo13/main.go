package main

import "fmt"

type Animal struct {
	Age    int
	Weight float32
}

func (an *Animal) Shout() {
	fmt.Println("大声喊叫")
}
func (an *Animal) ShowInfo() {
	fmt.Printf("动物年龄:%v 体重:%v \n", an.Age, an.Weight)
}

type Cat struct {
	Name string
	Age  int
	Animal
}

func (c *Cat) scratch() {
	fmt.Printf("我是%v 我可以挠人 \n", c.Name)
}

func (c *Cat) ShowInfo() {
	fmt.Printf("~~~动物年龄:%v 体重:%v \n", c.Age, c.Weight)
}

type Dog struct {
	*Animal
	Name string
}

func (d *Dog) scratch() {
	fmt.Printf("我是%v 我可以挠人 \n", d.Name)
}

// type Dog struct {
// 	Animal
// }

func main() {
	var cat *Cat = &Cat{}
	cat.Name = "小花"
	cat.Animal.Age = 21
	cat.Animal.Weight = 21.12
	cat.Shout()
	cat.Animal.ShowInfo()

	// var dog *Dog = &Dog{Name: "小白", Animal: &Animal{Age: 11, Weight: 10.2}}
	// dog.Shout()
	// dog.ShowInfo()
}
