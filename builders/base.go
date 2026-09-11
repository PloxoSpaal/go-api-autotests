package builders

import "github.com/PloxoSpaal/go-api-autotests/fake"

type Builder struct {
	generator *fake.Fake
}

func New(generator *fake.Fake) *Builder {
	return &Builder{generator: generator}
}
