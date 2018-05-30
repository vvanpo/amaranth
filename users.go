package cmf

import (
	"database/sql"
)

type Users struct {
}

// List returns a slice of users sorted by username.
func (u *Users) List(limit int, offset int) ([]*User, error) {

}

func (u *Users) Count(filter string) (int, error) {

}
