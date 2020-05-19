package middleware

import (
	"github.com/kataras/iris"
)

//捕获panic异常
func Recover(ctx iris.Context)  {
	defer func() {
		if err := recover(); err != nil {
			ctx.Redirect("/errors",301)
			return
		}
	}()

	ctx.Next()
}
