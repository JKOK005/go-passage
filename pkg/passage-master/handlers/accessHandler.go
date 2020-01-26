package handlers

import (
	"fmt"
	"strconv"
	"github.com/golang/glog"
	_ "github.com/go-sql-driver/mysql"
	"go-passage/pkg/passage-master/dao"
	"go-passage/pkg/passage-master/schema"
)

type AccessHandler struct {
	ModelDAO dao.DAOInterface
}

func (m AccessHandler) RegisterServer(serverUrl string, serverPort int) error {
	glog.Info(fmt.Sprintf("Register server: %s:%d", serverUrl, serverPort))
	srv := &schema.ServerModel{
		Url:   serverUrl,
		Port:  serverPort,
	}
	return m.ModelDAO.CreateServer(srv)
}

func (m AccessHandler) SearchServer(serverUrl string, serverPort int) (*schema.ServerModel, error) {
	glog.Info(fmt.Sprintf("Searching for server: %s:%d", serverUrl, serverPort))
	return m.ModelDAO.GetServer(map[string]string{"url": serverUrl, "port": strconv.Itoa(serverPort)})
}

func (m AccessHandler) RegisterApp(serverID int, appName string) error {
	glog.Info(fmt.Sprintf("Register app: %s under server id: %d", appName, serverID))
	app := &schema.AppModel{
		Name: appName,
		DeployedServerUid: serverID,
	}
	return m.ModelDAO.CreateApp(app)
}

func (m AccessHandler) SearchApp(appName string) ([]*schema.ServerModel, error) {
	/*
		Search for App location across servers
	*/
	glog.Info(fmt.Sprintf("Searching for deployments of app: %s", appName))
	apps, _ := m.ModelDAO.GetApps(map[string]string{"name" : appName})
	return m.ModelDAO.GetDeployedServers(apps)
}

func (m AccessHandler) DeleteApp(serverID int, appName string) error {
	glog.Info(fmt.Sprintf("Deleting app: %s under server id: %d", appName, serverID))
	app, _ := m.ModelDAO.GetApp(map[string]string{"name": appName, "deployed_server_uid": strconv.Itoa(serverID)})
	return m.ModelDAO.DeleteApp(app)
}

func (m AccessHandler) DeleteServer(serverID int) error {
	glog.Info(fmt.Sprintf("Deleting server id: %d", serverID))
	srv, _ := m.ModelDAO.GetServer(map[string]string{"id": strconv.Itoa(serverID)})
	return m.ModelDAO.DeleteServer(srv)
}