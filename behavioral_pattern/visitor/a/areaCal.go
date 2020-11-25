package a

import "fmt"

// Thêm cái hành vi tính diện tích cảu hình này, nó ko ảnh hưởng tới các thành phàn khác của các file dưới

type AreaCal struct {
	area int
}



func (a *AreaCal) visitForSquare(s *Square){
 fmt.Println("Caculator area for square")
}

func (a *AreaCal) visitForCirCle(c *Circle){
	fmt.Println("Caculator area for circle")
}

func (a *AreaCal) visitForRectange(r *Rectange){
	fmt.Println("Caculator area for rectange")
}

