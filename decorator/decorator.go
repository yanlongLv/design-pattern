package decorator

import "fmt"

type Idraw interface {
	Draw()
}

type Roll struct{}

func (r *Roll) Draw() {

	fmt.Printf("I draw roll")
}

type RollLen struct {
	Idraw,
	Len int
}

func (r *RollLen) Draw() {

	fmt.Printf("I draw roll", r.Len)
}
