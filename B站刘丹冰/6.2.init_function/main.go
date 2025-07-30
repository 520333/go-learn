package main

import (
	_ "learn/6.2.init_function/lib1" // _ 是匿名别名 lib1种的init方法会被执行
	// mylib2 "learn/5.init_function/lib2"  // 起别名
	. "learn/6.2.init_function/lib2" //将包导入到当前main方法种
)

func main() {
	//lib1.Lib1Test()
	// mylib2.Lib2Test()
	Lib2Test()
}
