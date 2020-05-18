package controller

import (
	"codwiki.cn/goblog/internal/common"
	"codwiki.cn/goblog/internal/model"
	"encoding/json"
	"github.com/kataras/iris"
	"net/url"
	"path/filepath"
)

func List(ctx iris.Context)  {
	res := model.CateModel.ArticleList(filepath.Join(common.GetDocsPath(), ctx.Params().Get("path")))

	data,_ := json.Marshal(res)

	ctx.ViewData("data",string(data))

	ctx.View("list.html")
}

func Post(ctx iris.Context)  {
	dir,_ := url.QueryUnescape(ctx.Params().Get("path"))
	name,_ := url.QueryUnescape(ctx.Params().Get("name"))
	path := filepath.Join(common.GetDocsPath(), dir,name)
	res := model.CateModel.ArticleContent(path + ".md")

	data,_ := json.Marshal(res)

	ctx.ViewData("data",string(data))

	ctx.View("marked_view.html")
}