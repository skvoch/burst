package apiclient

type UnableConnectToServer struct {

}

func (u *UnableConnectToServer) Error() string {
	return "unable connect to Burst server"
}

type WrongResponseStatus struct {
	
}

func (w *WrongResponseStatus) Error() string {
	return "wrong response status"
}