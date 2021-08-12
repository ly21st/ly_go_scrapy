package controller

import (
	_ "fmt"
	"log"
	"net/http"
	"time"
	md "yannscrapy/service/user_login/middleware"
	"yannscrapy/service/user_login/model"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type RegisterInfo struct {
	Phone int64  `json:"phone"`
	Name  string `json:"name"`
	Pwd   string `json:"password"`
	Email string `json:"email"`
}

func RegisterUser(c *gin.Context) {
	var registerInfo RegisterInfo
	bindErr := c.BindJSON(&registerInfo)
	if bindErr == nil {
		err := model.Register(registerInfo.Name, registerInfo.Pwd, registerInfo.Phone, registerInfo.Email)

		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    "success ",
				"data":   nil,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg":    "fail" + err.Error(),
				"data":   nil,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "bad request body" + bindErr.Error(),
			"data":   nil,
		})
	}
}

// 登陆结果
type LoginResult struct {
	Token string `json:"token"`
	// 用户模型
	Name string `json:"name"`
	//model.User
}

// name,password
func Login(c *gin.Context) {
	var loginReq model.LoginReq
	if c.BindJSON(&loginReq) == nil {
		isPass, user, err := model.LoginCheck(loginReq)
		if isPass {
			generateToken(c, user)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": -1,
				"msg":    "verification failed:" + err.Error(),
				"data":   nil,
			})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "bad request",
			"data":   nil,
		})
	}
}

func generateToken(c *gin.Context, user model.User) {
	j := md.NewJWT()

	claims := md.CustomClaims{
		user.Name,
		user.Email,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),
			ExpiresAt: int64(time.Now().Unix() + 3600),
			Issuer:    "yannscrapy",
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
			"data":   nil,
		})
	}

	log.Println(token)
	// 获取用户相关数据
	data := LoginResult{
		Name:  user.Name,
		Token: token,
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "login success",
		"data":   data,
	})
	return
}

// 测试一个需要认证的接口
func GetDataByTime(c *gin.Context) {
	claims := c.MustGet("claims").(*md.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token valid",
			"data":   claims,
		})
	}
}
