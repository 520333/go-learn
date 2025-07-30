package main

/*
	type 结构体的名称{
		属性1 属性的类型
		属性2 属性的类型
	}
*/
// type People struct {
// 	Name    string
// 	Age     int
// 	Address string
// 	Gender  string
// 	Hobby   []string
// }

// func main() {
// 	var p People
// 	p.Name = "chuang"
// 	p.Age = 20
// 	p.Address = "福建省泉州市中正区一段122号"
// 	p.Hobby = []string{"freestyle", "photography", "piano"}
// 	fmt.Println(p, "类型:", reflect.TypeOf(p)) //stdout: {dawn 20 福建省泉州市中正区一段122号 } 类型: main.People

// 	p2 := People{
// 		Name:    "chuang",
// 		Age:     20,
// 		Address: "福建省泉州市中正区上直街99号",
// 		Gender:  "男",
// 	}
// 	fmt.Println(p2, "类型:", reflect.TypeOf(p2)) //stdout: {chuang 20 福建省泉州市中正区上直街99号 男} 类型: main.People

// 	var yoda People = People{"Yoda Jedi Master", 1500, "In a galaxy far, far away....", "Male", []string{"光剑", "倒装句"}}
// 	fmt.Println("My favorite star wars Jedi Master is:", yoda.Name, "homeland:", yoda.Address)
// 	yoda.Address = "Unknown"
// 	p = yoda
// 	fmt.Println(p)
// 	fmt.Println("yodo和chuang是否相等:", p.Name == yoda.Name)

// }
