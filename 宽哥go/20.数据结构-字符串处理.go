package main

// myfunc main() {
// 	s := "\t\txxx\n"
// 	fmt.Println("双引号字符串:", s) //stdout: 双引号字符串:           xxx
// 	s2 := `\t\txxx\n`
// 	fmt.Println("反引号字符串:", s2) //stdout: 反引号字符串: \t\txxx\n
// 	s3 := `
// 	我是chuang,爱好Linux、云原生，每天花大量的时间学习kubenertes、docker技术栈
// 	简单学过Java、python、go语言，很爱devops。
// 	项目经历:从0到1手拿把掐玩转基于微服务下的k8s生产级系统 林被做到了。
// 	作为一名出色的云原生小学徒，我的口号是：上代码,上报错信息，上页面显示。抱歉你找别人我不会。
// 	`
// 	fmt.Println("这是多行文本:", s3)

// 	s4 := "dawn.chuang"
// 	s5 := "obiwan master"
// 	s4Length := len(s4)
// 	s5Length := len(s5)
// 	fmt.Println("s4的长度:", s4Length, "s5的长度:", s5Length) //stdout: s4的长度: 11 s5的长度: 13

// 	//字符串截取
// 	fmt.Println("前两位:", s4[:2]) //stdout: 前两位: da

// 	s7 := "chuang"
// 	fmt.Println("小写转大写:", strings.ToUpper(s7)) //stdout: 小写转大写: CHUANG
// 	fmt.Println("首字母大写:", strings.Title(s7))   //stdout: 首字母大写: Chuang

// 	s8 := "Yoda"
// 	fmt.Println("转小写:", strings.ToLower(s8)) //stdout: 转小写: yoda

// 	fmt.Println("查看字符串是否包含ch这个元素:", strings.Contains(s7, "ch"))    //stdout: 查看字符串是否包含ch这个元素: true
// 	fmt.Println("查看字符串是否包含任意一个字符:", strings.ContainsAny(s7, "aa")) //stdout: 查看字符串是否包含任意一个字符: true

// 	fmt.Println("忽略大小写比较:", strings.EqualFold(s7, s8)) //stdout: 查看字符串是否包含任意一个字符: true

// 	s9 := "chuang and dawn is me,my age is 20"
// 	fmt.Println("n在字符串中出现了:", strings.Count(s9, "n")) //stdout: n在字符串中出现了: 3
// 	fmt.Println("使用逗号分隔:", strings.Split(s9, ","))    //stdout: 使用逗号分隔: [chuang and dawn is me my age is 20]

// 	s9SplitAfter := strings.SplitAfter(s9, ",")
// 	fmt.Println("使用逗号拆分字符串,并且保留逗号:", s9SplitAfter) //stout: 使用逗号拆分字符串,并且保留逗号: [chuang and dawn is me, my age is 20]

// 	slice1 := []string{"obiwan", "yoda", "rey", "bb8"}
// 	fmt.Println("拼接字符串:", strings.Join(slice1, ",")) //stdout: 拼接字符串: obiwan,yoda,rey,bb8

// 	s10 := "我是一个中国人，我非常热爱福建"
// 	fmt.Println("字符串是以'我'开头的:", strings.HasPrefix(s10, "我")) //stdout: 字符串是以'我'开头的: true
// 	fmt.Println("字符串是以'爱'结尾的:", strings.HasSuffix(s10, "爱")) //stdout: 字符串是以'爱'结尾的: false

// 	fmt.Println("打印重复字符串:", strings.Repeat("我", 5)) //stdout: 打印重复字符串: 我我我我我

// 	s11 := "13131"
// 	fmt.Println("把3替换为chuang:", strings.ReplaceAll(s11, "3", "chuang")) //stdout: 把3替换为chuang: 1chuang1chuang1
// 	fmt.Println("把3替换为阿宝:", strings.Replace(s11, "3", "阿宝", 1))         //stdout: 把3替换为阿宝: 1阿宝131

// 	s12 := "  dawn  "
// 	fmt.Println("去掉字符串的前后空格:", strings.Trim(s12, " ")) //stdout: 去掉字符串的前后空格: dawn

// }
