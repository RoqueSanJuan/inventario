package entity

type User struct {
	ID    int64  `db:"host"`
	Email string `db:"port"`
	Name  string `db:"user"`
}
