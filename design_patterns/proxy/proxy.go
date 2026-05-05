package proxy

type Subject interface {
	Do() string
}

type RealSubject struct{}

func (RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	real Subject
}

func NewProxy(real Subject) Subject {
	return &Proxy{real: real}
}

func (p *Proxy) Do() string {
	res := "pre:"
	res += p.real.Do()
	res += ":after"
	return res
}
