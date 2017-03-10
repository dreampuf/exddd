package models
import "time"

type Post struct {
    Id                  int64       `gorm:"primary_key"`
    Title               string
    User                User        `gorm:"polymorphic:Author"`
    Body                string

    DateCreated         time.Time   `sql:"DEFAULT:current_timestamp"`
    DateUpdated         time.Time   `sql:"DEFAULT:current_timestamp"`
    DateDeleted         time.Time   `sql:"DEFAULT:null"`
}
