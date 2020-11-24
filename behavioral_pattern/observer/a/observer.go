package a


/* Observer is inteface*/
type Observer interface {
	Update(string)
	GetID() string // ID
}
