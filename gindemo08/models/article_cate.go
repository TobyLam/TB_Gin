package models

type ArticleCate struct {
	Id int
	Title string
	State int
	Article []Article `gorm:"foreignKey:CateId"`
}

func (article_cate ArticleCate) TableName() string{
	return "article_cate"
}
