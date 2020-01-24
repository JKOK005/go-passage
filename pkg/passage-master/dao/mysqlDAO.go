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

func (m *MysqlDAO) CreateApp(aModel *schema.AppModel) error {
	glog.Info(fmt.Sprintf("Inserting new app: %s", aModel.ToString()))
	db, err := m.openConn(); if err != nil {
		return err
	} else {
		db.Create(aModel)
		return nil
	}
}

func (m *MysqlDAO) GetApp(pred map[string]string) (*schema.AppModel, error) {
	glog.Info(fmt.Sprintf("Fetching app based on predicates: %s", pred))
	db, err := m.openConn(); if err != nil {
		return nil, err
	} else {
		appModel := &schema.AppModel{}
		for eachCondition, conditionValue := range pred {
			db = db.Where(fmt.Sprintf("%s = '%s'", eachCondition, conditionValue))
		}
		rec := db.Take(appModel)
		if rec.RecordNotFound() {
			return &schema.AppModel{}, nil
		} else {return appModel, nil}
	}
}

func (m *MysqlDAO) GetApps(pred map[string]string) ([]*schema.AppModel, error) {
	glog.Info(fmt.Sprintf("Finding all deployments for app based on predicates: %s", pred))
	db, err := m.openConn(); if err != nil {
		return nil, err
	} else {
		appModel := make([]*schema.AppModel, 1)
		for eachCondition, conditionValue := range pred {
			db = db.Where(fmt.Sprintf("%s = '%s'", eachCondition, conditionValue))
		}
		rec := db.Find(&appModel)
		if rec.RecordNotFound() {
			return make([]*schema.AppModel, 0), nil
		} else {return appModel, nil}
	}
}

func (m *MysqlDAO) GetDeployedServers(aModel []*schema.AppModel) ([]*schema.ServerModel, error) {
	glog.Info(fmt.Sprintf("Finding all deployments of apps"))
	db, err := m.openConn(); if err != nil {
		return nil, err
	} else {
		srvModel := make([]*schema.ServerModel, 1)
		for indx, eachApp := range aModel {
			if indx == 0 {
				db = db.Where(fmt.Sprintf("%s = '%d'", "id", eachApp.DeployedServerUid))
			} else {
				db = db.Or(fmt.Sprintf("%s = '%d'", "id", eachApp.DeployedServerUid))
			}
		}
		rec := db.Find(&srvModel)
		if rec.RecordNotFound() {
			return make([]*schema.ServerModel, 0), nil
		} else {return srvModel, nil}
	}
}

func (m *MysqlDAO) DeleteApp(aModel *schema.AppModel) error {
	glog.Info(fmt.Sprintf("Deleting app: %s", aModel.ToString()))
	db, err := m.openConn(); if err != nil {
		return err
	} else {
		db.Unscoped().Where(fmt.Sprintf("id=%d", aModel.ID)).Delete(aModel)
		return nil
	}
}