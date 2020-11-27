package a

type Ibuilder interface {
	setWindowType()
	setDoorType()
	setFloor()
	getHouse() House // Build ra cai house

}

/*la 1 phuong thuc, muc dich giups chung ta tra ve 1 builder tuong ung khi ngta truyen vao chuoi, chung ta se tra ve 1 tuong ung voi concrete builder*/
func GetBuilder(builderType string) Ibuilder {
	switch builderType{
	case "nomal":
		return &nomalBuilder{}
	case "igloo":
		return &iglooBuilder{}
	}
	return nil
}
