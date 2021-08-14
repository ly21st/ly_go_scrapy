package model

import (
	"encoding/json"
	"fmt"
	"time"
	"yannscrapy/logger"
)

// 构造用户表
type User struct {
	Id          int32  `json:"-"`
	Name        string `json:"name"`
	Pwd         string `json:"password"`
	Phone       int64  `json:"-"`
	Email       string `json:"-"`
	CreatedAt *time.Time `json:"-"`
	UpdateTAt  *time.Time `json:"-"`
}

// LoginReq请求参数
type LoginReq struct {
	Name string `json:"name"`
	Pwd  string `json:"password"`
}


// 插入数据
func (user *User) Insert() error {

	userInfo, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = DB.Put([]byte(user.Name), userInfo, nil)
	return err
}


// 用户注册
func Register(username, pwd string, phone int64, email string) error {
	logger.Infof("%s %s %s %s", username, pwd, phone, email)

	if CheckUser(username) {
		logger.Errorf("user " + username + " already exists")
		return fmt.Errorf("user " + username + " already exists")
	}

	// 需要生成一个uuid: Id为自增
	// 构造用户注册信息
	user := User{
		Name:  username,
		Pwd:   pwd,
		Phone: phone,
		Email: email,
	}
	insertErr := user.Insert()
	return insertErr
}

// 用户检查
func CheckUser(username string) bool {

	result := false
	_, err := DB.Get([]byte(username), nil)
	if err == nil {
		logger.Errorf("user " + username + " already exists")
		result = true
	}

	return result
}

// LoginCheck验证
func LoginCheck(login LoginReq) (bool, User, error) {
	userData := User{}
	userExist := false

	var user User
	//dbErr := DB.Where("name = ?", login.Name).Find(&user).Error

	userInfo, dbErr := DB.Get([]byte(login.Name), nil)
	if dbErr != nil {
		return userExist, userData, dbErr
	}
	err := json.Unmarshal(userInfo, &user)
	if err != nil {
		logger.Error(err)
		return userExist, userData, err
	}

	if login.Name == user.Name && login.Pwd == user.Pwd {
		userExist = true
		userData.Name = user.Name
		userData.Email = user.Email
	}

	if !userExist {
		return userExist, userData, fmt.Errorf("login error")
	}
	return userExist, userData, nil
}


