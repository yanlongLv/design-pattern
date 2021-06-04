package adapter

type Target interface {
	Request() string
}

type Adaptee interface {
	SpecificRequest() string
}

type adapteeImpl struct{}

type adapter struct {
	Adaptee
}

func (*adapteeImpl) SpecificRequest() string {
	return "adapeteeimpl"
}

func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

//Request 实现Target接口
func (a *adapter) Request() string {
	return a.SpecificRequest()
}
