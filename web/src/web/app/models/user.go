package models

import (
    "fmt"
    "github.com/revel/revel"
    "regexp"
    "time"
)

type User struct {
	Id                  int64       `gorm:"primary_key"`
	Name                string
	Nickname            string
    Password            string      `sql:"-"`
	HashedPassword      []byte

	WeiboID				string
    WeiboToken          string
	WeiboExpires		time.Time	`sql:"DEFAULT:current_timestamp`

    DateCreated         time.Time   `sql:"DEFAULT:current_timestamp"`
    DateUpdated         time.Time   `sql:"DEFAULT:current_timestamp"`
	DateDeleted         time.Time   `sql:"DEFAULT:null"`
}

func (u *User) String() string {
	return fmt.Sprintf("User(%s)", u.Nickname)
}

var userRegex = regexp.MustCompile("^\\w*$")

func (user *User) Validate(v *revel.Validation) {
	v.Check(user.Nickname,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{4},
		revel.Match{userRegex},
	)

	ValidatePassword(v, user.Password).
		Key("user.Password")

	v.Check(user.Name,
		revel.Required{},
		revel.MaxSize{100},
	)
}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{5},
	)
}
