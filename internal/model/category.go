package model

import (
	"goblog/internal/common"
	"os"
	"path/filepath"
)

type Categories []Category

type Category struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Number int `json:"number"`
}

/**
获取分类列表（当前不支持嵌套文件夹计数）
 */
func GetCategoryList() Categories {
	list := Categories{}

	//循环遍历文章目录，获取目录结构
	err := filepath.Walk(common.GetDocsPath(), func(path string, info os.FileInfo, err error) error {
		//跳过根目录
		if path == common.GetDocsPath(){
			return nil
		}
		//计数
		if info.IsDir() {
			list = append(list,Category{Name:info.Name(),Path:info.Name(),Number:0})
		}

		//判断是否是markdown文档
		if filepath.Ext(path) == ".md" {
			list[len(list) - 1].Number++
		}

		return nil
	})

	if err != nil{
		panic(err)
	}

	return list
}

// Len 实现 Sort 的接口
func (cate Categories) Len() int {
	return len(cate)
}

// Swap 实现的 Sort 接口
func (cate Categories) Swap(i, j int) {
	cate[i], cate[j] = cate[j], cate[i]
}

// Less 实现的 Sort 接口， 按照标签数量排序
func (cate Categories) Less(i, j int) bool {
	return cate[i].Number > cate[j].Number
}