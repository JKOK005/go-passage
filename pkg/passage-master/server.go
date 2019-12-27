package passage_master

type MasterServer struct
{
	httpServerPort int
}

func (m *MasterServer) startHttpServer () (int, error) {
	/*
	Begins http server setup.
	This is used to service application requests via http api calls.

	Params
		nil
	Returns
		int: Http status code
		error: Error when initializing server
	*/
}