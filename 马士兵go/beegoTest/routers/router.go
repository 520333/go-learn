package routers

import (
	"beegoTest/api/content"
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
	web.Get("/content/", content.Retrieve)
	web.Router("/our", &OurController{})
	// post方法映射到OtherFunc
	ourController := &OurController{}
	web.Router("/our/other", ourController, "get,post:OtherFunc;*:Register")
}

func NameHandler(ctx *context.Context) {
	_ = ctx.Output.Body([]byte("Go 海绵宝宝! Name Handler!"))
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
