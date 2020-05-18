package main

import (
	"codwiki.cn/Irisgo/config"
	"codwiki.cn/goblog/routes"
	"github.com/kataras/iris"
)

func main()  {
	app := iris.New()

	routes.Register(app)

	app.HandleDir("/static","web/static")

	addr := config.Conf.Get("app.addr").(string)
	_ = app.Run(iris.Addr(addr),iris.WithoutServerError())
}
