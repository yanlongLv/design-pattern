package observer

import "fmt"

type ISubject interface {
	Register(observer IObserver)
	Remove(observer IObserver)
	Notify(msg string)
}


type IObserver interface {
	Update(m string)
}

type Subject struct {
	observers []IObserver
}

func (s *Subject) Register(observer IObserver) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) Remove(observer IObserver) {
	for i,ob:=range s.observers {
		if ob == observer {
			s.observers = append(s.observers[:i],s.observers[i+1:]...)
		}
	}
}

func (s *Subject) Notify(smg string) {
	for _,o:=range s.observers {
		o.Update(smg)
	}
}

type IObserver1 struct {}

type IObserver2 struct {}

func (o *IObserver1) Update(msg string){
	fmt.Printf("Observice1")
}

func (o *IObserver2) Update(msg string){
	fmt.Printf("Observice2")
}