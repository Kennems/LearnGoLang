package chain_of_responsibility
import "strconv"

type Request struct {
	RequestType string
	Number      int
}

type Manager interface {
	SetSuperior(manager Manager)
	RequestApplications(request Request) string
}

type CommonManager struct {
	name     string
	superior Manager
}
func NewCommonManager(name string) *CommonManager { return &CommonManager{name: name} }
func (c *CommonManager) SetSuperior(manager Manager) { c.superior = manager }
func (c *CommonManager) RequestApplications(request Request) string {
	if request.RequestType == "请假" && request.Number <= 2 {
		return c.name + ":" + request.RequestType + " 数量" + strconv.Itoa(request.Number) + " 被批准"
	} else if c.superior != nil {
		return c.superior.RequestApplications(request)
	}
	return ""
}

type Majordomo struct {
	name     string
	superior Manager
}
func NewMajordomo(name string) *Majordomo { return &Majordomo{name: name} }
func (m *Majordomo) SetSuperior(manager Manager) { m.superior = manager }
func (m *Majordomo) RequestApplications(request Request) string {
	if request.RequestType == "请假" && request.Number <= 5 {
		return m.name + ":" + request.RequestType + " 数量" + strconv.Itoa(request.Number) + " 被批准"
	} else if m.superior != nil {
		return m.superior.RequestApplications(request)
	}
	return ""
}

type GeneralManager struct {
	name     string
	superior Manager
}
func NewGeneralManager(name string) *GeneralManager { return &GeneralManager{name: name} }
func (g *GeneralManager) SetSuperior(manager Manager) { g.superior = manager }
func (g *GeneralManager) RequestApplications(request Request) string {
	if request.RequestType == "请假" {
		return g.name + ":" + request.RequestType + " 数量" + strconv.Itoa(request.Number) + " 被批准"
	} else if request.RequestType == "加薪" && request.Number <= 500 {
		return g.name + ":" + request.RequestType + " 数量" + strconv.Itoa(request.Number) + " 被批准"
	} else if request.RequestType == "加薪" && request.Number > 500 {
		return g.name + ":" + request.RequestType + " 数量" + strconv.Itoa(request.Number) + " 再说吧"
	}
	return ""
}
