package tokens

import (
	"GuTikTok/config"
	"errors"
	"github.com/golang-jwt/jwt/v5"

	"time"
)

type MyClaims struct {
	ID int64 `json:"id"`
	jwt.RegisteredClaims
}

// GetToken 生成token
func GetToken(id int64) (string, error) {
	expireTime := time.Now().Add(time.Hour * 24 * 90) // 三个月过期
	SetClaims := MyClaims{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    "GuTikTok",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	return reqClaim.SignedString([]byte(config.Conf.JwtSecret))
}

// CheckToken 验证token
func CheckToken(token string) (*MyClaims, error) {
	key, err := jwt.ParseWithClaims(token, &MyClaims{}, func(*jwt.Token) (any, error) {
		return []byte(config.Conf.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := key.Claims.(*MyClaims); ok && key.Valid {
		return claims, nil
	} else {
		return nil, errors.New("你的token已过期")
	}
}
