package passage_master

/*
Generic handlers module that we can use for different communication protocols
*/

func helloWorldHandler() []byte {return []byte("Hello world")}

func heartbeatHandler() []byte {return []byte("1")}