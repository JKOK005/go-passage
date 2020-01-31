package handlers

type DispatchHandlerServer struct {

}

func (d DispatchHandlerServer) DispatchHeartbeatMaster() {}

func (d DispatchHandlerServer) DispatchHeartbeatApp() {}

func (d DispatchHandlerServer) DispatchPayload() {}