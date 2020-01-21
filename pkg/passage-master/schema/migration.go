package schema

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

func StartMigration() {
	glog.Info("Initiating GORM migration")
	db, err := gorm.Open("mysql", "passage:passage@(localhost)/Passage_master?charset=utf8&parseTime=True&loc=Local")

	if err != nil {panic(err)}
	defer db.Close()

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&ServerModel{},
		&AppModel{},
	)
}