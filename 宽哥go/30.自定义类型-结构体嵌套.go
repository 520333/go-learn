package main

type Phone struct {
	Name  string
	Mode  string
	Price float32
}

type Person struct {
	Name    string
	Age     int
	Address string
	Gender  string
}
type People struct {
	Name    string
	Age     int
	Address string
	Gender  string
	Mobile  Phone
}
type Info struct {
	Person
	Phone
}

// func main() {
// 	// var p Person
// 	// p.Name = "阿宝"
// 	// p.Age = 20
// 	// p.Mobile.Mode = "apple"
// 	// p.Mobile.Price = 8799.00
// 	// fmt.Printf("用户:%s 手机型号:%s\n", p.Name, p.Mobile.Mode)
// 	// var m Phone = Phone{"iphone15Pro", 10999.00}
// 	// var p2 Person
// 	// p2.Name = "dawn"
// 	// p2.Age = 20
// 	// p2.Mobile = m
// 	// fmt.Printf("用户:%s,手机型号:%s\n", p2.Name, p2.Mobile.Mode)

// 	var i Info
// 	i.Person.Name = "宝总"
// 	i.Phone.Name = "小米"
// 	// i.Name = "chuang" //如果成本变量有冲突，就不能直接赋值
// 	i.Age = 20
// 	i.Mode = "小米"
// 	i.Price = 3299.00
// 	fmt.Println("Info:", i)
// }
