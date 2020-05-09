package controller

import "github.com/kataras/iris"

func Index(ctx iris.Context)  {
	ctx.JSON(iris.Map{"message": "Hello Iris!"})
}