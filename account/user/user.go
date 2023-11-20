package user

import "time"

type User struct {
	ID        int
	Account   string
	Password  string
	Status    int
	UserApp   []UserApp
	CreatedAt time.Time
	UpdatedAt time.Time
}
