package a

import "fmt"

type PerimateCal struct {
	perimater int
}

func (p *PerimateCal) visitorForSquere(s *Square){
	fmt.Println("Caculator perimater for square")
}

func (p *PerimateCal) visitorForCircle(c *Circle){
	fmt.Println("Caculator perimater for cirlce")
}

func (p *PerimateCal) visitorForRectange(r *Rectange){
	fmt.Println("Caculator perimater for rectange")
}