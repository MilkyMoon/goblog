package common

import (
	"goblog/config"
	"path/filepath"
)

func GetRootPath() string {
	return filepath.Join(config.Root,config.Conf.Get("res.root_dir").(string))
}

//获取文章文件夹根目录
func GetDocsPath() string {
	return filepath.Join(config.Root,config.Conf.Get("res.docs_dir").(string))
}

//获取文档文件夹根目录
func GetBookPath() string  {
	return filepath.Join(config.Root,config.Conf.Get("res.books_dir").(string))
}

func CodeToText()  {
	
}

type Error struct {
	code  int
	msg   string
}

func (e *Error) Error() string {
	return e.msg
}