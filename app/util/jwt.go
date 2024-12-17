// app/util/jwt.go
package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"strconv"
	"time"
)

// 用于对 JWT 进行签名的密钥
var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

type Claims struct {
	Username       string `json:"username"`
	StandardClaims jwt.StandardClaims
}

func (c *Claims) Valid() error {
	// 这里你可以根据业务需求自定义有效性验证逻辑
	// 比如检查令牌是否过期，签发者是否匹配等
	if c.StandardClaims.ExpiresAt < time.Now().Unix() {
		return fmt.Errorf("token expired")
	}

	// 如果需要进一步验证Issuer、Audience等字段，也可以在这里检查
	if c.StandardClaims.Issuer != "Login_Function_PY" {
		return fmt.Errorf("invalid issuer")
	}

	return nil
}

// GenerateToken 为给定用户名生成新的 JWT
func GenerateToken(username string) (string, error) {

	// 从环境变量中获取 JWT 令牌的过期时间
	expirationHours, _ := strconv.Atoi(os.Getenv("JWT_TOKEN_EXPIRE_TIME_HOURS"))

	// 计算令牌的过期时间
	expirationTime := time.Now().Add(time.Duration(expirationHours) * time.Hour)

	//创建 JWT 声明
	//claims := &jwt.StandardClaims{
	//	Subject:   username,              // User ID or username
	//	ExpiresAt: expirationTime.Unix(), // Expiry time
	//	IssuedAt:  time.Now().Unix(),     // Issue time
	//}

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // 设置令牌过期时间
			Issuer:    "Login_Function_PY",   // 设置令牌发行者
		},
	}

	// 使用 HS256 签名算法创建 JWT 令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用我们的密钥对令牌进行签名
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Println("Error signing token:", err)
		return "", err
	}

	return signedToken, nil
}

// ValidateToken 验证 JWT 令牌，并在有效时返回用户名
func ValidateToken(tokenString string) (string, error) {
	// 解析令牌
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Return the JWT secret key for validation
		return jwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims.Subject, nil // Return the username (subject)
	}

	return "", nil
}
