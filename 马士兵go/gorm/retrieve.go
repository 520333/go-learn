package gorm

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

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

func GeOne() {
	c := Content{}
	if err := DB.First(&c, "id > ?", 42).Error; err != nil {
		log.Println(err)
	}
	o := Content{}
	if err := DB.Last(&o, "id > ?", 42).Error; err != nil {
		log.Println(err)
	}
	n := Content{}
	if err := DB.Take(&n, "id > ?", 42).Error; err != nil {
		log.Println(err)
	}
	f := Content{}
	if err := DB.Limit(1).Find(&f, "id > ?", 42).Error; err != nil {
		log.Println(err)
	}
	fs := Content{}
	if err := DB.Find(&fs, "id > ?", 42).Error; err != nil {
		log.Println(err)
	}
}

func GetToMap() {
	c := map[string]any{}
	if err := DB.Model(&Content{}).First(&c, 13).Error; err != nil {
		log.Println(err)
	}
	fmt.Println(c, c["id"] == 13)
	if c["id"].(uint) == 13 {
		fmt.Println("id bingo")
	}
	// time类型处理
	fmt.Println(c["created_at"])
	t, _ := time.Parse("2006-01-02 15:04:05.00 -0700 CST", "2025-12-02 23:23:34.22 +0800 CST")
	if c["created_at"].(time.Time) == t {
		fmt.Println("created_at bingo")
	}

	// 多条
	var cs []map[string]any
	if err := DB.Model(&Content{}).Find(&cs, []uint{13, 14, 15}).Error; err != nil {
		log.Println(err)
	}
	for _, c := range cs {
		fmt.Println(c["id"], c["subject"].(string), c["created_at"].(time.Time))
	}
}

func GetPluck() {
	//使用切片存储
	var subject []sql.NullString
	if err := DB.Model(&Content{}).Pluck("subject", &subject).Error; err != nil {
		fmt.Println(err)
	}
	for _, subject := range subject {
		if subject.Valid {
			fmt.Println(subject.String)
		} else {
			fmt.Println("[NULL]")
		}
	}
}

func GetPluckExp() {
	//使用切片存储
	var subjects []sql.NullString
	if err := DB.Model(&Content{}).Pluck("concat(subject,'-',likes)", &subjects).Error; err != nil {
		fmt.Println(err)
	}
	for _, subject := range subjects {
		fmt.Println(subject.String)
	}
}

func GetSelect() {
	var c Content
	if err := DB.Select("subject", "likes", "concat(subject,'-',views)").First(&c, 13).Error; err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", c)
}

func GetDistinct() {
	var c []Content
	if err := DB.Distinct("*").Find(&c, 13).Error; err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", c)
}

func WhereMethod() {
	var cs []Content
	// inline条件，内联条件
	//if err := DB.Find(&cs, "likes > ? AND subject like ?", 100, "gorm%").Error; err != nil {
	//	log.Fatalln(err)
	//}
	//query := DB.Where("likes > ?", 100)
	//subject := "gorm"
	//if subject != "" {
	//	query.Where("subject like ?", subject+"%")
	//}
	//if err := query.Find(&cs).Error; err != nil {
	//	log.Fatalln(err)
	//}

	// OR 逻辑运算
	//query := DB.Where("likes > ?", 100)
	//subject := "gorm"
	//if subject != "" {
	//	query.Or("subject like ?", subject+"%")
	//}
	//if err := query.Find(&cs).Error; err != nil {
	//	log.Fatalln(err)
	//}
	// Not 逻辑运算
	query := DB.Where("likes > ?", 100)
	subject := "gorm"
	if subject != "" {
		query.Not("subject like ?", subject+"%")
		query.Or(DB.Not("subject like ?", subject+"%"))
	}
	if err := query.Find(&cs).Error; err != nil {
		log.Fatalln(err)
	}
}

func WhereType() {
	var cs []Content
	// (1 or 2) and (3 and (4 or 5))
	//condA := DB.Where("likes > ?", 10).
	//	Or("likes <= ?", 100)
	//condB := DB.Where("views > ?", 100).
	//	Where(DB.Where("views <= ?", 200).
	//		Or("subject like ?", "gorm%"))
	//query := DB.Where(condA).Where(condB)
	//if err := query.Find(&cs).Error; err != nil {
	//	log.Fatalln(err)
	//}

	// Map类型构建查询
	//query := DB.Where(map[string]any{
	//	"views": 100,
	//	"id":    []uint{1, 2, 3, 4, 5},
	//})
	query := DB.Where(Content{
		Views:   100,
		Subject: "GORM",
	})
	if err := query.Find(&cs).Error; err != nil {
		log.Fatalln(err)
	}

}

func PlaceHolder() {
	var cs []Content
	// 匿名
	//query := DB.Where("likes = ? AND subject like ?", 100, "gorm%")
	query := DB.Where("likes = @like AND subject like @subject",
		sql.Named("subject", "gorm%"),
		sql.Named("like", 1000))

	if err := query.Find(&cs).Error; err != nil {
		log.Fatalln(err)
	}

}

func OrderBy() {
	var cs []Content
	ids := []uint{2, 3, 1}
	//query := DB.Order("Field(id,2,3,1)")
	query := DB.Clauses(clause.OrderBy{
		Expression: clause.Expr{
			SQL:                "Field(id,?)",
			Vars:               []any{ids},
			WithoutParentheses: true,
		},
	})
	if err := query.Find(&cs, ids).Error; err != nil {
		log.Fatalln(err)
	}
	for _, c := range cs {
		fmt.Println(c.ID)
	}
}

// Pager 定义分页必要数据结构
type Pager struct {
	Page, PageSize int
}

// 默认的值
const (
	DefaultPage     = 1
	DefaultPageSize = 12
)

// Pagination 翻页查询
func Pagination(pager Pager) {
	// 确定 offset 和 pageSize
	page := DefaultPage
	if pager.Page != 0 {
		page = pager.Page
	}
	pageSize := DefaultPageSize
	if pager.PageSize != 0 {
		pageSize = pager.PageSize
	}
	// 计算offset
	// page pageSize offset
	// 1,10,0
	// 2,10,10
	// 3,10,20
	// 4,10,30
	offset := pageSize * (page - 1)
	var cs []Content
	if err := DB.Offset(offset).Limit(pageSize).Find(&cs).Error; err != nil {
		log.Fatalln(err)
	}
}

// Paginate 分页复用
func Paginate(pager Pager) func(db *gorm.DB) *gorm.DB {
	var cs []Content
	page := DefaultPage
	if pager.Page != 0 {
		page = pager.Page
	}
	pageSize := DefaultPageSize
	if pager.PageSize != 0 {
		pageSize = pager.PageSize
	}
	offset := pageSize * (page - 1)
	if err := DB.Offset(offset).Limit(pageSize).Find(&cs).Error; err != nil {
		log.Fatalln(err)
	}
	return func(db *gorm.DB) *gorm.DB {
		// 使用闭包的变量实现翻页的逻辑
		return db.Offset(offset).Limit(pageSize)
	}
}

func PaginationScope(pager Pager) {
	var cs []Content
	if err := DB.Scopes(Paginate(pager)).Find(&cs).Error; err != nil {
		log.Fatalln(err)
	}
	for _, c := range cs {
		fmt.Println(c.ID, c.Subject, c.Likes)
	}

	var ps []Post
	if err := DB.Scopes(Paginate(pager)).Find(&ps).Error; err != nil {
		log.Fatalln(err)
	}
}

func GroupHaving() {
	DB.AutoMigrate(&Content{})
	type Result struct {
		AuthorID uint // 分组字段
		// 合计字段
		TotalViews int
		TotalLikes int
		AvgViews   float64
	}
	// 执行分组合计过滤查询
	var rs []Result
	if err := DB.Model(&Content{}).
		Select("author_id", "SUM(views) as total_views", "SUM(likes) as total_likes", "AVG(views) as avg_views").
		Group("author_id").Having("total_views > ?", 99).Find(&rs).Error; err != nil {
		log.Fatalln(err)
	}
	for _, r := range rs {
		fmt.Println(r.AuthorID, r.TotalViews, r.TotalLikes, r.AvgViews)
	}
}

func Count(pager Pager) {
	query := DB.Model(&Content{}).
		Where("likes > ?", 99)
	var count int64
	if err := query.Count(&count).Error; err != nil {
		log.Fatalln(err)
	}
	// 计算总页数 ceil(count / pageSize)
	var cs []Content
	if err := query.Scopes(Paginate(pager)).Find(&cs).Error; err != nil {
		log.Fatalln(err)
	}
}

func Iterator() {
	// 利用DB.Rows() 获取Rows对象
	rows, err := DB.Model(&Content{}).Rows()
	if err != nil {
		log.Fatalln(err)
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	fmt.Println(rows)
	for rows.Next() {
		var c Content
		if err := DB.ScanRows(rows, &c); err != nil {
			log.Fatalln(err)
		}
		fmt.Println(c.Subject)
	}
}

func Locking() {
	var cs []Content
	// SELECT * FROM `go_content` WHERE `go_content`.`deleted_at` IS NULL FOR UPDATE
	if err := DB.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Find(&cs).Error; err != nil {
		log.Fatalln(err)
	}

	// SELECT * FROM `go_content` WHERE `go_content`.`deleted_at` IS NULL FOR SHARE
	if err := DB.
		Clauses(clause.Locking{Strength: "SHARE"}).
		Find(&cs).Error; err != nil {
		log.Fatalln(err)
	}
}
