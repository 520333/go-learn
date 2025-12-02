package gorm

import "log"

func GetByPK() {
	DB.AutoMigrate(&Content{}, &ContentStrPK{})
	// 查询单条
	c := Content{}
	if err := DB.First(&c, 10).Error; err != nil {
		log.Println(err)
	}
	// 字符串类型主键
	cStr := ContentStrPK{}
	if err := DB.First(&cStr, "id = ?", "some pk").Error; err != nil {
		log.Println(err)
	}
	// 查询多条
	var cs []Content
	if err := DB.Find(&cs, []uint{10, 11, 12}).Error; err != nil {
		log.Println(err)
	}

	// 字符串类型的主键
	var cStrs []ContentStrPK
	if err := DB.Find(&cStrs, "id IN ?", []string{"some", "pk", "item"}).Error; err != nil {
		log.Println(err)
	}
}
