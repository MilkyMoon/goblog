package config

import (
	"flag"
	"fmt"
	"github.com/pelletier/go-toml"
	"os"
	"path/filepath"
)

var (
	FlagPort     int //提供服务端口号
	FlagConfName string
	rootPath     string
	configFile   string
	debug        = true
	conf         *toml.Tree
)

func init() {
	FlagParse()

	env := "online"
	if FlagConfName != "" {
		env = FlagConfName
	}

	dir, _ := os.Getwd()
	rootPath = filepath.Join(dir,"..")
	configFile = fmt.Sprintf("%s/config/conf_%s.toml", rootPath, env)

	fmt.Println(configFile)
	_, err := os.Stat(configFile)
	if err != nil && os.IsNotExist(err) {
		panic(fmt.Sprintf("config file (%s) not exist", configFile))
	}
	config, err := toml.LoadFile(configFile)
	if err != nil {
		panic(fmt.Sprintf("load conf file(%s) error(%+v)", configFile, err))
	}
	fmt.Println("load config " + configFile)
	conf = config
}

//FlagParse 解析命令行参数
func FlagParse() {
	if !flag.Parsed() {
		flag.IntVar(&FlagPort, "port", 0, "the server port")
		flag.StringVar(&FlagConfName, "conf", "", "the server load conf")

		flag.Parse()
	}
}

func GetRootPath() string {
	return rootPath
}

func String(key string) string {
	return conf.Get(key).(string)
}

func Int(key string) int {
	return int(conf.Get(key).(int64))
}
