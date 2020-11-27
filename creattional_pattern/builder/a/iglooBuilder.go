package a

type iglooBuilder struct {
	windowType string
	doorType string
	floor int
}

func newIglooBuilder() *iglooBuilder{
	return &iglooBuilder{}
}

func (i *iglooBuilder) setWindowType(){

	i.windowType="Snow windows"
}
func (i *iglooBuilder) setDoorType(){
	i.doorType="Snow Door"
}

func (i *iglooBuilder) setFloor(){
	i.floor=1
}

func (i *iglooBuilder) getHouse() House{
	return House{
		windowType:i.windowType,
		doorType:i.doorType,
		floor:i.floor,
	}
}
