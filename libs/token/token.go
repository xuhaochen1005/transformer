// Package token 实现token的生成和验证
package token

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"time"

	"gopkg.in/dgrijalva/jwt-go.v3"
	"transformer_mz/libs/errors"
	"transformer_mz/libs/log"
)

// 用于JWT加解密的密钥
const key = `
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCybeR5cNcwpvMzzlsKdbbyZZStMCuTv5f/5c85RtQIXiVqYo10
7DOBVlczcTBy4A4wPcULYlD8zNnGz3dd8kspkhruE92a/aV3EeULFGEvXh+41F5l
0sHZbSS9eOTePvMUvUaQZLvch/Byzj124HuR9bg+gxtAwZhOnX5qVKGHeQIDAQAB
AoGAGu3Q9K8eGx6nR+NWBC8d7Fl+ZeEGQqdA5oMlkkxpRdYHqZW0KbhYUaXZRU9I
851zJr671WsCNBUqrQG0zz35L81pIwpvRAn38UyXmwMyR8Jiq2Fx9uT6pWBealMA
BcvM791RsQ/GDtLBWOT4uLuJpTkiXr/5c4JRHaTze0Kv6/ECQQDd/tkf2rTgQIGH
pmCbZtsBF5LjyCsUhYuTDzeHeio45nus6VBGU9Skx0hZXupA8gc/B/tHH4XAJSqG
pdLEoehtAkEAzcKsX+fq3v454X7xAadaRe8TEZp9wNyOlSpXIy3cEV3ZvJhoOVi4
El2I0ICAWS4OBPDQzPGfu17m9DZow2VLvQJBAIYWZc1QzAnbzKFDxHx5HvVE6Ot+
v06w70JPYaoKBzDBPpiNgHhKMFmrKS8aHoJF7kyLJKKsDpyllk8oH+u3I+0CQH1P
pz1Nd/xQ5h1P+yVfr5nRzVI2PYn8iRWNUL5LCrDE8HtscmUihBAjSNR6vjAR3VXi
97cpHHe3h36JuGnvwh0CQHxcjzSQ5vO4BdV6wvgAW6y2vN7ZwxfeRGb9P7m7iFXR
o2xtIpb0FTvnTvxEwnYxdus+qyt0Pz+f9FW9reAmGEw=
-----END RSA PRIVATE KEY-----`

// privateKey RSA格式的私钥
var privateKey *rsa.PrivateKey

// Token token数据格式
type Token struct {
	ID        int   `json:"id"`
	ExpiresAt int64 `json:"exp"`
}

// Valid 校验token是否过期
func (token Token) Valid() error {
	now := time.Now().Unix()
	if token.ExpiresAt <= now {
		delta := time.Unix(now, 0).Sub(time.Unix(token.ExpiresAt, 0))
		return errors.New("登陆已过期，请重新登陆" + delta.String())
	}
	return nil
}

func init() {
	var err error
	data, _ := pem.Decode([]byte(key))
	privateKey, err = x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		log.Fatalln(err)
	}
}

// Create 根据给定ID生成token
func Create(id int) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)
	token.Claims = Token{
		ID:        id,
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}
	tokenStr, err := token.SignedString(privateKey)
	return tokenStr, errors.New(err, "内部服务出错，请联系开发人员处理")
}

// Validate 校验token有效性
func Validate(tokenStr string) (*Token, error) {
	claim, err := jwt.ParseWithClaims(tokenStr, &Token{}, func(_ *jwt.Token) (interface{}, error) {
		return &privateKey.PublicKey, nil
	})
	if err != nil {
		return nil, errors.New(err)
	}
	token := claim.Claims.(*Token)
	return token, token.Valid()
}
