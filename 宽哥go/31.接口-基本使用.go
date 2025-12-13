package main

//1.创建数据库结构体，用来存放数据库连接信息
// type DBConfig struct {
// 	Host     string
// 	Port     int
// 	Database string
// 	User     string
// 	Password string
// }

// myfunc main() {
// 	//2.声明一个数据库实例
// 	db := DBConfig{"127.0.0.1", 3306, "interface_test", "root", "123456"}
// 	fmt.Println("数据库配置:", db)
// 	//3.插入一条数据
// 	fmt.Println("在mysql中插入一条数据:db.row('insert into test').Rows()")

// 	//4.换成postgresql
// 	dbPg := DBConfig{"127.0.0.1", 3306, "interface_test", "root", "123456"}
// 	fmt.Println("数据库配置:", dbPg)
// 	fmt.Println("在mysql中插入一条数据:db.QueryRow('insert into test')")

// }
