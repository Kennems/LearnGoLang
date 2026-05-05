package prototype

// Cloneable 是原型对象需要实现的接口
type Cloneable interface {
	Clone() Cloneable
}

type PrototypeManager struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prototypes: make(map[string]Cloneable),
	}
}

// Get 返回指定名称的原型克隆对象
func (p *PrototypeManager) Get(name string) Cloneable {
	proto, ok := p.prototypes[name]
	if !ok {
		return nil
	}
	return proto.Clone()
}

// Set 注册原型对象
func (p *PrototypeManager) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}
