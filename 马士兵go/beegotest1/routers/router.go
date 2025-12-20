// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beegotest1/controllers"
	"encoding/json"
	"log"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)

	test := beego.NewNamespace("/test")
	test.Router("request-data/:id", &TestRequestController{})
	test.Router("request-data/other/?:key", &TestRequestController{}, "post:Other")
	test.Router("request-data/upload", &TestRequestController{}, "post:Upload")
	test.Router("request-data/cookie", &TestRequestController{}, "post:Cookie")
	test.Router("request-data/header", &TestRequestController{}, "get:Header")
	test.Router("response-data", &TestRequestController{}, "get:Resp")
	beego.AddNamespace(test)
	//beego.Post("/test/post", func(ctx *context.Context) {
	//	ctx.Input.Query("name")
	//})
}

type TestRequestController struct {
	beego.Controller
}
type Resp struct {
	Code    int    `json:"code" xml:"code"`
	Message string `json:"message" xml:"message"`
	Data    any    `json:"data" xml:"data"`
}
type Article struct {
	ID        uint   `json:"id,omitempty"`
	Subject   string `json:"subject,omitempty"`
	Views     int    `json:"views,omitempty"`
	Published bool   `json:"published,omitempty"`
}

func (c *TestRequestController) Resp() {
	// 响应数据
	//data := Resp{
	//	Code:    0,
	//	Message: "success",
	//	Data: Article{
	//		ID:      1,
	//		Subject: "Beego 一个功能齐全的web框架",
	//		Views:   1024,
	//		//Published: false,
	//	},
	//}
	//data1 := map[string]any{
	//	"message": "success",
	//}
	// 设置不同的响应格式
	//c.Data["json"] = data // json
	//_ = c.ServeJSON()
	//c.Data["xml"] = data  // xml
	//c.Data["xml"] = data1 // xml
	//_ = c.ServeXML()
	//c.Data["yaml"] = data // yaml
	//_ = c.ServeYAML()

	// 基于请求头Accept 完成响应格式的转换
	//_ = c.Ctx.Output.ServeFormatted(data, false, false)

	//c.Data["jsonp"] = data
	//_ = c.ServeJSONP() //get /response-data?callback=func

	//c.Ctx.WriteString("go")
	//c.Ctx.WriteString("lang")
	//c.Ctx.WriteString("海绵宝宝")

	// 文件下载
	c.Ctx.Output.Download("./beegotest1.exe", "test.exe")
}
func (c *TestRequestController) Header() {
	// 获取header
	value1 := c.Ctx.Input.Header("Content-Type")
	value2 := c.Ctx.Input.Header("Accept")
	value3 := c.Ctx.Input.Header("Authorization")
	c.Data["json"] = map[string]interface{}{
		"Request:Content-Type":  value1,
		"Request:Accept":        value2,
		"Request:Authorization": value3,
	}
	c.Ctx.Output.Header("X-Powered-By", "golang")
	_ = c.ServeJSON()
}
func (c *TestRequestController) Cookie() {
	// 设置cookie
	c.Ctx.Output.Cookie("token", "some token value")
	// 获取cookie
	token := c.Ctx.Input.Cookie("token")
	// 安全cookie
	secret := "your secret key"
	// 设置安全cookie
	//c.Ctx.SetSecureCookie(secret,"a","a")
	c.Ctx.SetSecureCookie(secret, "user", "12")
	// 获取安全cookie
	value, ok := c.GetSecureCookie(secret, "user")
	if !ok {
		log.Println("cookie not found")
	}

	c.Data["json"] = map[string]interface{}{
		"code":        200,
		"msg":         "success",
		"data":        token,
		"cookie:user": value,
	}
	_ = c.ServeJSON()
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
	requestData[":key"] = c.GetString(":key")
	requestData["name"] = c.GetString("name")
	ids := []uint{}
	if err := c.Ctx.Input.Bind(&ids, "ids"); err != nil {
		log.Println(err)
	}
	requestData["ids"] = ids
	type User struct {
		Name   string
		Status bool
	}
	user := &User{}
	if err := c.Ctx.Input.Bind(user, "user"); err != nil {
		log.Println(err)
	}
	requestData["user"] = user
	type Article struct {
		Subject   string `form:"subject" json:"subject"`
		Content   string `form:"content" json:"content"`
		Published bool   `form:"published" json:"published"`
		Views     int    `form:"views" json:"views"`
	}
	article := &Article{}
	_ = c.ParseForm(article)
	requestData["article"] = article

	c.Data["json"] = requestData
	_ = c.ServeJSON()
}
func (c *TestRequestController) Upload() {
	// 1.获取文件信息
	f, h, err := c.GetFile("logo")
	if err != nil {
		log.Fatalln(err)
	}

	// 2.校验文件信息是否满足
	// 2.1
	var maxSize int64 = 100 * 1024 // 100k
	if h.Size > maxSize {
		//log.Fatalln("file size too large")
		c.Data["json"] = map[string]any{
			"message":  "file size too large",
			"size":     h.Size,
			"Type":     h.Header.Get("Content-Type"),
			"filename": h.Filename,
		}
		_ = c.ServeJSON()
	}
	// 2.1.2 类型校验
	allowTypes := []string{"image/jpeg", "image/png", "image/gif", "image/webp"}
	allow := false
	for _, t := range allowTypes {
		if h.Header.Get("Content-Type") == t {
			allow = true
			break
		}
	}
	buffer := make([]byte, 512)
	f.Read(buffer)
	contentType := http.DetectContentType(buffer)
	// 以文件到N个字节为判断依据
	allowServer := false // server自己判定类型结果
	for _, t := range allowTypes {
		if contentType == t {
			allowServer = true
			break
		}
	}
	if !(allow && allowServer) {
		c.Data["json"] = map[string]any{
			"message":     "file type not allowed",
			"size":        h.Size,
			"Type":        h.Header.Get("Content-Type"),
			"Server Type": contentType,
			"filename":    h.Filename,
		}
		_ = c.ServeJSON()
	}
	// 2.2 规范存储 hash文件名防止文件名中出现乱码 保证相同文件存储一次

	// 3.存储到合理位置
	uploadPath := "./volume/upload/"
	if err := c.SaveToFile("logo", uploadPath+h.Filename); err != nil {
		log.Fatalln(err)
	}
	c.Data["json"] = map[string]any{
		"Size":     h.Size,
		"Type":     h.Header.Get("Content-Type"),
		"filename": h.Filename,
	}
	_ = c.ServeJSON()

}
