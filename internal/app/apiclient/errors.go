package apiclient

// UnableConnectToServer - 404
type UnableConnectToServer struct {
}

func (u *UnableConnectToServer) Error() string {
	return "unable connect to Burst server"
}

// WrongResponseStatus - when server sent unexpected status code
type WrongResponseStatus struct {
}

func (w *WrongResponseStatus) Error() string {
	return "wrong response status"
}
