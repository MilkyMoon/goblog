package main

import (
	"fmt"
	"goblog/config"
	"goblog/routes"
	"github.com/kataras/iris"
	"os"
)

func main()  {
	port := config.FlagPort
	if port == 0 {
		port = config.Int("app.port")
		if port == 0 {
			fmt.Println("port must be given")
			os.Exit(1)
		}
	}
	addr := fmt.Sprintf(":%d", port)

	app := iris.New()
	routes.Register(app)

	_ = app.Run(iris.Addr(addr),iris.WithoutServerError())
}
