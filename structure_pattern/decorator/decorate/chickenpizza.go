package decorate

type ChickenPizza struct {
ad string
}

func (cp *ChickenPizza)Dopizza() string {
	return "ChickenPizza"

}