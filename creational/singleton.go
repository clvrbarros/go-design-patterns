package creational

type Singleton interface {
	AddOne() int
}
// Singleton for single thread
type singleton struct {
	count int
}

var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		// create an instance and return pointer
		instance = new(singleton)
	}
	return instance
}
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}