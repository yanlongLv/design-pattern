package factorymethod

type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) {
	o.a = a
}

func (o *OperatorBase) SetB(b int) {
	o.b = b
}

type PlusOperator struct {
	*OperatorBase
}

type OperatorFactory interface {
	Create() Operator
}
type PlusOperatorFactory struct{}

func (p PlusOperator) Result() int {
	return p.a + p.b
}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}
