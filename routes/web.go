package routes

import (
	"codwiki.cn/goblog/internal/controller"
	"codwiki.cn/goblog/internal/middleware"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

func Register(api *iris.Application)  {
	//设置请求头，允许跨域
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})

	HTML := iris.HTML("./web/views", ".html")
	//修改默认的{{}}标记符为{%%}，因为前端使用Vue，与Vue冲突
	HTML.Delims("{%","%}")
	HTML.Layout("layout.html")
	api.RegisterView(HTML)

	app := api.Party("/", crs,middleware.SiteInfo).AllowMethods(iris.MethodOptions)
	// 首页模块

	app.Get("/", controller.List)

	docs := app.Party("/",middleware.NavList,middleware.CategoryList,middleware.SliderInfo)
	docs.Get("/",controller.List)
	docs = docs.Party("/docs")
	docs.Get("/{path}",controller.List)
	docs.Get("/{path}/{name}",controller.Post)

	//book := app.Party("/book")
	//book.Get("/")
}