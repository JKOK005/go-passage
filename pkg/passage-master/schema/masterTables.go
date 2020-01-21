package schema

import (
	"github.com/jinzhu/gorm"
)

type ServerModel struct {
	gorm.Model
	Url string `gorm:"unique_index:indx; type:varchar(20); NOT NULL"`
	Port int `gorm:"unique_index:indx; NOT NULL"`
}

func (s ServerModel) TableName() string {
	return "server_model"
}

type AppModel struct {
	gorm.Model
	Name string `gorm:"type:varchar(20); NOT NULL; unique_index:indx"`
	Server ServerModel `gorm:"foreignkey:ID"`
}

func (a AppModel) TableName() string {
	return "app_model"
}