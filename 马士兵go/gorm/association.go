package gorm

import "log"

func StdAssocModel() {
	if err := DB.AutoMigrate(&Author{}, &Essay{}, &Tag{}, &EssayMate{}); err != nil {
		log.Fatalln(err)
	}
	log.Println("migrate successful")

}
