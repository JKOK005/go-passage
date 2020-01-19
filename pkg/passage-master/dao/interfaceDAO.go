package dao

type DAOInterface interface {
	InsertServer(serverUrl string, serverPort uint32) error
	GetServer(serverUrl string, serverPort uint32) (uint32, error)
	DeleteServer(serverID uint32)
	InsertApp(serverID uint32, appName string) error
	GetApp(serverID uint32, appName string) (uint32, error)
	GetApps(appName string) ([]ServerTbl, error)
	DeleteApp(appID uint32) error
}
