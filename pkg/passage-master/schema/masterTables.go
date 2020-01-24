package schema

import (
	"encoding/json"
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

func (s ServerModel) ToString() string {
	str, _ := json.Marshal(s)
	return string(str)
}

type AppModel struct {
	gorm.Model
	Name string `gorm:"type:varchar(20); NOT NULL; unique_index:indx"`
	DeployedServerUid uint `gorm:"NOT NULL; unique_index:indx"`
}

func (a AppModel) TableName() string {
	return "app_model"
}

func (a AppModel) ToString() string {
	str, _ := json.Marshal(a)
	return string(str)
}