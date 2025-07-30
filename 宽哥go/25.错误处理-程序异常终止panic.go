package main

// func connectDatabase(address string, port int) (string, error) {
// 	//如果address和port为空
// 	if address == "" || port == 0 {
// 		return "", errors.New("无法连接数据库")
// 	} else {
// 		return "连接成功", nil

// 	}
// }

// func main() {
// 	//panic: 可以在异常的时候让程序终止执行,退出程序。或者是程序所强依赖的基础组件不可用
// 	s, err := connectDatabase("", 0)
// 	for { //模拟程序一直执行
// 		time.Sleep(5 * time.Second)
// 		if err != nil {
// 			fmt.Println("连接数据库失败:", err)
// 			panic(err) //终止程序
// 		} else {
// 			fmt.Println(s, "连接数据库ok")
// 		}
// 	}

// }
