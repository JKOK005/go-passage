package storage

import (
	"fmt"
	"github.com/golang/glog"
	_ "github.com/go-sql-driver/mysql"
	"go-passage/pkg/passage-master/dao"
)

type AccessHandler struct {
	modelDAO dao.DAOInterface
}

func (m *AccessHandler) RegisterServer(serverUrl string, serverPort uint32) error {
	glog.Info(fmt.Sprintf("Registering server: %s:%d", serverUrl, serverPort))
	return m.modelDAO.InsertServer(serverUrl, serverPort)
}

func (m *AccessHandler) GetServerID(serverUrl string, serverPort uint32) (uint32, error) {return -1, nil}
func (m *AccessHandler) RemoveServer(serverID uint32) error {return nil}
func (m *AccessHandler) RegisterApp(serverID uint32, appName string) error {return nil}
func (m *AccessHandler) GetApps(appName string) ([]ServerClient, error) {return nil, nil}
func (m *AccessHandler) RemoveApp(appID uint32) error {return nil}

func CreateClient() (*MysqlHandler, error) {
	/*
		Factory dispatcher pattern for new ZK client
	*/
	return nil, nil
}