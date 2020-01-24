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

	// Separate migration of schemas done due to the need to create a foreign key for AppModel
	// Issue tracked at: https://github.com/jinzhu/gorm/issues/450
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&ServerModel{},
	)

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&AppModel{},
	).AddForeignKey("deployed_server_uid", "server_model(id)", "CASCADE", "RESTRICT")
}