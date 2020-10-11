package salestypes

import (
	"strings"

	"github.com/jinzhu/gorm"
)

type SalesTypes struct {
	ID         int    `gorm:"id",json:"id,omitempty"`
	Title      string `gorm:"title",json:"title,omitempty"`
	Cover      string `gorm:"cover",json:"cover,omitempty"`
	Aftersale  int    `gorm:"aftersale",json:"aftersale,omitempty"`
	Beforesale int    `gorm:"beforesale",json:"beforesale,omitempty"`
	Numsperson int    `gorm:"numsperson",json:"numsperson,omitempty"`
	Urls       string `gorm:"urls",json:"urls,omitempty"`
	Type       int    `gorm:"type",json:"type,omitempty"`
}

//gorm框架与数据库表的映射主要这个方法返回的值
func (u *SalesTypes) TableName() string {
	return "salestypes"
}

func (u *SalesTypes) AllTypes(db *gorm.DB, salestype int) []*SalesTypes {
	sales := []*SalesTypes{}
	db.Where("type = ?", salestype).Find(&sales)
	return sales
}

func (u *SalesTypes) SalesById(db *gorm.DB, id int) *SalesTypes {
	sales := SalesTypes{}
	db.Where("id = ?", id).First(&sales)
	return &sales
}

func (u *SalesTypes) InsertSale(db *gorm.DB) bool {
	ds := []string{"8", "9", "10"}
	var animal = SalesTypes{Title: "qwerqwerwqe", Cover: "qwerqwerwer", Aftersale: 99, Beforesale: 99, Numsperson: 99, Urls: strings.Join(ds, ",")}
	_ = db.Create(&animal)
	ret := db.NewRecord(animal)
	return ret
}
