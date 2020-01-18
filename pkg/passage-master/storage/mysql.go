package storage

import (
	"database/sql"
	"fmt"
	"github.com/golang/glog"
	_ "github.com/go-sql-driver/mysql"
)

var dbSchema string 	= "Passage_master"
var serverTbl string 	= "servers"
var appTbl string 		= "apps"

type MysqlClient struct {
	user string
	password string
}

type ServerClient struct {
	exposeUrl string
	exposePort uint32
}

func (m *MysqlClient) execute(query string) (*sql.Rows, error) {
	/*
		Executes a given query statement.
	*/
	glog.Info("Executing query: ", query)
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", m.user, m.password, dbSchema))
	if err != nil {return nil, err}
	defer db.Close()

	if results, err := db.Query(query); err != nil {
		return nil, err
	} else {return results, err}
}

func (m *MysqlClient) AddServer(serverUrl string, serverPort uint32) error {
	glog.Info(fmt.Sprintf("Registering server: %s:%d", serverUrl, serverPort))
	query := fmt.Sprintf("INSERT INTO %s.%s (url,port) VALUES (%s,%d)", dbSchema, serverTbl, serverUrl, serverPort)
	if _, err := m.execute(query); err != nil {
		return err
	} else {return nil}
}

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