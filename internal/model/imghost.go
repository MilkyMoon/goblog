package model

import (
	"goblog/config"
	"goblog/internal/common"
	"errors"
	"github.com/kataras/iris"
	"io"
	"os"
	"path"
	"path/filepath"
	"sort"
	"time"
)

type ImgHost interface {
	ImageUpload(iris.Context) (string,error)
	ImageList() Images
	ImageDel(string) bool
}

type ImageHost struct {

}

type Images []Image

type Image struct {
	Url string `json:"url"`
	//创建时间
	CreateTime time.Time `json:"create_time"`
}

func (list *ImageHost)ImageList() []Image {
	var images = Images{}
	filepath.Walk(common.GetImagePath(), func(path string, info os.FileInfo, err error) error {
		suffix := filepath.Ext(path)
		//跳过文件夹
		if info.IsDir() || (suffix != ".jpg" && suffix != ".jpeg" && suffix != ".png" && suffix != ".gif"){
			return nil
		}

		image := Image{}
		//只获取web文件夹之后的路径
		root := filepath.Join(config.Root,"web")
		image.Url = path[len(root):]
		image.CreateTime = info.ModTime()

		images = append(images,image)

		return nil
	})

	sort.Sort(&images)

	return images
}

func (image *ImageHost)ImageUpload(ctx iris.Context) (string,error) {
	file, info, err := ctx.FormFile("file")
	if err != nil {
		return "",errors.New("上传图片失败！")
	}
	defer file.Close()

	//生成文件名，当前纳秒转为36进制
	fname := common.BaseTo36(time.Now().UnixNano()) + path.Ext(info.Filename)

	//创建文件夹
	date := time.Now().Format("2006-01-02")
	path := filepath.Join(common.GetImagePath(),date)
	err = os.MkdirAll(path,os.ModePerm)
	if err != nil {
		panic("文件夹创建失败！")
	}

	out,err := os.OpenFile(path + "/" +fname,os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "",errors.New("保存图片失败！")
	}
	defer out.Close()
	_,err = io.Copy(out, file)

	if err != nil {
		return "",errors.New("保存图片失败！")
	}

	return "/" + config.Conf.Get("image_host.img_dir").(string) + "/" + fname,nil
}

//实现了sort接口
func (a Images) Len() int {
	return len(a)
}

func (a Images) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Images) Less(i, j int) bool {
	if a[i].Url > a[j].Url {
		return true
	}

	if a[i].CreateTime.After(a[j].CreateTime) {
		return true
	}

	return false
}