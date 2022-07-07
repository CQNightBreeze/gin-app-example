package user

import (
	"gin-app-example/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

type GetUserInfoConditionalReq struct {
	UserId uint
}
type GetUserInfoConditionalRes struct {
	UserId uint
}

// Auth godoc
// @Summary      根绝条件获取用户
// @Description  根绝条件获取用户
// @Tags         user
// @Param        UserId  query  int  true  "用户id"
// @Accept       json
// @Produce      json
// @Success      200  {object}  GetUserInfoConditionalRes
// @Security     auth
// @Router       /c/UserHandler/GetUserInfoConditional [get]
func (u *UserHandler) GetUserInfoConditional(c *gin.Context) {
	var r GetUserInfoConditionalReq
	err := c.ShouldBindQuery(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	resData := GetUserInfoConditionalRes{
		UserId: r.UserId,
	}
	c.JSON(http.StatusOK, resData)
}

type GetUserInfoRes struct {
	Username string `json:"username" binding:"required"` // 账号
}

// Auth godoc
// @Summary      获取用户信息
// @Description  获取用户信息
// @Tags         user
// @Param        Authorization  header  string  true  "授权信息"
// @Accept       json
// @Produce      json
// @Success      200  {object}  GetUserInfoRes
// @Security     auth
// @Router       /c/UserHandler/GetUserInfo [get]
func (u *UserHandler) GetUserInfo(c *gin.Context) {

	authorization := c.Request.Header["Authorization"]
	if authorization == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "not Authorization"})
		return
	}
	authorizations := strings.Split(authorization[0], " ")
	user := models.User{}

	if authorizations[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "not Authorization"})
		return
	}
	user.VerifyToken(authorizations[1])

	getUserInfoRes := GetUserInfoRes{
		Username: user.UserName,
	}
	c.JSON(http.StatusOK, getUserInfoRes)
}
