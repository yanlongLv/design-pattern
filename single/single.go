package single

import (
	"fmt"
	"sync"
)

type Single interface {
	foo()
}

type singleTon struct{}

func (s *singleTon) foo() {
	fmt.Println("foo single")
}

var (
	instance *singleTon
	once     sync.Once
)

func GetInstance() Single {
	once.Do(func() {
		instance = &singleTon{}
	})

	return instance
}
