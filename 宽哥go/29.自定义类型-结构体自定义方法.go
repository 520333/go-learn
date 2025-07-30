package main

// type People struct {
// 	Name    string
// 	Age     int
// 	Address string
// 	Gender  string
// }

// /*
// 小写开头的方法只能在本包内使用
// 大写开头的方法可以被外部包调用
// */
// func (p *People) getInfo() string { //定义一个方法
// 	p.Age = 99
// 	return fmt.Sprintf("当前用户名:%s 性别:%s 年龄:%d 家乡:%s\n", p.Name, p.Gender, p.Age, p.Address)

// }

// func (p *People) Eat(food string) {
// 	fmt.Printf("%s吃了%s\n", p.Name, food)
// }

// func main() {
// 	// var p People = People{"chuang", 20, "福建省", "男"}
// 	// fmt.Printf("当前用户名:%s 性别:%s 年龄:%d 家乡:%s\n", p.Name, p.Gender, p.Age, p.Address)

// 	// var p2 People = People{"庄丽芳", 20, "北京", "女"}
// 	// info := p2.getInfo()
// 	// fmt.Println(info)

// 	// p2.Eat("韭菜盒子")
// 	// p.Eat("青椒肉丝")
// 	var p People = People{"chuang", 20, "福建省", "男"}
// 	info := p.getInfo()
// 	fmt.Println(info)
// 	fmt.Println("查看原有值是否被修改:", p)

// }
