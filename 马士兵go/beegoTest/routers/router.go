package routers

import (
	"beegoTest/controllers"
	"encoding/json"
	"strings"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {
	web.Router("/", &controllers.MainController{})
	web.Get("/easy", func(ctx *context.Context) {
		_ = ctx.Output.Body([]byte("Go 海绵宝宝!"))
	})
	web.Get("/named-handler", NameHandler)
	//web.Get("/content/", content.Retrieve)
	web.Router("/our", &OurController{})
	// post方法映射到OtherFunc
	ourController := &OurController{}
	web.Router("/our/other", ourController, "get,post:OtherFunc;*:Register")
	//web.Include(&CPController{})

	// 路由自动匹配
	web.AutoRouter(&ContentController{}) //GET  /content/select

	// 路由参数 正则匹配
	web.Get("/article/:id", func(ctx *context.Context) {
		//web.Get("/article/?:id", func(ctx *context.Context) {
		//web.Get("/article/*:id", func(ctx *context.Context) {
		//web.Get("/article/*.*", func(ctx *context.Context) {
		//web.Get("/article/:id:int", func(ctx *context.Context) {
		//web.Get("/article/:id(\\d{4})", func(ctx *context.Context) {
		//web.Get("/article_:id(\\d{4})", func(ctx *context.Context) {
		//ctx.Output.Body([]byte("Router Param! id:" + ctx.Input.Param(":id")))
		//body := fmt.Sprintf("Router Param! *.:%s,.*:%s", ctx.Input.Param(":path"), ctx.Input.Param(":ext"))
		//ctx.Output.Body([]byte(body))
		//ctx.Output.Body([]byte("Router Param! *:" + ctx.Input.Param(":splat")))
		ctx.Output.Body([]byte("Router Param! id:" + ctx.Input.Param(":id")))
	})

	// namespace
	// 创建
	v1 := web.NewNamespace("/v1",
		// 条件限定 返回true 执行路由转发否则不执行
		web.NSCond(func(ctx *context.Context) bool {
			agent := ctx.Request.Header.Get("User-Agent")
			if strings.Contains(agent, "GO") {
				return true
			}
			return false
		}),

		// 前置过滤器
		web.NSBefore(func(ctx *context.Context) {
			// TODO 身份认证、鉴权、追踪、请求日志
		}, func(ctx *context.Context) {

		}),
		// 后置过滤器
		web.NSAfter(func(ctx *context.Context) {
			// TODO 资源回收、统计
		}, func(ctx *context.Context) {

		}),
		web.NSRouter("/content", &ContentController{}),
		web.NSNamespace("/tag",
			web.NSRouter("group", &TagGroupController{}),
		),
	)
	// 注册
	web.AddNamespace(v1)

	// 另一种路由命名空间的使用
	v2 := web.NewNamespace("/v2")
	v2.Cond(func(ctx *context.Context) bool {
		return true
	})
	v2.Filter("before", func(ctx *context.Context) {})
	v2.Filter("after", func(ctx *context.Context) {})
	v2.Router("/article", &ContentController{})
	v2Tag := web.NewNamespace("tag")
	v2.Namespace(v2Tag)
	web.AddNamespace(v2)

	test := web.NewNamespace("/test")
	test.Router("request-data/:id", &TestRequestController{})
	test.Router("request-data/other/?:key", &TestRequestController{}, "post:Other")
	web.AddNamespace(test)
}

type TestRequestController struct {
	web.Controller
}

func (c *TestRequestController) Post() {
	startAt, _ := c.GetInt("startAt")
	type Article struct {
		Subject string `json:"subject,omitempty"`
		Content string `json:"content,omitempty"`
	}
	article := &Article{}
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, article)
	requestData := map[string]interface{}{
		"ID":   c.Ctx.Input.Param(":id"),
		"name": c.GetString("name"),
		//"courses": c.GetString("courses"),
		"courses": c.GetStrings("courses"),
		"startAt": startAt,
		// body
		"keyword":       c.GetString("keyword"), // urlencoded get
		"content":       c.GetString("content"), // form-data
		"body":          &article,               // raw,json
		"Authorization": c.Ctx.Input.Header("Authorization"),
	}
	c.Data["json"] = requestData
	_ = c.ServeJSON()
}

func (c *TestRequestController) Other() {
	requestData := map[string]interface{}{}
	requestData["name"] = c.Ctx.Input.Param("name")

	c.Data["json"] = requestData
	c.ServeJSON()
}

type TagGroupController struct {
	web.Controller
}

func (this *TagGroupController) Get() {
	_ = this.Ctx.Output.Body([]byte("\"GO 海绵宝宝! TagGroupController Get()"))
}

type ContentController struct {
	web.Controller
}

func (this *ContentController) Get() {
	_ = this.Ctx.Output.Body([]byte("\"GO 海绵宝宝! Content Get()"))
}
func (this *ContentController) Post() {
	_ = this.Ctx.Output.Body([]byte("\"GO 海绵宝宝! Content Post()"))
}
func (this *ContentController) Select() {
	_ = this.Ctx.Output.Body([]byte("\"GO 海绵宝宝! Content Select()"))
}

type CPController struct {
	web.Controller
}

func (this *CPController) URLMapping() {
	this.Mapping("Create", this.Create)
	this.Mapping("Retrieve", this.Retrieve)
}

// @router /content [post,put]
func (this *CPController) Create() {
	this.Ctx.Output.Body([]byte("Go Create!"))
}

// @router /content/:id [get]
func (this *CPController) Retrieve() {
	this.Ctx.Output.Body([]byte("Go Retrieve!"))
}

type OurController struct {
	web.Controller
}

func (this *OurController) Get() {
	_ = this.Ctx.Output.Body([]byte("Go 海绵宝宝! Controller Method GET!"))
}

func (this *OurController) Post() {
	_ = this.Ctx.Output.Body([]byte("Go 海绵宝宝! Controller Method POST!"))
}
func (this *OurController) Put() {
	_ = this.Ctx.Output.Body([]byte("Go 海绵宝宝! Controller Method PUT!"))
}
func (this *OurController) Delete() {
	_ = this.Ctx.Output.Body([]byte("Go 海绵宝宝! Controller Method DELETE!"))
}
func (this *OurController) Options() {
	_ = this.Ctx.Output.Body([]byte("Go 海绵宝宝! Controller Method Options!"))
}
func (this *OurController) Head() {
	_ = this.Ctx.Output.Body([]byte("Go 海绵宝宝! Controller Method Head!"))
}
func (this *OurController) Patch() {
	_ = this.Ctx.Output.Body([]byte("Go 海绵宝宝! Controller Method Patch!"))
}
func (this *OurController) OtherFunc() {
	_ = this.Ctx.Output.Body([]byte("Go 海绵宝宝! Controller Method OtherFunc!"))
}
func (this *OurController) Register() {
	_ = this.Ctx.Output.Body([]byte("Go 海绵宝宝! Controller Method Register!"))
}

func NameHandler(ctx *context.Context) {
	_ = ctx.Output.Body([]byte("Go 海绵宝宝! Name Handler!"))
}
