package models

type Article struct {
	Id int
	Title string
	CateId int
	State int
	//ArticleCate ArticleCate `gorm:"foreignKey:CateId"` //重写外键
}

func (article Article) TableName() string {
	return "article"
}
