package a

type Department interface {
	Execute(*Patient)
	SetNext(Department)
}
