package main

// import (
// 	"crypto/rand"
// 	"crypto/sha1"
// 	"encoding/base64"
// 	"fmt"
// 	"os"
// )

// myfunc generateSSHA(password string) (string, error) {
// 	// 生成 4 字节的随机盐值（slappasswd 默认使用 4 字节）
// 	salt := make([]byte, 4)
// 	_, err := rand.Read(salt)
// 	if err != nil {
// 		return "", err
// 	}

// 	// 创建 SHA-1 哈希对象
// 	hash := sha1.New()
// 	// 将密码和盐值写入哈希
// 	hash.Write([]byte(password))
// 	hash.Write(salt)
// 	// 计算哈希值
// 	hashBytes := hash.Sum(nil)

// 	// 合并哈希值和盐值
// 	hashWithSalt := append(hashBytes, salt...)

// 	// 使用 base64 编码
// 	encoded := base64.StdEncoding.EncodeToString(hashWithSalt)

// 	// 返回 SSHA 格式的字符串
// 	return "{SSHA}" + encoded, nil
// }

// myfunc main() {
// 	// 检查命令行参数
// 	if len(os.Args) < 2 {
// 		fmt.Println("用法: go run script.go <password>")
// 		os.Exit(1)
// 	}

// 	password := os.Args[1]
// 	result, err := generateSSHA(password)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println(result)
// }
