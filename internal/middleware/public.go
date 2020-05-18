package middleware

import (
	"codwiki.cn/goblog/config"
	"codwiki.cn/goblog/internal/model"
	"encoding/json"
	"github.com/kataras/iris"
)

func SiteInfo(ctx iris.Context){
	ctx.ViewData("title",config.Conf.Get("app.name").(string))
	ctx.ViewData("slogan",config.Conf.Get("app.slogan").(string))
	ctx.ViewData("icp",config.Conf.Get("app.icp").(string))
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