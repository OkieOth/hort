package types

import "time"

type PersonDao struct {
	_id_ int64 // Storage Primary Key

	age int

	_contact int64

	isActive bool

	createdAt time.Time // DateTime

	iEnum PersonIEnum

	id string // UUID

	_name int64
}

type PersonRolesItemsDao struct {
	_id_ int64 // Storage Primary Key

	_person int64

	value string
}

type PersonContactDao struct {
	_id_ int64 // Storage Primary Key

	email string

	phone string

	_address int64
}

type PersonContactAddressDao struct {
	_id_ int64 // Storage Primary Key

	city string

	postalCode string

	country string

	street string
}

type PersonNameDao struct {
	_id_ int64 // Storage Primary Key

	first string

	middle string

	last string
}
