package grabbing_ticket_service

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"time"
	"yannscrapy/logger"
	"yannscrapy/logging"

	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb"
)

/**
获取用户列表
*/
func GetUserList(c *gin.Context) {

	result := make([]UserType, 0)
	db, err := leveldb.OpenFile(TargetUser, nil)
	if err != nil {
		logger.Error(err)
	}
	defer db.Close()

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.

		val := iter.Value()
		user := UserType{}
		json.Unmarshal(val, &user)
		result = append(result, user)
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		logger.Error(err)
	}

	c.JSON(200, gin.H{
		"data": result,
	})
}

func AddUser(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error(err)
		rspMsg := UserRegisterRspMsg{
			Code: "0001",
			Msg:  "read body error",
		}
		c.JSON(500, &rspMsg)
		return
	}
	logger.Infof("receive body:%s", string(body))

	user := UserType{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		logger.Error(err)
		rspMsg := UserRegisterRspMsg{
			Code: "0002",
			Msg:  "request body error",
		}
		c.JSON(400, &rspMsg)
		return
	}

	userId := strings.Trim(user.UserId, " ")
	password := strings.Trim(user.Password, " ")
	if userId == "" || password == "" {
		logger.Errorf("userId or password is empty")
		rspMsg := UserRegisterRspMsg{
			Code: "0003",
			Msg:  "request body error",
		}
		c.JSON(400, &rspMsg)
		return
	}

	rspMsg := UserRegisterRspMsg{
		Code: "0000",
		Msg:  "ok",
	}
	user.Ctime = time.Now()

	db, err := leveldb.OpenFile(TargetUser, nil)
	if err != nil {
		logging.Error(err)
		rspMsg.Code = "0004"
		rspMsg.Msg = err.Error()
		c.JSON(500, &rspMsg)
		return
	}
	defer db.Close()
	data, err := db.Get([]byte(userId), nil)
	if err == nil && len(data) != 0 {
		msg := "user " + userId + " already exists"
		logger.Errorf(msg)
		rspMsg.Code = "0005"
		rspMsg.Msg = msg
		c.JSON(400, &rspMsg)
		return
	}

	err = db.Put([]byte(userId), body,nil)
	if err != nil {
		logger.Error(err)
		rspMsg.Code = "0006"
		rspMsg.Msg = err.Error()
		c.JSON(500, &rspMsg)
		return
	}

	c.JSON(200, &rspMsg)
}

func DeleteUser(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error(err)
		rspMsg := UserRegisterRspMsg{
			Code: "0001",
			Msg:  "read body error",
		}
		c.JSON(500, &rspMsg)
		return
	}
	logger.Infof("receive body:%s", string(body))

	user := UserType{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		logger.Error(err)
		rspMsg := UserRegisterRspMsg{
			Code: "0002",
			Msg:  "request body error",
		}
		c.JSON(400, &rspMsg)
		return
	}

	userId := strings.Trim(user.UserId, " ")
	if userId == "" {
		logger.Errorf("userId is empty")
		rspMsg := UserRegisterRspMsg{
			Code: "0003",
			Msg:  "request body error",
		}
		c.JSON(400, &rspMsg)
		return
	}

	rspMsg := UserRegisterRspMsg{
		Code: "0000",
		Msg:  "ok",
	}
	user.Ctime = time.Now()

	db, err := leveldb.OpenFile(TargetUser, nil)
	if err != nil {
		logging.Error(err)
		rspMsg.Code = "0004"
		rspMsg.Msg = err.Error()
		c.JSON(500, &rspMsg)
		return
	}
	defer db.Close()
	data, err := db.Get([]byte(userId), nil)
	if err != nil || len(data) == 0 {
		msg := "user " + userId + " not exists"
		logger.Errorf(msg)
		rspMsg.Code = "0005"
		rspMsg.Msg = msg
		c.JSON(400, &rspMsg)
		return
	}

	err = db.Delete([]byte(userId),nil)
	if err != nil {
		logger.Error(err)
		rspMsg.Code = "0006"
		rspMsg.Msg = err.Error()
		c.JSON(500, &rspMsg)
		return
	}

	c.JSON(200, &rspMsg)
}
