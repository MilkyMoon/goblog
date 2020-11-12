package middleware

import (
	"encoding/json"
	"github.com/kataras/iris"
	"goblog/config"
	"goblog/internal/model"
)

func SiteInfo(ctx iris.Context){
	ctx.ViewData("title", config.String("app.name"))
	ctx.ViewData("slogan", config.String("app.slogan"))
	ctx.ViewData("icp", config.String("app.icp"))
	ctx.Next()
}

//菜单栏
func NavList(ctx iris.Context) {
	data,_ := json.Marshal(model.CateModel.NavList())

	ctx.ViewData("menu",string(data))
	ctx.Next()
}

//分类菜单
func CategoryList(ctx iris.Context) {
	data,_ := json.Marshal(model.CateModel.CategoryList())

	ctx.ViewData("cate",string(data))
	ctx.Next()
}

func SliderInfo(ctx iris.Context)  {
	data,_ := json.Marshal(model.CateModel.GetSliderInfo())

	ctx.ViewData("slider",string(data))
	ctx.Next()
}

func DocsHandler(ctx iris.Context){
	ctx.Next()
}