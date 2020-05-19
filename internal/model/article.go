package model

import (
	"bufio"
	"codwiki.cn/goblog/internal/common"
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
	"time"
)

type Articles []Article

type Article struct {
	//文章标题
	Title string `json:"title"`
	// 文章标签
	Category string `json:"category"`
	//创建时间
	CreateTime time.Time `json:"create_time"`
	//更新时间
	UpdateTime time.Time `json:"update_time"`
	//文章描述
	Description string `json:"description"`
	//文章内容
	Body string `json:"body"`
	//文章所在路径
	Path string `json:"path"`
	//文章类型
	Type int `json:"type"`
}

func GetArticleList(name string) []Article {
	list := make([]Article, 0, 10)

	//循环遍历文章目录，获取文章
	err := filepath.Walk(name, func(path string, info os.FileInfo, err error) error {
		//判断是否是markdown文档
		if filepath.Ext(path) == ".md" {
			//获取文档信息
			article,err := GetArticleContent(path)

			if err != nil {
				return nil
			}

			//获取文件后缀
			//suffix := filepath.Ext(path)
			//获取文件名
			//article.Title = strings.TrimSuffix(filepath.Base(path),suffix)
			article.Body = ""

			list = append(list,article)
		}

		return nil
	})

	if err != nil{
		panic(err)
	}

	return list
}

/**
获取文章详情
 */
func GetArticleContent(path string) (Article,error) {
	article := Article{}

	//以只读方式打开文档
	if fileObj,err := os.Open(path);err == nil {
		defer fileObj.Close()

		//判断文章是否存在标题
		buf := bufio.NewReader(fileObj)

		var content string

		is_begin := false

		for {
			//逐行读取
			line, err := buf.ReadString('\n')

			//找到不为空的第一行文字
			if line != "" && !is_begin{
				is_begin = true

				//判断是否为一二三级标题
				//去除两端多余空格
				str := strings.TrimSpace(line)
				reg := "^(#{1,3})\\s+(.*)"

				ok,err :=regexp.MatchString(reg,str)

				if err != nil || !ok{
					//不为标题时加入文章详情
					content += line
					continue
				}

				fly := regexp.MustCompile("^(#{1,3})\\s+")
				titles := fly.FindStringSubmatch(str)

				//获取标题
				article.Title = strings.TrimSpace(string(str)[len(titles[1]):])
			} else {
				content += line
			}

			//遇到任何错误立即结束
			if err != nil {
				break
			}
		}

		//如果文章内容中不存在标题，就使用文件名作为标题
		if article.Title == "" {
			//获取文件后缀
			suffix := filepath.Ext(path)
			//获取文件名
			article.Title = strings.TrimSuffix(filepath.Base(path),suffix)
		}

		//截取文章内容
		body := strings.Split(content,"<br description/>")

		//判断文章内容的结构
		if len(body) != 3 && len(body) != 1 {
			return Article{},errors.New("文章内容结构不正确")
		}

		//截取文章描述
		if len(body) == 3 {
			article.Description = body[1]
			article.Body = body[1] + body[2]
		} else {
			article.Body = body[0]
		}

		//获取文件最后更新时间
		fi,err := fileObj.Stat()
		if err != nil {
			article.UpdateTime = time.Now()
		} else {
			article.UpdateTime = fi.ModTime()
		}

		//获取创建时间

		//linux下获取创建时间（此方法无法获取到创建时间，后续更新shell脚本解决方案）
		//stat_t := fi.Sys().(*syscall.Stat_t)
		//ctim := stat_t.Ctim

		//windows下获取创建时间
		//wFileSys := fi.Sys().(*syscall.Win32FileAttributeData)
		//ctim := wFileSys.CreationTime.Nanoseconds()/1e9

		//mac下获取文件创建时间
		stat_t := fi.Sys().(*syscall.Stat_t)
		ctim := stat_t.Birthtimespec

		article.CreateTime = time.Unix(int64(ctim.Sec), int64(ctim.Nsec))

		//获取Tag标签（父文件夹）
		dirs := strings.Split(path,"/")
		article.Category = dirs[len(dirs) - 2]

		//构建文件访问路径
		//判断文件所在目录
		//var path_arr []string
		if(strings.Contains(path,common.GetDocsPath())){
			//path_arr = strings.Split(string([]byte(path)[len(common.GetDocsPath()) + 1:len(path) - 3]),"/")
			article.Path = string([]byte(path)[len(common.GetDocsPath()):len(path) - 3])
			article.Type = 1
		} else {
			//path_arr = strings.Split(string([]byte(path)[len(common.GetBookPath()) + 1:len(path) - 3]),"/")
			article.Path = string([]byte(path)[len(common.GetBookPath()):len(path) - 3])
			article.Type = 2
		}

		//path_str := ""
		//for _,v := range path_arr {
		//	//逐个字转换
		//	for _,word := range []rune(v) {
		//		fmt.Println(string(word))
		//	}
		//	path_str += "/"
		//	ascii := strconv.QuoteToASCII(v)
		//	ascii = strings.Replace(ascii,"\\","",-1)
		//	path_str += ascii[1:len(ascii) - 2]
		//}

		//article.Path = path_str
	} else {
		return article,err
	}

	return article,nil
}

//实现了sort接口
func (a Articles) Len() int {
	return len(a)
}

func (a Articles) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Articles) Less(i, j int) bool {
	if a[i].UpdateTime.After(a[j].UpdateTime) {
		return true
	}

	if a[i].CreateTime.After(a[j].CreateTime) {
		return true
	}

	return false
}