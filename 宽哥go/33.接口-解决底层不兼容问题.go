package main

/*
type 接口名称 interface {
	方法名称(传参类型) 返回值
	方法名称(传参类型) 返回值
	方法名称(传参类型) 返回值
}
*/
// type DBCommon interface {
// 	Select(string) error
// 	Insert(string) error
// 	Update(string) error
// 	Delete(string) error
// }

// type DBConfig struct {
// 	Host     string
// 	Port     int
// 	Database string
// 	User     string
// 	Password string
// }

// //定义类型去实现接口
// type MySQL struct {
// 	config  DBConfig
// 	charset string
// }

// type PostGreSQL struct {
// 	config  DBConfig
// 	charset string
// }

// myfunc (m MySQL) Select(data string) error {
// 	fmt.Println("从MySQL查询数据:", data)
// 	return nil
// }

// myfunc (m MySQL) Insert(data string) error {
// 	fmt.Println("插入数据到MySQL:", data)
// 	return nil

// }

// myfunc (m MySQL) Update(data string) error {
// 	fmt.Println("更新数据到MySQL:", data)
// 	return nil
// }

// myfunc (m MySQL) Delete(data string) error {
// 	fmt.Println("删除数据:", data)
// 	return nil
// }

// myfunc (m PostGreSQL) Select(data string) error {
// 	fmt.Println("从postgresql查询数据:", data)
// 	return nil
// }

// myfunc (m PostGreSQL) Insert(data string) error {
// 	fmt.Println("插入数据到postgresql:", data)
// 	return nil

// }

// myfunc (m PostGreSQL) Update(data string) error {
// 	fmt.Println("更新数据到postgresql:", data)
// 	return nil
// }

// myfunc (m PostGreSQL) Delete(data string) error {
// 	fmt.Println("postgresql删除数据:", data)
// 	return nil
// }

// myfunc main() {
// 	dbType := "pgsql"
// 	// db := DBConfig{"127.0.0.1", 3306, "interface_test", "root", "123456"}
// 	var dbCommonInterface DBCommon
// 	if dbType == "mysql" {
// 		var m MySQL
// 		dbCommonInterface = m
// 	} else {
// 		var pg PostGreSQL
// 		dbCommonInterface = pg
// 	}
// 	// var m MySQL
// 	// var pg PostGreSQL
// 	dbCommonInterface.Select("select")
// 	dbCommonInterface.Insert("insert")
// 	dbCommonInterface.Update("update")
// 	dbCommonInterface.Delete("delete")
// }
