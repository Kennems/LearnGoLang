package builder

// Builder 是生成器接口
type Builder interface {
	Part1()
	Part2()
	Part3()
}

// Director 构建指挥者
type Director struct {
	builder Builder
}

// NewDirector ...
func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

// Construct 构建产品
func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

// ---------------- Builder1 ----------------

type Builder1 struct {
	result string
}

func (b *Builder1) Part1()            { b.result += "1" }
func (b *Builder1) Part2()            { b.result += "2" }
func (b *Builder1) Part3()            { b.result += "3" }
func (b *Builder1) GetResult() string { return b.result }

// ---------------- Builder2 ----------------

type Builder2 struct {
	result int
}

func (b *Builder2) Part1()         { b.result += 1 }
func (b *Builder2) Part2()         { b.result += 2 }
func (b *Builder2) Part3()         { b.result += 3 }
func (b *Builder2) GetResult() int { return b.result }
