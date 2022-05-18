package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName string // 用户账号
	Password string // 密码
}

type CustomClaims struct {
	User *User
	jwt.StandardClaims
}

const key = "gin-app-example"

func (u *User) GetToken() (string, error) {
	expireToken := time.Now().Add(time.Hour * 72 * 10).Unix()

	// Create the Claims
	claims := CustomClaims{
		u,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "gin-app-example",
		},
	}
	// DefaultLogger.Error("key is：", util.GetEnvConfig().Key)
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString([]byte(key))

}

func (u *User) VerifyToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	// Validate the token and return the custom claims
	if token == nil {
		return nil, nil
	}
	switch token.Claims.(type) {
	case *CustomClaims:
		result := token.Claims.(*CustomClaims)
		u.UserName = result.User.UserName
		u.Password = result.User.Password
		return result, nil
	default:
		return nil, err
	}

}
