package grabbing_ticket_service

import (
	"sync"
	"time"
)

type UserType struct {
	UserId string `json:"userId"`
	Password string `json:"password"`
	Ctime      time.Time `json:"-"`
}


type UserMgrType struct {
	UserMap map[string]*UserType
	Mutex sync.Mutex
	FilePath string
}

type UserRegisterRspMsg struct {
	Code string `json:"code"`
	Msg string `json:"msg"`
}

var  userMgr = &UserMgrType{
	UserMap: make(map[string]*UserType, 0),
	FilePath: UserFilePathConst,
}

func GetUserMgr() *UserMgrType {
	return userMgr
}
