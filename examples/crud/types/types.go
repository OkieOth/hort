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

  _id_ int64 // Storage Primary Key

	age int

	contact PersonContact

	roles []PersonRolesItems

	isActive bool

	createdAt time.Time // DateTime

	iEnum PersonIEnum

	id string // UUID

	name PersonName

}

type PersonContact struct {

	email string

	phone string

	address PersonContactAddress

}

type PersonContactAddress struct {

	city string

	postalCode string

	country string

	street string

}

type PersonName struct {

	first string

	middle string

	last string

}


