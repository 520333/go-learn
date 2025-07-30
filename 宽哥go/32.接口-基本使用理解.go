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

// func (m MySQL) Select(data string) error {
// 	fmt.Println("从MySQL查询数据:", data)
// 	return nil
// }

// func (m MySQL) Insert(data string) error {
// 	fmt.Println("插入数据到MySQL:", data)
// 	return nil

// }

// func (m MySQL) Update(data string) error {
// 	fmt.Println("更新数据到MySQL:", data)
// 	return nil
// }

// func (m MySQL) Delete(data string) error {
// 	fmt.Println("删除数据:", data)
// 	return nil
// }

// func (m PostGreSQL) Select(data string) error {
// 	fmt.Println("从postgresql查询数据:", data)
// 	return nil
// }

// func (m PostGreSQL) Insert(data string) error {
// 	fmt.Println("插入数据到postgresql:", data)
// 	return nil

// }

// func (m PostGreSQL) Update(data string) error {
// 	fmt.Println("更新数据到postgresql:", data)
// 	return nil
// }

// func (m PostGreSQL) Delete(data string) error {
// 	fmt.Println("postgresql删除数据:", data)
// 	return nil
// }

// func main() {
// 	db := DBConfig{"127.0.0.1", 3306, "interface_test", "root", "123456"}
// 	var dbCommonInterface DBCommon
// 	var m MySQL
// 	var pg PostGreSQL
// 	m.config = db
// 	m.charset = "utf8"
// 	dbCommonInterface = m
// 	dbCommonInterface.Select("select")
// 	dbCommonInterface.Insert("insert")
// 	dbCommonInterface.Update("update")
// 	dbCommonInterface.Delete("delete")

// 	pg.config = db
// 	pg.charset = "utf8"
// 	dbCommonInterface = pg
// 	dbCommonInterface.Select("select")

// }
