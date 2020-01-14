package storage

type MysqlClient struct {}
type ServerClient struct {
	exposeUrl string
	exposePort uint32
}

func (m *MysqlClient) AddServer(serverUrl string, serverPort uint32) (uint32, error) {return -1, nil}
func (m *MysqlClient) RemoveServer(serverID uint32) error {return nil}
func (m *MysqlClient) AddApp(serverID uint32, appName string) (uint32, error) {return -1, nil}
func (m *MysqlClient) RemoveApp(appID uint32) error {return nil}
func (m *MysqlClient) GetAppLoc(appName string) ([]ServerClient, error) {return nil, nil}

func CreateClient() (*MysqlClient, error) {
	/*
		Factory dispatcher pattern for new ZK client
	*/
	return nil, nil
}