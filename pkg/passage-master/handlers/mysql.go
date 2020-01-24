package storage

import (
	"fmt"
	"github.com/golang/glog"
	_ "github.com/go-sql-driver/mysql"
	"go-passage/pkg/passage-master/dao"
	"go-passage/pkg/passage-master/schema"
)

type AccessHandler struct {
	modelDAO dao.DAOInterface
}

func (m *AccessHandler) RegisterServer(serverUrl string, serverPort int) error {
	glog.Info(fmt.Sprintf("Register server: %s:%d", serverUrl, serverPort))
	srv := &schema.ServerModel{
		Url:   serverUrl,
		Port:  serverPort,
	}
	return m.modelDAO.CreateServer(srv)
}

func (m *AccessHandler) RegisterApp(serverID int, appName string) error {
	glog.Info(fmt.Sprintf("Register app: %s under server id: %d", appName, serverID))
	app := &schema.AppModel{
		Name: appName,
		DeployedServerUid: serverID,
	}
	return m.modelDAO.CreateApp(app)
}

func (m *AccessHandler) SearchApp(appName string) ([]*schema.ServerModel, error) {
	/*
		Search for App location across servers
	*/
	glog.Info(fmt.Sprintf("Searching for deployments of app: %s", appName))
	apps, _ := m.modelDAO.GetApps(map[string]string{"name" : appName})
	return m.modelDAO.GetDeployedServers(apps)
}

func (m *AccessHandler) DeleteApp(serverID int, appName string) error {
	glog.Info(fmt.Sprintf("Deleting app: %s under server id: %d", appName, serverID))
	app, _ := m.modelDAO.GetApp(map[string]string{"Name": appName, "DeployedServerUid": string(serverID)})
	return m.modelDAO.DeleteApp(app)
}

func (m *AccessHandler) DeleteServer(serverID int) error {
	glog.Info(fmt.Sprintf("Deleting server id: %d", serverID))
	srv, _ := m.modelDAO.GetServer(map[string]string{"id": string(serverID)})
	return m.modelDAO.DeleteServer(srv)
}

func CreateClient() (*AccessHandler, error) {
	/*
		Factory dispatcher pattern for new ZK client
	*/
	return nil, nil
}