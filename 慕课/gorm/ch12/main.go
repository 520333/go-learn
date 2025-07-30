package main

// many to many
import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User3 struct {
	gorm.Model
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
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
	// 插入数据
	// db.AutoMigrate(&User3{})

	// Languages := []Language{}
	// Languages = append(Languages, Language{Name: "go"})
	// Languages = append(Languages, Language{Name: "java"})
	// Languages = append(Languages, Language{Name: "python"})
	// Languages = append(Languages, Language{Name: "vue"})
	// user := User3{
	// 	Languages: Languages,
	// }
	// db.Create(&user)

	// user := User3{}
	// db.Preload("Languages").First(&user)
	// for _, Language := range user.Languages {
	// 	fmt.Println(Language.ID, Language.Name)
	// }

	var user User3
	db.First(&user)
	var languages []Language
	db.Model(&user).Association("Languages").Find(&languages)
	for _, language := range languages {
		fmt.Println(language.Name)
	}

}
