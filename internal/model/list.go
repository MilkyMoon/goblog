package model

import (
	"sort"
)

var CateModel = newListCate()

type List interface {
	ArticleContent(string) Article
	CategoryList() Categories
	ArticleList(string) Articles
	NavList() Navbars
	GetSliderInfo() Slider
}

type ListCate struct {
	Navbars Navbars
	Categories Categories
	Articles Articles
	Slider Slider
}

type ListBook struct {
	Articles Articles
}

func newListCate() List {
	l := ListCate{
		Navbars : GetNavList(),
		Categories: GetCategoryList(),
		Articles:   nil,
		Slider: GetSliderInfo(),
	}

	return &l
}

func (list *ListCate) ArticleContent(path string) Article {
	content,err := GetArticleContent(path)

	if err != nil {
		panic(err)
	}

	return content
}

func (list *ListCate) ArticleList(name string) Articles {
	var articles Articles
	articles = GetArticleList(name)

	sort.Sort(&articles)
	return articles
}

func (list *ListCate) CategoryList() Categories {
	sort.Sort(&list.Categories)
	return list.Categories
}

func (list *ListCate) NavList() Navbars {
	return list.Navbars
}

func (list *ListCate) GetSliderInfo() Slider {
	return list.Slider
}

func (l ListBook) ArticleContent(string) Article {
	panic("implement me")
}

func (l ListBook) CategoryList() Categories {
	panic("implement me")
}

func (l ListBook) ArticleList(string) Articles {
	panic("implement me")
}
