package dao

import (
	"database/sql"
	"fmt"
	"github.com/golang/glog"
)

var dbSchema string 	= "Passage_master"
var serverTbl string 	= "servers"
var appTbl string 		= "apps"

type MysqlDAO struct {
	user string
	password string
}

type ServerTbl struct {
	id uint32
	url string
	port uint32
}

func (m *MysqlDAO) execute(query string) (*sql.Rows, error) {
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

func (m *MysqlDAO) InsertServer(serverUrl string, serverPort uint32) error {
	glog.Info(fmt.Sprintf("Inserting new server entry: %s:%d", serverUrl, serverPort))
	query := fmt.Sprintf("INSERT INTO %s.%s (url,port) VALUES (%s,%d)", dbSchema, serverTbl, serverUrl, serverPort)
	if _, err := m.execute(query); err != nil {
		return err
	} else {return nil}
}

func (m *MysqlDAO) GetServer(serverUrl string, serverPort uint32) (*ServerTbl, error) {
	glog.Info(fmt.Sprintf("Fetching server information: %s:%d", serverUrl, serverPort))
	query := fmt.Sprintf("SELECT * FROM %s.%s WHERE url='%s' AND port='%d'", dbSchema, serverTbl, serverUrl, serverPort)
	if row, err := m.execute(query); err != nil {
		return nil, err
	} else {
		var resp *ServerTbl
		row.Scan(resp)
		return resp, nil
	}
}

func (m *MysqlDAO) DeleteServer(serverID uint32) error {return nil}
func (m *MysqlDAO) InsertApp(serverID uint32, appName string) error {return nil}
func (m *MysqlDAO) GetApp(serverID uint32, appName string) (uint32, error) {return -1, nil}
func (m *MysqlDAO) GetApps(appName string) ([]ServerTbl, error) {return nil, nil}
func (m *MysqlDAO) DeleteApp(appID uint32) error {return nil}

func NewMysqlDAO() error {

}