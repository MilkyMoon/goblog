package model

import (
	"goblog/internal/common"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Navbars []Navbar

type Navbar struct {
	Title string `json:"title"`
	Path string `json:"path"`
	Folder string `json:"folder"`
	Description string `json:"description"`
	OutLink bool `json:"out_link"`
}

func GetNavList() Navbars {
	f,err := os.Open(filepath.Join(common.GetRootPath(),"navbar.json"))

	if err != nil{
		return nil
	}

	content,_ := ioutil.ReadAll(f)

	var nav Navbars

	_ = json.Unmarshal(content,&nav)

	return nav
}