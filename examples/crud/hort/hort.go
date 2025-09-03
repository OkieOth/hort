package hort

import "github.com/okieoth/hort/examples/crud/types"

type Hort struct {
	name string
}

func (h *Hort) Open(hortName string) error {
	return nil // TODO
}

func (h *Hort) Close() error {
	return nil // TODO
}

func (h *Hort) AddPerson(person types.Person) (types.Person, error) {
	return person, nil // TODO
}

func (h *Hort) UpdatePerson(person types.Person) (types.Person, error) {
	return person, nil // TODO
}

func (h *Hort) GetPerson(id int64) (types.Person, error) {
	return types.Person{}, nil // TODO
}

func (h *Hort) ListPerson() ([]types.Person, error) {
	return []types.Person{}, nil // TODO
}

func (h *Hort) DeletePerson(id int64) error {
	return nil // TODO
}
