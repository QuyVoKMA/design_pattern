package a

type INode interface {
	Clone() INode
	Print(s string)
}
