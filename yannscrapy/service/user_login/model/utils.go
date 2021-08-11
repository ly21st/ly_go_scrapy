package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/syndtr/goleveldb/leveldb"
    "yannscrapy/service/user_login/api"

)

var (
	DB *leveldb.DB
)

func InitDBCon() (err error) {
	DB, err = leveldb.OpenFile(api.DbConfig.Database, nil)
	return err
}
