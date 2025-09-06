package types
import "time"


type PersonIEnum int
const (
    PersonIEnum_1 PersonIEnum = iota
    PersonIEnum2
    PersonIEnum4
    PersonIEnum8
    PersonIEnum16
)


type PersonRolesItems string
const (
	PersonRolesItemsadmin PersonRolesItems = "admin"
	PersonRolesItemseditor PersonRolesItems = "editor"
	PersonRolesItemsviewer PersonRolesItems = "viewer"
)

type Person struct {

  Id_ int64 // Storage Primary Key

	Roles []PersonRolesItems

	IsActive bool

	CreatedAt time.Time // DateTime

	IEnum PersonIEnum

	Id string // UUID

	Name PersonName

	Age int

	Contact PersonContact

}

type PersonContact struct {

	Email string

	Phone string

	Address PersonContactAddress

}

type PersonContactAddress struct {

	City string

	PostalCode string

	Country string

	Street string

}

type PersonName struct {

	Middle string

	Last string

	First string

}


