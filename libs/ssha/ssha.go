// Package ssha ssha加解密
package ssha

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"

	"transformer_mz/libs/errors"
)

func createSHA1Sum(password string, salt []byte) []byte {
	sum := sha1.Sum(append([]byte(password), salt[:]...))
	return append(sum[:], salt[:]...)
}

func generateSSHA(password string, salt []byte) string {
	return "{SSHA}" + base64.StdEncoding.EncodeToString(createSHA1Sum(password, salt))
}

// GenerateSSHA 生成SSHA格式的密码
func GenerateSSHA(password string) (string, error) {
	salt := make([]byte, 4)
	_, err := rand.Reader.Read(salt)
	if err != nil {
		return "", errors.New(err, "内部服务出错，请联系开发人员进行处理")
	}
	return generateSSHA(password, salt), nil
}

// ValidateSSHA 校验密码是否正确
func ValidateSSHA(password string, hash string) (bool, error) {
	if len(hash) < 9 {
		return false, nil
	}
	data, err := base64.StdEncoding.DecodeString(hash[6:])
	if len(data) < 24 || err != nil {
		return false, errors.New(err, "内部服务出错，请联系开发人员进行处理")
	}
	return generateSSHA(password, data[20:]) == hash, nil
}
