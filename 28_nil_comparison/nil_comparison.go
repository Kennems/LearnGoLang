package main

type nilCoder interface {
	Code() string
}

type gopher struct {
	Name string
}

func (g *gopher) Code() string {
	if g == nil {
		return ""
	}
	return g.Name
}

func typedNilAsInterface() any {
	var p *int
	return p
}

func nilCoderValue() nilCoder {
	var g *gopher
	return g
}

func compareInterfaces(a, b any) bool {
	return a == b
}

func compareMayPanic(a, b any) (equal bool, panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()

	return a == b, false
}
