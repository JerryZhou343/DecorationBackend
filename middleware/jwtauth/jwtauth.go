package jwtauth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// JWTAuth json web token
//gin 中间间，用于在分组路由中校验请求是否带有token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.DefaultQuery("token", "")
		if token == "" {
			token = c.Request.Header.Get("Authorization")
			if s := strings.Split(token, " "); len(s) == 2 {
				token = s[1]
			}
		}
		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == ErrTokenExpired {
				if token, err = j.RefreshToken(token); err == nil {
					c.Header("Authorization", "Bear "+token)
					//c.JSON(http.StatusOK, gin.H{"error": 0, "message": "refresh token", "token": token})
					return
				}
			}
			c.JSON(http.StatusUnauthorized, gin.H{"error": 1, "message": err.Error()})
			return
		}
		c.Set("claims", claims)
	}
}

//JWT jwt对象定义，含有jwt需要的签名KEY
type JWT struct {
	SigningKey []byte
}

var (
	//ErrTokenExpired token 已经过期
	ErrTokenExpired = errors.New("Token is expired")
	//ErrTokenNotValidYet token不可用
	ErrTokenNotValidYet = errors.New("Token not active yet")
	//ErrTokenMalformed token不合法
	ErrTokenMalformed = errors.New("That's not even a token")
	//ErrTokenInvalid token 不可用
	ErrTokenInvalid = errors.New("Couldn't handle this token")
	//SignKey  签名key
	SignKey = "test"
)

//CustomClaims 用户关键信息
type CustomClaims struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

//NewJWT 生成一个新的json web token
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

//GetSignKey 返回签名Key
func GetSignKey() string {
	return SignKey
}

//SetSignKey 设置签名KEY
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

//CreateToken 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

//ParseToken token逆转，获得用户信息
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

//RefreshToken  刷新Token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", ErrTokenInvalid
}
