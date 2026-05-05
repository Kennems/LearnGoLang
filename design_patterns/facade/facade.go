package facade

type API interface {
	Test() string
}

func NewAPI() API {
	return &apiImpl{
		a: NewAImpl(),
		b: NewBImpl(),
	}
}

type apiImpl struct {
	a AModule
	b BModule
}

func (api *apiImpl) Test() string {
	// 把两部分结果拼接返回
	return api.a.TestA() + api.b.TestB()
}

// --- AModule ---

type AModule interface {
	TestA() string
}

type AImpl struct{}

func NewAImpl() AModule {
	return &AImpl{}
}

func (a *AImpl) TestA() string {
	return "testA ... "
}

// --- BModule ---

type BModule interface {
	TestB() string
}

type BImpl struct{}

func NewBImpl() BModule {
	return &BImpl{}
}

func (b *BImpl) TestB() string {
	return "testB ... "
}
