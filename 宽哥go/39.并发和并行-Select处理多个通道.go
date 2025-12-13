package main

// myfunc cookDish(chef, dishName string, c chan string) { //做菜函数
// 	fmt.Printf("厨师:\033[1;32;40m%s\033[0m   正在做:\033[1;32;40m%s\033[0m\n", chef, dishName)
// 	time.Sleep(time.Second * 5)
// 	//做好的菜放进管道内
// 	c <- dishName
// }
// myfunc main() {
// 	//定义2个channel存储不同的数据
// 	chef1 := make(chan string)
// 	chef2 := make(chan string)
// 	go cookDish("chef1", "烤鸭", chef1)
// 	defer close(chef1) //关闭通道
// 	defer close(chef2) //关闭通道
// 	go cookDish("chef2", "开水白菜", chef2)

// 	//等待获取数据
// 	select { //select包含多个case语句，每个case用于接收一个通道的数据 当某个通道有个数据之后，执行对应的case语句
// 	case dish := <-chef1:
// 		fmt.Println("厨师chef1已经做好了:", dish)
// 	case dish := <-chef2:
// 		fmt.Println("厨师chef1已经做好了:", dish)
// 	case <-time.After(time.Second * 3):
// 		fmt.Println("高峰时段上菜时间慢")
// 	}
// }
