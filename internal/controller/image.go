package controller

import (
	"goblog/config"
	"goblog/internal/model"
	"github.com/kataras/iris"
)

func ImgLoginView(ctx iris.Context)  {
	ctx.ViewData("path",config.Conf.Get("image_host.path").(string) + "/login")
	ctx.View("login.html")
}

//图床工具登录
func ImgLogin(ctx iris.Context)  {
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")

	if username != config.Conf.Get("image_host.username").(string) || password != config.Conf.Get("image_host.password").(string) {
		ctx.JSON(iris.Map{"code":"400","message": "用户名或密码错误！"})
	}

	//记录session
	ctx.JSON(iris.Map{"code":"200","data":config.Conf.Get("image_host.path").(string) + "/index","message": "登录成功！"})
}

func ImgIndex(ctx iris.Context)  {
	ctx.ViewData("path",config.Conf.Get("image_host.path").(string))
	ctx.View("img_index.html")
}

/**
图片列表
 */
func ImgList(ctx iris.Context)  {
	image := model.ImageHost{}
	res := image.ImageList()
	//data,_ := json.Marshal(res)

	ctx.JSON(iris.Map{"code":200,"data":res})
}

func ImgDel(ctx iris.Context)  {

}

func ImgUpload(ctx iris.Context)  {
	image := model.ImageHost{}
	path,err := image.ImageUpload(ctx)

	if err != nil {
		ctx.JSON(iris.Map{"code":500,"data":"","msg":err})
	}

	ctx.JSON(iris.Map{"code":200,"data":path})
}
