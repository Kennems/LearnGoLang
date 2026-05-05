package observer

type Subject struct {
	observers []Observer
	context   string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

type Observer interface {
	Update(*Subject)
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}

func (s *Subject) Context() string {
	return s.context
}

type Reader struct {
	name     string
	received []string
}

func NewReader(name string) *Reader {
	return &Reader{
		name:     name,
		received: make([]string, 0),
	}
}

func (r *Reader) Update(s *Subject) {
	r.received = append(r.received, s.Context())
}

func (r *Reader) Received() []string {
	return r.received
}

func (r *Reader) Name() string {
	return r.name
}
