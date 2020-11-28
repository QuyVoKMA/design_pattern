package a

type Originator struct {
	State string
}

func (e *Originator) CreateMemeto() *Memento{
	return &Memento{
		state:e.State,
	}
}
func (e *Originator) RestoreMemeto(m *Memento){
	e.State=m.state
}

func (e *Originator) GetState() string{
	return e.State
}

func (e *Originator) SetState(state string){
	 e.State=state
}