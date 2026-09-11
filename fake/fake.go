package fake

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

type Fake struct {
	faker *gofakeit.Faker
}

func New() *Fake {
	return &Fake{faker: gofakeit.New(0)}
}

func (f *Fake) UUID() string {
	return f.faker.UUID()
}

func (f *Fake) Email() string {
	return f.faker.Email()
}

func (f *Fake) EmailWithDomain(domain string) string {
	return fmt.Sprintf("student-%s@%s", f.faker.UUID(), domain)
}

func (f *Fake) Password() string {
	return f.faker.Password(true, true, true, false, false, 16)
}

func (f *Fake) FirstName() string {
	return f.faker.FirstName()
}

func (f *Fake) LastName() string {
	return f.faker.LastName()
}

func (f *Fake) MiddleName() string {
	return f.faker.MiddleName()
}
