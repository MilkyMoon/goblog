package controller

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"github.com/kataras/iris"
	"goblog/config"
	"goblog/internal/common"
	"goblog/internal/model"
	"io/ioutil"
	"net/url"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"unsafe"
)

func List(ctx iris.Context)  {
	res := model.CateModel.ArticleList(filepath.Join(common.GetDocsPath(), ctx.Params().Get("path")))

	//分页
	page := ctx.URLParam("page")
	var page_int int
	page_int,err := strconv.Atoi(page)

	if page == "" || err != nil {
		page_int = 1
	}

	var limit int
	limit_64 := config.Int("app.limit")
	limit =  *(*int)(unsafe.Pointer(&limit_64))

	length := len(res)

	if length > limit {
		begin  := (page_int - 1) * limit
		end    := (page_int - 1) * limit + limit

		if begin > (length - 1) {
			begin = length - 1
			end   = length - 1
		}

		if end > (length - 1) {
			end = length
		}

		res = res[begin:end]
	}

	data,_ := json.Marshal(res)

	ctx.ViewData("data",string(data))
	ctx.ViewData("page",page_int)
	ctx.ViewData("records",length)
	ctx.ViewData("perPage",limit)

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

func Webhook(ctx iris.Context){
	singn := ctx.GetHeader("X-Hub-Signature")
	body, err := ioutil.ReadAll(ctx.Request().Body)

	if err != nil {
		return
	}
	check := checkSecret(singn, body)
	if !check {
		return
	}

	gitpull()

	//刷新内存中的数据
	model.CateModel.Reload()

	ctx.StatusCode(200)
}

func gitpull() {
	path := common.GetRootPath()
	// 执行 git pull
	cmd := exec.Command("git", "pull")
	// 切换到命令要执行的目录
	cmd.Dir = path

	// 执行，并返回结果
	_, err := cmd.Output()

	if err != nil {
		panic(err)
	}
}

//验证签名
func checkSecret(singn string, body []byte) bool {
	if len(singn) != 45 || !strings.HasPrefix(singn, "sha1=") {
		return false
	}

	// 获取github配置的加密串
	secret := []byte(config.String("app.secret"))

	// 计算签名
	mac := hmac.New(sha1.New, secret)
	mac.Write(body)
	key := mac.Sum(nil)

	// Hex 解码
	singnature := make([]byte, 20)
	hex.Decode(singnature, []byte(singn[5:]))

	// 比较签名是否一直
	if hmac.Equal(singnature, key) {
		return true
	}

	return false
}

func Reload(ctx iris.Context)  {
	model.CateModel.Reload()
	ctx.StatusCode(200)
}