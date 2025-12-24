package serviceListen

import (
	"log"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	//s := g.Server()
	//s.BindHandler("/", func(r *ghttp.Request) {
	//	r.Response.Writeln("hello world111")
	//})
	//s.SetAddr(":8080")
	//s.SetPort(80, 81, 82) // 多端口
	//s.Run()
	//MultiService()
}

// MultiService 多个服务
func MultiService() {
	s1 := g.Server("one")
	s1.SetPort(81)
	s1.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writeln("Service one")
	})
	if err := s1.Start(); err != nil {
		log.Println("service one err:", err)
	}

	s2 := g.Server("two")
	s2.SetPort(82)
	s2.SetServerRoot("./static/")
	s2.SetIndexFolder(true) // 列出文件聊吧
	//s2.BindHandler("/", func(r *ghttp.Request) {
	//	r.Response.Writeln("Service two")
	//})
	if err := s2.Start(); err != nil {
		log.Println("service one err:", err)
	}
	// 阻塞等待
	g.Wait()
}
