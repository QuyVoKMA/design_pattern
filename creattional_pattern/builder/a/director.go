package a

type Director struct {
	builder Ibuilder
}

// Hàm khỏi tạo
// Hướng 1 để truyền vào
func NewDirector(b Ibuilder) *Director {
	return &Director{
		builder: b,
	}
}
// cai cái SetBuilder vào Director
// Hướng 2 để truyền vào.
func (d *Director) SetBuilder(b Ibuilder){
	d.builder=b
}

func (d *Director)BuildHouse() House{
	d.builder.setWindowType()
	d.builder.setDoorType()
	d.builder.setFloor()
	return d.builder.getHouse()
}




