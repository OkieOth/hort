package hort

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/okieoth/hort/examples/crud/types"
	p "github.com/okieoth/hort/pkg/jsonschemaparser"
)

type Hort struct {
	name       string
	dbFileName string
	db         *sql.DB
}

func NewHort(hortName string) *Hort {
	return &Hort{
		name: hortName,
	}
}

func (h *Hort) Open() error {
	name := p.ToProperName(h.name)
	h.dbFileName = fmt.Sprintf("../temp/%s.db", name)
	db, err := sql.Open("sqlite3", h.dbFileName)
	if err != nil {
		log.Fatal(err)
		return err
	}
	h.db = db
	return nil // TODO
}

func (h *Hort) Close() error {
	return h.db.Close()
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
