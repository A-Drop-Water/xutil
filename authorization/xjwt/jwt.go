package xjwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

// 对JWT的一个简单封装
// 但其实你会发现这个基本上没区别

type XJwt struct {
	cfg *Config
}

func NewXJwt(cfg *Config) *XJwt {
	return &XJwt{cfg: cfg}
}

type Config struct {
	Key []byte
}

type Claims interface {
	jwt.Claims
}

type RegisteredClaims = jwt.RegisteredClaims
type NumericDate = jwt.NumericDate

// 一共两个方法: 创建JWT Token,解析JWT Token

// CreateToken 默认了加密方式
func (j *XJwt) CreateToken(claims Claims) (string, error) {
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tk.SignedString(j.cfg.Key)
}

// ParseToken 解析
func (j *XJwt) ParseToken(token string, claims Claims) error {
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return j.cfg.Key, nil
	})
	if err != nil {
		return err
	}
	// 校验是不是有效
	if parsedToken == nil || !parsedToken.Valid {
		return errors.New("token is invalid")
	}
	// 有效
	return nil
}

// 一些配置信息: 过期时间、加密的key
