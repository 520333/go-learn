package main

/*
1.提高代码复用性
2.实现接口隔离
3.简化接口使用
*/
// type Animal interface {
// 	Eat()
// 	Sleep()
// }

// //Dog和Cat继承Animal
// type Dog interface {
// 	Animal
// }

// type Cat interface {
// 	Animal
// }

// type GenericUser interface {
// 	Login()
// 	Logout()
// }
// type VIP interface {
// 	ToDesk()
// }
// type User interface {
// 	GenericUser
// }
// type VIPUser interface {
// 	GenericUser
// 	VIP
// }

// type UserImpl struct {
// 	Name string
// }
// type VIPUserImpl struct {
// 	Name string
// }

// func (u UserImpl) Login() {
// 	fmt.Println("用户登录:", u.Name)
// }
// func (u UserImpl) Logout() {
// 	fmt.Println("退出登录:", u.Name)
// }

// func (u VIPUserImpl) Login() {
// 	fmt.Println("VIP用户登录:", u.Name)
// }
// func (u VIPUserImpl) Logout() {
// 	fmt.Println("VIP退出登录:", u.Name)
// }

// func (u VIPUserImpl) ToDesk() {
// 	fmt.Println("VIP远程支撑服务:", u.Name)
// }
// func main() {
// 	var userInterface User
// 	var u UserImpl
// 	u.Name = "chuang"
// 	userInterface = u
// 	userInterface.Login()

// 	// vip用户
// 	var vipUserInterface VIPUser
// 	var vipu VIPUserImpl
// 	vipu.Name = "obiwan"
// 	vipUserInterface = vipu
// 	vipUserInterface.Login()
// 	vipUserInterface.ToDesk()
// }
