package simple_factory

const (
	APITypeHi = iota + 1
	APITypeHello
)

type Greeting interface {
	Say(msg string) string
}

func NewGreeting(t int) Greeting {
	if t == APITypeHi {
		return &Hi{}
	} else if t == APITypeHello {
		return &Hello{}
	}
	return nil
}

type Hi struct{}

func (h *Hi) Say(msg string) string {
	return "Hi! " + msg
}

type Hello struct{}

func (h *Hello) Say(msg string) string {
	return "Hello! " + msg
}
