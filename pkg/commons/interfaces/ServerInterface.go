package interfaces

type ServerInterface interface
{
	startHttpServer() (int, error)
}
