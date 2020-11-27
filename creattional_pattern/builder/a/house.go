package a

/*House is struct*/
type House struct {
	windowType string
	doorType string
	floor int
}

/*GetwindowType is getter*/
func (h *House)GetWindowType() string {
	return h.windowType
}

/*GetdoorType is getter*/
func (h *House)GetDoorType() string{
	return h.doorType
}
/*Getfloor is getter*/

func (h *House)Getfloor() int{
	return h.floor
}