package model

import (
	"goblog/internal/common"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Link struct {
	Title string `json:"title"`
	Url string `json:"url"`
	Icon string `json:"icon"`
	Description string `json:"description"`
}

type Slider struct {
	Author string `json:"author"`
	Avatar string `json:"avatar"`
	Description string `json:"description"`
	Email string `json:"email"`
	Github string `json:"github"`
	OutLink []Link `json:"out_link"`
}

func GetSliderInfo() Slider {
	f,err := os.Open(filepath.Join(common.GetResRootPath(),"app.json"))

	if err != nil{
		return Slider{}
	}

	content,_ := ioutil.ReadAll(f)

	var slider Slider

	_ = json.Unmarshal(content,&slider)

	return slider
}
