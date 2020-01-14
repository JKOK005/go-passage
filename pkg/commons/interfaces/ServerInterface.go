package interfaces

type ServerInterface interface
{
	startHttpServer() (bool, error)
	startGrpcServer() (bool, error)
}
