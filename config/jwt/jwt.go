package jwt

import (
	"easy-go-admin/app/constant"
	"easy-go-admin/app/entity"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type userStdClaims struct {
	entity.JwtSystemAccountData
	jwt.StandardClaims
}

// TokenExpireDuration 过期时间
const TokenExpireDuration = time.Hour * 2

// Secret 密钥
const Secret = "easyadmin"

var (
	ErrAbsent  = "令牌不存在"
	ErrInvalid = "令牌无效"
)

// GenerateTokenByAdmin 生成用户token
func GenerateTokenByAdmin(account entity.SystemAccount) (string, error) {
	var jwtAdmin = entity.JwtSystemAccountData{
		Id:       account.ID,
		UserName: account.UserName,
		NickName: account.NickName,
		Avatar:   account.Avatar,
		Email:    account.Email,
	}
	c := userStdClaims{
		jwtAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "admin",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(Secret))
}

// ValidateToken 解析token
func ValidateToken(tokenString string) (*entity.JwtSystemAccountData, error) {
	if len(tokenString) == 0 {
		return nil, errors.New(ErrAbsent)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if token == nil {
		return nil, errors.New(ErrInvalid)
	}
	claims := userStdClaims{}
	_, err = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	return &claims.JwtSystemAccountData, err
}

// GetAccountID 返回用户id
func GetAccountID(c *gin.Context) (uint, error) {
	u, exist := c.Get(constant.ContextKeyUserObj)
	if !exist {
		return 0, errors.New("未找到id")
	}
	admin, ok := u.(entity.JwtSystemAccountData)
	if !ok {
		return admin.Id, nil
	}
	return 0, errors.New("未找到id")
}
