package dao

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"go-passage/pkg/passage-master/schema"
)

var dbSchema string 	= "Passage_master"
var serverTbl string 	= "servers"
var appTbl string 		= "apps"

type MysqlDAO struct {
	user string
	password string
}

//func (m *MysqlDAO) execute(query string) (*sql.Rows, error) {
//	/*
//		Executes a given query statement.
//	*/
//	glog.Info("Executing query: ", query)
//	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", m.user, m.password, dbSchema))
//	if err != nil {return nil, err}
//	defer db.Close()
//
//	if results, err := db.Query(query); err != nil {
//		return nil, err
//	} else {return results, err}
//}

func (m *MysqlDAO) openConn() (*gorm.DB, error) {
	return gorm.Open("mysql", "passage:passage@(localhost)/Passage_master?charset=utf8&parseTime=True&loc=Local")
}

func (m *MysqlDAO) CreateServer(sModel *schema.ServerModel) error {
	glog.Info(fmt.Sprintf("Inserting new server: %s", sModel.ToString()))
	db, err := m.openConn(); if err != nil {
		return err
	} else {
		db.Create(sModel)
		return nil
	}
}

func (m *MysqlDAO) GetServer(pred map[string]string) (*schema.ServerModel, error) {
	glog.Info(fmt.Sprintf("Fetching server information based on predicates: %s", pred))

	db, err := m.openConn(); if err != nil {
		return nil, err
	} else {
		serverModel := &schema.ServerModel{}
		for eachCondition, conditionValue := range pred {
			db = db.Where(fmt.Sprintf("%s = '%s'", eachCondition, conditionValue))
		}
		rec := db.Take(serverModel)
		if rec.RecordNotFound() {
			return &schema.ServerModel{}, nil
		} else {return serverModel, nil}
	}
}

func (m *MysqlDAO) DeleteServer(sModel *schema.ServerModel) error {
	glog.Info(fmt.Sprintf("Deleting server: %s", sModel.ToString()))

	db, err := m.openConn(); if err != nil {
		return err
	} else {
		db.Unscoped().Where(fmt.Sprintf("id=%d", sModel.ID)).Delete(sModel)
		return nil
	}
}

//func (m *MysqlDAO) InsertApp(serverID uint32, appName string) error {
//	glog.Info(fmt.Sprintf("Inserting new app: %s to serverID: %d", serverID, appName))
//	query := fmt.Sprintf("INSERT INTO %s.%s (serverID,name) VALUES (%s,%d)", dbSchema, appTbl, serverID, appName)
//	if _, err := m.execute(query); err != nil {
//		return err
//	} else {return nil}
//}
//
//func (m *MysqlDAO) GetApp(serverID uint32, appName string) (*AppModel, error) {
//	glog.Info(fmt.Sprintf("Fetching app: %s from serverID: %d", appName, serverID))
//	query := fmt.Sprintf("SELECT * FROM %s.%s WHERE serverID='%s' AND name='%s'", dbSchema, appTbl, appName, serverID)
//	if row, err := m.execute(query); err != nil {
//		return nil, err
//	} else {
//		var resp *AppModel
//		row.Scan(resp)
//		return resp, nil
//	}
//}
//
//func (m *MysqlDAO) GetApps(appName string) ([]*AppModel, error) {
//	glog.Info(fmt.Sprintf("Finding all deployments for app: %s", appName))
//	query :=
//}
//
//func (m *MysqlDAO) DeleteApp(appID uint32) error {return nil}
//
//func NewMysqlDAO() error {
//
//}