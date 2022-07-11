package user

import (
	"fmt"
	"gin-app-example/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignInReq struct {
	Username string `json:"username" binding:"required"` // 账号
	Password string // 密码
}
type SignInRes struct {
	Type  string
	Token string // token
}

// Auth godoc
// @Summary      账号密码登录
// @Description  账号密码登录
// @Tags         user
// @Param        body  body  SignInReq  true  "body传参"
// @Accept       json
// @Produce      json
// @Success      200  {object}  SignInRes
// @Router       /c/UserHandler/SignIn [post]
func (u *UserHandler) SignIn(c *gin.Context) {
	var r SignInReq
	err := c.BindJSON(&r)
	if err != nil {
		fmt.Println("register failed")
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		return
	}
	user := models.User{
		UserName: r.Username,
		Password: r.Password,
	}
	token, err := user.GetToken()
	if err != nil {
		fmt.Println("register failed")
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	signInRes := SignInRes{
		Type:  "Bearer",
		Token: token,
	}
	//验证 存储操作省略.....
	fmt.Println("register success")
	c.JSON(http.StatusOK, signInRes)

}
