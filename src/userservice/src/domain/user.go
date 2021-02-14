package domain

type Users []User

type User struct {
	ID         string
	FirstName  string
	MiddleName string
	LastName   string
	Gender     uint32
	Birthday   Birthday
	Email      string
	Password   string
}

type Birthday struct {
	Year  uint32
	Month uint32
	Day   uint32
}
