package bridge

type Software interface {
	Run() string
}

type SoftwareA struct{}
func (s *SoftwareA) Run() string { return "run soft a" }

type SoftwareB struct{}
func (s *SoftwareB) Run() string { return "run soft b" }

type Phone interface {
	SetSoft(software Software)
	Run() string
}

type PhoneA struct {
	soft Software
	name string
}
func NewPhoneA(name string) *PhoneA { return &PhoneA{name: name} }
func (p *PhoneA) SetSoft(software Software) { p.soft = software }
func (p *PhoneA) Run() string { return "PhoneA " + p.name + " run: " + p.soft.Run() }

type PhoneB struct {
	soft Software
	name string
}
func NewPhoneB(name string) *PhoneB { return &PhoneB{name: name} }
func (p *PhoneB) SetSoft(software Software) { p.soft = software }
func (p *PhoneB) Run() string { return "PhoneB " + p.name + " run: " + p.soft.Run() }
