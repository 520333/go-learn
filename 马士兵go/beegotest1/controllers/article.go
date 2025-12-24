package controllers

import (
	"beegotest1/models"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Read() {
	article := &models.Article{} // 模型对象
	article.Id = 200
	isCreate, id, err := models.OrmDft.ReadOrCreate(article, "id")
	if err != nil {
		c.Data["json"] = map[string]any{
			"code":    1,
			"message": err.Error(),
		}

		_ = c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]any{
		"code":      0,
		"message":   "ReadOrCreate SUCCESS",
		"is_create": isCreate,
		"id":        id,
	}
	_ = c.ServeJSON()
}

func (c *ArticleController) Create() {

	// 1.做测试数据切片
	articles := make([]*models.Article, 100)
	for i := range articles {
		article := &models.Article{}
		article.Subject = fmt.Sprintf("批量插入测试标题-%d", i)
		articles[i] = article
	}
	// 2.批量插入
	var bulk = 10 // 一次性插入10条
	rows, err := models.OrmDft.InsertMulti(bulk, articles)
	//models.OrmDft.InsertOrUpdate(article)
	if err != nil {
		c.Data["json"] = map[string]any{
			"code":    1,
			"message": err.Error(),
		}
		_ = c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]any{
		"code":    0,
		"message": "insert multi SUCCESS",
		"rows":    rows,
	}
	_ = c.ServeJSON()
}

func (c *ArticleController) Get() {
	article := &models.Article{} // 模型对象
	article.Id = 1
	err := models.OrmDft.ReadForUpdate(article)
	if err != nil {
		c.Data["json"] = map[string]any{
			"code":    1,
			"message": err.Error(),
		}

		_ = c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]any{
		"code":    0,
		"message": "Read SUCCESS",
		"article": article,
	}
	_ = c.ServeJSON()
}

func (c *ArticleController) Post() {
	article := &models.Article{} // 模型对象
	article.Subject = c.GetString("subject")

	// 调用插入方法 完成数据插入
	id, err := models.OrmDft.Insert(article)
	if err != nil {
		c.Data["json"] = map[string]any{
			"code":    1,
			"message": err.Error(),
		}

		_ = c.ServeJSON()
		return
	}
	// 做出响应
	c.Data["json"] = map[string]any{
		"code":    0,
		"message": "Insert SUCCESS",
		"id":      id,
	}
	_ = c.ServeJSON()
}

func (c *ArticleController) Delete() {
	article := &models.Article{} // 模型对象
	article.Id = 1
	id, err := models.OrmDft.Delete(article)
	if err != nil {
		c.Data["json"] = map[string]any{
			"code":    1,
			"message": err.Error(),
		}

		_ = c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]any{
		"code":    0,
		"message": "Delete SUCCESS",
		"rows":    id,
	}
	_ = c.ServeJSON()
}

func (c *ArticleController) Put() {
	article := &models.Article{} // 模型对象
	article.Id = 2
	err := models.OrmDft.Read(article)
	if err != nil {
		c.Data["json"] = map[string]any{
			"code":    1,
			"message": err.Error(),
		}
		_ = c.ServeJSON()
		return
	}

	article.Summary = "beegoSummary"
	article.Subject = "beego Subject"
	id, err := models.OrmDft.Update(article)
	if err != nil {
		c.Data["json"] = map[string]any{
			"code":    1,
			"message": err.Error(),
		}

		_ = c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]any{
		"code":    0,
		"message": "Update SUCCESS",
		"rows":    id,
	}
	_ = c.ServeJSON()
}
