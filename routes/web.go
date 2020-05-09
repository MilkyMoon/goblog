package routes

import (
	"codwiki.cn/goblog/internal/controller"
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

	app := api.Party("/", crs).AllowMethods(iris.MethodOptions)

	// 首页模块
	app.Get("/", controller.Index)

}