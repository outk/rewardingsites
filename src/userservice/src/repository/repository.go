package repository

type User struct {
	ID          uint32
	first_name  string
	middle_name string
	last_name   string
	gender      uint32
	birthday    Birthday
	email       string
	password    string
}

type Birthday struct {
	Year  uint32
	Month uint32
	day   uint32
}
