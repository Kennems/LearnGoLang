package singleton

import "sync"

type Singleton interface {
	Foo()
}

type singleton struct{}

func (s *singleton) Foo() {}

var (
	once     sync.Once
	instance Singleton
)

func GetSingleton() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
