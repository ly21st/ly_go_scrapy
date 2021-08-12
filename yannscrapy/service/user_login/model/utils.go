package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/syndtr/goleveldb/leveldb"
	"yannscrapy/logger"
	"yannscrapy/service/user_login/api"

)

var (
	DB *leveldb.DB
)

func InitDBCon() (err error) {
	api.ParserConfig()
	DB, err = leveldb.OpenFile(api.DbConfig.Database, nil)
	if err != nil {
		logger.Error(err)
	}
	return err
}
