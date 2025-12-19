package routers

import (
	"beegoTest/controllers"

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
}

type ContentController struct {
	web.Controller
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
