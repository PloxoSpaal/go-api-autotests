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

func (f *Fake) Title() string {
	return f.faker.Sentence(4)
}

func (f *Fake) Description() string {
	return f.faker.Sentence(12)
}

func (f *Fake) Filename() string {
	return fmt.Sprintf("file-%s.txt", f.faker.UUID())
}

func (f *Fake) Directory() string {
	return fmt.Sprintf("uploads/%s", f.faker.UUID())
}

func (f *Fake) Content() []byte {
	return []byte(f.faker.Sentence(20))
}

func (f *Fake) MaxScore() int {
	return f.faker.Number(50, 100)
}

func (f *Fake) MinScore() int {
	return f.faker.Number(1, 30)
}

func (f *Fake) EstimatedTime() string {
	return fmt.Sprintf("%d weeks", f.faker.Number(1, 10))
}

func (f *Fake) OrderIndex() int {
	return f.faker.Number(1, 100)
}
