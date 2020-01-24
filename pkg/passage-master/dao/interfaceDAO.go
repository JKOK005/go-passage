package dao

import "go-passage/pkg/passage-master/schema"

type DAOInterface interface {
	CreateServer(sModel *schema.ServerModel) error
	GetServer(pred map[string]string) (*schema.ServerModel, error)
	DeleteServer(sModel *schema.ServerModel) error
	CreateApp(aModel *schema.AppModel) error
	GetApp(pred map[string]string) (*schema.AppModel, error)
	GetApps(pred map[string]string) ([]*schema.AppModel, error)
	GetDeployedServers(aModel []*schema.AppModel) ([]*schema.ServerModel, error)
	DeleteApp(aModel *schema.AppModel) error
}
