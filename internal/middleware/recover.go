package middleware

import (
	"fmt"
	"github.com/kataras/iris"
)

//捕获panic异常
func Recover(ctx iris.Context)  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			ctx.Redirect("/errors",301)
			return
		}
	}()

	ctx.Next()
}
