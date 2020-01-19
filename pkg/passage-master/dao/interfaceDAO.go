package dao

type DAOInterface interface {
	AddServer(serverUrl string, serverPort uint32) error
	GetServer(serverUrl string, serverPort uint32) (uint32, error)
	RemoveServer(serverID uint32)
	AddApp(serverID uint32, appName string) error
	GetApp(serverID uint32, appName string) (uint32, error)
	GetApps(appName string) ([]ServerClient, error)
	RemoveApp(appID uint32) error
}
