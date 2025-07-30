package main

/*
import (
	"fmt"
	"reflect"
)

type AnimalIf interface { //本质

	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	name  string // 名称
	Color string // 颜色
	Type  string // 品种
}

type Dog struct {
	name  string // 名称
	Color string // 颜色
	Type  string // 品种
}

//*********猫类实现animal接口*******
func (kitty Cat) Sleep()           { fmt.Println("kitty is sleep...") }
func (kitty Cat) GetColor() string { return kitty.Color }
func (kitty Cat) GetType() string  { return kitty.Type }

//*********狗类实现animal接口********
func (puppy Dog) Sleep()           { fmt.Println("puppy is sleep...") }
func (puppy Dog) GetColor() string { return puppy.Color }
func (puppy Dog) GetType() string  { return puppy.Type }

func main() {
	//*********猫类********
	var animal AnimalIf
	fmt.Println(reflect.TypeOf(animal))
	animal = Cat{"小白", "白色", "田园猫"}
	animal.Sleep()
	animal.GetColor()
	fmt.Println(animal.GetColor(), animal.GetType())

	//*********狗类********
	animal = Dog{"小黑", "黑色", "阿拉斯加"}
	animal.Sleep()
	animal.GetColor()
	fmt.Println(animal.GetColor(), animal.GetType())
}
*/
