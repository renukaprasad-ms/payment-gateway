package middleware

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"net/http"
	"os"
	"payment-gateway/internal/config"
	"payment-gateway/internal/modules/auth"
	"payment-gateway/internal/response"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	publicKeyPath := config.LoadConfig().JWTPublicKeyPath
	pubKey, err := loadPublicKey(publicKeyPath)
	if err != nil {
		log.Fatalf("failed to load JWT public key from %s: %v", publicKeyPath, err)
	}

	return func(c *gin.Context) {

		token, err := c.Cookie(config.LoadConfig().AccessTokenCookieName)
		if err != nil {
			response.AbortError(c, http.StatusUnauthorized, "missing token")
			return
		}

		claims, err := auth.VerifyToken(token, pubKey)
		if err != nil {
			response.AbortError(c, http.StatusUnauthorized, "invalid token")
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("roles", claims.Roles)

		c.Next()
	}
}

func loadPublicKey(path string) (*rsa.PublicKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("invalid PEM data")
	}

	parsedKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err == nil {
		publicKey, ok := parsedKey.(*rsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf("PEM public key is not RSA")
		}

		return publicKey, nil
	}

	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}
