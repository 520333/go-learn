package content

import "github.com/beego/beego/v2/server/web/context"

func Retrieve(ctx *context.Context) {
	_ = ctx.Output.Body([]byte("Go 海绵宝宝! Content Retrieve Handler!"))
}
