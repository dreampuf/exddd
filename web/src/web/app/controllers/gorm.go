package controllers

import (
    r "github.com/revel/revel"
    "github.com/jinzhu/gorm"
    "database/sql"

    _ "github.com/mattn/go-sqlite3"

    "web/app/models"
)

// type: revel controller with `*gorm.DB`
// c.Txn will keep `Gdb *gorm.DB`
type GormController struct {
    *r.Controller
    Txn *gorm.DB
}

// it can be used for jobs
var Gdb gorm.DB

// init db
func InitDB() {
    var err error
    db_driver := r.Config.StringDefault("db.driver", "sqlite3")
    db_url := r.Config.StringDefault("db.url", "sqlite://sqlite.db")
    // open db
    Gdb, err = gorm.Open(db_driver, db_url)
    if err != nil {
        r.ERROR.Println("FATAL", err)
        panic( err )
    }
    Gdb.AutoMigrate(&models.User{})
    // uniquie index if need
    //Gdb.Model(&models.User{}).AddUniqueIndex("idx_user_name", "name")
}

// transactions

// This method fills the c.Txn before each transaction
func (c *GormController) Begin() r.Result {
    txn := Gdb.Begin()
    if txn.Error != nil {
        panic(txn.Error)
    }
    c.Txn = txn
    return nil
}

// This method clears the c.Txn after each transaction
func (c *GormController) Commit() r.Result {
    if c.Txn == nil {
        return nil
    }
    c.Txn.Commit()
    if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
        panic(err)
    }
    c.Txn = nil
    return nil
}

// This method clears the c.Txn after each transaction, too
func (c *GormController) Rollback() r.Result {
    if c.Txn == nil {
        return nil
    }
    c.Txn.Rollback()
    if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
        panic(err)
    }
    c.Txn = nil
    return nil
}
