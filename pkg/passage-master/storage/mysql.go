package storage

import (
	_ "github.com/go-sql-driver/mysql"
)

type MysqlClient struct {
	user string
	password string
	schema string
}

type ServerClient struct {
	exposeUrl string
	exposePort uint32
}

func (m *MysqlClient) AddServer(serverUrl string, serverPort uint32) error {return nil}
func (m *MysqlClient) GetServer(serverUrl string, serverPort uint32) (uint32, error) {return -1, nil}
func (m *MysqlClient) RemoveServer(serverID uint32) error {return nil}
func (m *MysqlClient) AddApp(serverID uint32, appName string) error {return nil}
func (m *MysqlClient) GetApp(serverID uint32, appName string) (uint32, error) {return -1, nil}
func (m *MysqlClient) GetApps(appName string) ([]ServerClient, error) {return nil, nil}
func (m *MysqlClient) RemoveApp(appID uint32) error {return nil}

func CreateClient() (*MysqlClient, error) {
	/*
		Factory dispatcher pattern for new ZK client
	*/
	return nil, nil
}