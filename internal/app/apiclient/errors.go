package apiclient

// UnableConnectToServer ...
type UnableConnectToServer struct {
}

func (u *UnableConnectToServer) Error() string {
	return "unable connect to Burst server"
}

// WrongResponseStatus ...
type WrongResponseStatus struct {
}

func (w *WrongResponseStatus) Error() string {
	return "wrong response status"
}
