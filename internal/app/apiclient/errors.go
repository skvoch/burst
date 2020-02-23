package apiclient

type UnableConnectToServer struct {
}

func (u *UnableConnectToServer) Error() string {
	return "unable connect to Burst server"
}
