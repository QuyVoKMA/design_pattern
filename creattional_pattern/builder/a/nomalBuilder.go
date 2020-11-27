package a

type nomalBuilder struct {
	// Giu cho trang thai hien tai, tam thoi khi chung ta khoi tao, house struct co j thi no co cai do
	windowType string
	doorType string
	floor int
}


func newNomalBuild() *nomalBuilder {
	return &nomalBuilder{}
}

func (n *nomalBuilder) setWindowType() {
	 n.windowType="Wooden windows"
}

func (n *nomalBuilder) setDoorType() {
	n.doorType="Wooden door"
}


func (n *nomalBuilder) setFloor() {
	n.floor=2
}

func (n *nomalBuilder) getHouse() House{
	return House{
		windowType: n.windowType,
		doorType: n.doorType,
		floor: n.floor,
	}
}