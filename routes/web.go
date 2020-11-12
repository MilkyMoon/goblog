package routes

import (
	"goblog/config"
	"goblog/internal/controller"
	"goblog/internal/middleware"
	"github.com/kataras/iris"
)

func Register(api *iris.Application)  {
	//设置请求头，允许跨域
	//crs := cors.New(cors.Options{
	//	AllowedOrigins:   []string{"*"},
	//	AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
	//	AllowedHeaders:   []string{"*"},
	//	ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	//	AllowCredentials: true,
	//})

	HTML := iris.HTML("../web/views", ".html")
	//修改默认的{{}}标记符为{%%}，因为前端使用Vue，与Vue冲突
	HTML.Delims("{%","%}")
	HTML.Layout("layout.html")
	api.RegisterView(HTML)

	//注册全局中间件，捕获异常
	api.Use(middleware.Recover)
	api.OnAnyErrorCode(notFound)

	app := api.Party("/",middleware.SiteInfo,middleware.NavList,middleware.CategoryList,middleware.SliderInfo).AllowMethods(iris.MethodOptions)
	app.Get("/",controller.List)

	docs := app.Party("/")
	docs = docs.Party("/docs")
	docs.Get("/{path}",controller.List)
	docs.Get("/{path}/{name}",controller.Post)

	//图床路由组
	img := app.Party(config.String("image_host.path"))
	img.Get("/login",controller.ImgLoginView)
	img.Post("/login",controller.ImgLogin)
	img.Get("/index",controller.ImgIndex)
	img.Get("/list",controller.ImgList)
	img.Post("/upload",controller.ImgUpload)

	//跳转错误页面
	app.Get("/errors",notFound)
	app.Post("/webhooks",controller.Webhook)
	app.Get("/reload",controller.Reload)



	//book := app.Party("/book")
	//book.Get("/")
}

func notFound(ctx iris.Context)  {
	ctx.View("404.html")
}