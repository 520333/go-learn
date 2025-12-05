package gorm

import (
	"context"
	"fmt"
	"log"
)

func ContextTimeOutCancel(ctx context.Context) {
	//ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*20)
	var cs []Content
	if err := DB.WithContext(ctx).Limit(10).Find(&cs).Error; err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cs)
	// SELECT * FROM `go_content` WHERE `go_content`.`deleted_at` IS NULL LIMIT 10

}
