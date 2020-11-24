package a

/*Subject in interface*/
type Subject interface {
	//hàm register để register các observer
	Register(observer Observer)
	DeRegister(observer Observer)
	NotifyAll()
}
