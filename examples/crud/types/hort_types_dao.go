package types

type DaoPersonName struct {
  Id_ int64 // Storage Primary Key
  First string
  Middle string
  Last string
}

type DaoPersonContactAddress struct {
  Id_ int64 // Storage Primary Key
  Country string
  Street string
  City string
  PostalCode string
}

type DaoPersonContact struct {
  Id_ int64 // Storage Primary Key
  Phone string
  AddressId_ int64
  Email string
}

type DaoPerson struct {
  Id_ int64 // Storage Primary Key
  IsActive int32 // bool
  CreatedAt int64 // DateTime
  IEnum int32
  Id string // UUID
  NameId_ int64
  Age int64
  ContactId_ int64
}

type DaoPerson_roles struct {
  Id_ int64 // Storage Primary Key
  PersonId_ int64 // Reference to the parent entry
  Value string
}
