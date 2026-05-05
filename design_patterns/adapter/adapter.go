package adapter

type Adapter interface {
	Request() string
}

func NewAdapter(adaptee Adaptee) Adapter {
	return &adapter{
		adaptee: adaptee,
	}
}

type adapter struct {
	adaptee Adaptee
}

func (a *adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

// --- Adaptee 接口与实现 ---

type Adaptee interface {
	SpecificRequest() string
}

type adaptee struct{}

func NewAdaptee() Adaptee {
	return &adaptee{}
}

func (a *adaptee) SpecificRequest() string {
	return "specific request called"
}
