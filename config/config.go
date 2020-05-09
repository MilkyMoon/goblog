package config

import "github.com/pelletier/go-toml"

var Conf = New()
/**
 * 返回单例实例
 * @method New
 */
func New() *toml.Tree {

	config, err := toml.LoadFile("./config/config.toml")

	if err != nil {
		panic(err.Error())
	}

	return config
}