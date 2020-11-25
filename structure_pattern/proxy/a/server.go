package a

/*Server is interface*/
type Server interface {
	//Hanh vi
	HandleRequest(string, string) (int, string)
}
