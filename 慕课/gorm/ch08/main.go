package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type NewUser struct {
	ID           uint           `gorm:"primarykey;comment:用户唯一ID"`
	MyName       string         `gorm:"column:name;comment:用户名称"`
	Email        *string        `gorm:"comment:邮箱"`
	Age          uint8          `gorm:"comment:年龄"`
	Birthday     *time.Time     `gorm:"comment:生日"`
	MemberNumber sql.NullString `gorm:"comment:会员号码"`
	ActivedAt    sql.NullTime   `gorm:"comment:激活时间"`
	CreatedAt    time.Time      `gorm:"comment:创建时间"`
	UpdatedAt    time.Time      `gorm:"comment:更新时间"`
	DeleteAt     gorm.DeletedAt `gorm:"comment:"删除时间"`
}

func main() {
	// 1.数据源云原生操练环境
	dsn := "root:123456@tcp(192.168.50.100:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, //日志级别
			Colorful:      true,        //禁用彩色打印
		},
	)
	// 全局logger 打印执行的sql语句
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&NewUser{})
	// db.Model(&NewUser{}).Create([]map[string]interface{}{
	// 	{"Name": "laoliu", "Age": 18},
	// 	{"Name": "laoliu2", "Age": 19},
	// })
	// var users = []NewUser{{MyName: "laoliu1", Age: 18}, {MyName: "laoliu2", Age: 18}, {MyName: "laoliu3", Age: 18}}
	// for _, user := range users {
	// 	fmt.Println(user.ID)
	// }
	// 软删除
	db.Delete(&NewUser{}, 1)
	var users []NewUser
	db.Find(&users)
	for _, user := range users {
		fmt.Println(user.ID)
	}
	// 硬删除
	db.Unscoped().Delete(&NewUser{ID: 2})
}
