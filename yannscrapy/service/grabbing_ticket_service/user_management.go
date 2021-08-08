package grabbing_ticket_service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"yannscrapy/logger"
)


/**
获取用户列表
 */
func GetUserList(c *gin.Context) {
	//var err error
	userMgr := GetUserMgr()
	result := make([]string, 0)
	{
		userMgr.Mutex.Lock()
		defer userMgr.Mutex.Unlock()
		for key, _ := range userMgr.UserMap {
			result = append(result, key)
		}
	}

	c.JSON(200, gin.H{
		"data": result,
	})
}


func AddUser(c *gin.Context) {
	//var err error

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error(err)
		rspMsg := UserRegisterRspMsg{
			Code : "0001",
			Msg: "read body error",
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
			Code : "0002",
			Msg: "request body error",
		}
		c.JSON(400, &rspMsg)
		return
	}

	userId := strings.Trim(user.UserId, " ")
	password := strings.Trim(user.Password, " ")
	if userId == "" || password == "" {
		logger.Errorf("userId or password is empty")
		rspMsg := UserRegisterRspMsg{
			Code : "0002",
			Msg: "request body error",
		}
		c.JSON(400, &rspMsg)
		return
	}

	rspMsg := UserRegisterRspMsg{
		Code: "0000",
		Msg: "ok",
	}
	user.Ctime = time.Now()
	userMgr := GetUserMgr()
	{
		dstFile, err := os.OpenFile(userMgr.FilePath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, UserFilePErmConst)
		defer dstFile.Close()
		if err != nil {
			logger.Error(err)
			rspMsg.Code = "0004"
			rspMsg.Msg = err.Error()
			c.JSON(500, &rspMsg)
			return
		}

		userMgr.Mutex.Lock()
		defer userMgr.Mutex.Unlock()
		for {
			_, ok := userMgr.UserMap[userId]
			if ok {
				msg := "user " + userId + " already exists"
				logger.Errorf(msg)
				rspMsg.Code = "0003"
				rspMsg.Msg = msg
				break
			}

			userMgr.UserMap[userId] = &user
			//fileBody, err := json.Marshal(&userMgr.UserMap)
			enc := json.NewEncoder(dstFile)
			enc.SetIndent("", "  ")
			// Dump json to the standard output
			err = enc.Encode(userMgr.UserMap)
			logger.Error(err)
			break
		}
	}

	if rspMsg.Code != "0000" {
		c.JSON(400, &rspMsg)
		return
	}

	c.JSON(200, &rspMsg)
}



func DeleteUser(c *gin.Context) {
	//var err error
	userMgr := GetUserMgr()
	result := make([]string, 0)
	{
		userMgr.Mutex.Lock()
		defer userMgr.Mutex.Unlock()
		for key, _ := range userMgr.UserMap {
			result = append(result, key)
		}
	}

	c.JSON(200, gin.H{
		"data": result,
	})
}