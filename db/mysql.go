package db

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

var dbMysqlMap map[string]*MysqlDBWrap

type MysqlDBWrap struct {
	Host     string
	Username string
	Password string
	Schema   string
	Charset  string
	Loc      string

	d     *sql.DB
	valid bool
}

func (db *MysqlDBWrap) genMysqlConnString() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&loc=%s&parseTime=true", db.Username, db.Password, db.Host, db.Schema, db.Charset, db.Loc)
}

func (db *MysqlDBWrap) init() error {
	var err error
	if db.Charset == "" {
		db.Charset = "utf8mb4"
	}

	if db.Loc == "" {
		db.Loc = "UTC"
	}

	connString := db.genMysqlConnString()
	db.d, err = sql.Open("mysql", connString)
	if err == nil {
		err = db.d.Ping()
		if err == nil {
			db.valid = true
			return nil
		}
	}
	return err
}

func (db *MysqlDBWrap) SetLoc(loc string) {
	db.Loc = url.QueryEscape(loc)
}

func (db *MysqlDBWrap) Valid() bool {
	return db.valid
}

func M(key string) *sql.DB {
	db, ok := dbMysqlMap[key]
	if !ok {
		return nil
	}
	if !db.Valid() {
		return nil
	}
	return db.d
}

func InitMysqlDB(key string, wrap *MysqlDBWrap) error {
	err := wrap.init()
	if err != nil {
		return err
	}

	dbMysqlMap[key] = wrap
	return nil
}
