package interf

import "github.com/okieoth/hort/examples/crud/types"

type PersonRepository interface {
	Open(hortName string) error
	Close() error
	AddPerson(person types.Person) (types.Person, error)
	UpdatePerson(person types.Person) (types.Person, error)
	GetPerson(id int64) (types.Person, error)
	ListPerson() ([]types.Person, error)
	DeletePerson(id int64) error
}
