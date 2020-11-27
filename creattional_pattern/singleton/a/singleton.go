package a

/*Singleton is interface*/
type Singleton interface {
	AddOne() int
}

type singleton struct{
	count int
}

var instance *singleton

/*Khi chạy chương trinh nó sẽ hàm init này, nó sẽ khỏi tạo đối tượng mối gán vào instance*/
func init() {
	instance =&singleton{count: 100}
}
/*GetInstance is interface*/
func GetInstance() Singleton{
	return instance
}

func (s *singleton) AddOne()  int{
	s.count++
	return s.count
}