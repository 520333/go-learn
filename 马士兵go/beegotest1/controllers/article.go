package controllers

import (
	"beegotest1/models"

	beego "github.com/beego/beego/v2/server/web"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	article := &models.Article{} // 模型对象
	article.Id = 1
	err := models.OrmDft.Read(article)
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
