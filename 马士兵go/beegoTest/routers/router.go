package routers

import (
	"beegoTest/api/content"
	"beegoTest/controllers"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Get("/easy", func(ctx *context.Context) {
		_ = ctx.Output.Body([]byte("Go 海绵宝宝!"))
	})
	beego.Get("/named-handler", NameHandler)
	beego.Get("/content/", content.Retrieve)
}

func NameHandler(ctx *context.Context) {
	_ = ctx.Output.Body([]byte("Go 海绵宝宝! Name Handler!"))
}
