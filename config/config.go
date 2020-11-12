package config

import (
	"github.com/pelletier/go-toml"
	"os"
)

var Conf = new()

var Root = root()
/**
 * 返回单例实例
 * @method New
 */
func new() *toml.Tree {

	config, err := toml.LoadFile("../config/config.toml")

	if err != nil {
		panic(err.Error())
	}

	return config
}

//获取当前所执行文件的根目录
func root() string {
	//file, _ := exec.LookPath(os.Args[0])
	//
	//path, _ := filepath.Abs(file)
	//
	//i := strings.LastIndex(path, "/")
	//if i < 0 {
	//	i = strings.LastIndex(path, "\\")
	//}

	path,_ := os.Getwd()

	// 获取文件的绝对路径
	//path, _ := filepath.Abs(file)

	// 获取最后一个文件名所在的位置，如 /home/bro/blog ，
	// 这里可以获取到 blog 的 b 所在的位置，为了下面的 slice 截取字符串
	//index := strings.LastIndex(path, string(os.PathSeparator))

	//fmt.Println(path,index)

	return string(path)
}