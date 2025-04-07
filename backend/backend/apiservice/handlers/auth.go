package handlers

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

var rsaPublicKey *rsa.PublicKey
var enforcer *casbin.Enforcer

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Failed to load .env file")
	}

	key := os.Getenv("publickey")
	if key == "" {
		panic("publickey not found in .env")
	}

	pemKey := fmt.Sprintf("-----BEGIN PUBLIC KEY-----\n%s\n-----END PUBLIC KEY-----", key)

	block, _ := pem.Decode([]byte(pemKey))
	if block == nil {
		panic("failed to decode PEM block")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic("failed to parse public key: " + err.Error())
	}

	var ok bool
	rsaPublicKey, ok = pub.(*rsa.PublicKey)
	if !ok {
		panic("key is not a valid RSA public key")
	}

	enforcer, err = casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		panic("failed to create casbin enforcer: " + err.Error())
	}
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return rsaPublicKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}
		if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
			c.JSON(401, gin.H{"error": "Token expired"})
			c.Abort()
			return
		}
		if !claims.VerifyIssuer("http://localhost:8080/realms/authgym", true) {
			c.JSON(401, gin.H{"error": "Invalid token issuer"})
			c.Abort()
			return
		}

		username, ok := claims["preferred_username"].(string)
		if !ok {
			c.JSON(401, gin.H{"error": "Username not found in token"})
			c.Abort()
			return
		}

		realmAccess, ok := claims["realm_access"].(map[string]interface{})
		if !ok {
			c.JSON(401, gin.H{"error": "Roles not found in token"})
			c.Abort()
			return
		}
		rawRoles, ok := realmAccess["roles"].([]interface{})
		if !ok || len(rawRoles) == 0 {
			c.JSON(401, gin.H{"error": "No roles found in token"})
			c.Abort()
			return
		}

		var rolesList []string
		for _, r := range rawRoles {
			if roleStr, ok := r.(string); ok {
				rolesList = append(rolesList, roleStr)
			}
		}

		resource := c.Request.URL.Path
		action := c.Request.Method
		allowed := false
		for _, role := range rolesList {
			permit, err := enforcer.Enforce(role, resource, action)
			if err != nil {
				c.JSON(500, gin.H{"error": "Error checking permission"})
				c.Abort()
				return
			}
			if permit {
				allowed = true
				break
			}
		}
		if !allowed {
			c.JSON(403, gin.H{"error": "Forbidden: Insufficient permissions"})
			c.Abort()
			return
		}

		c.Set("username", username)
		c.Set("roles", rolesList)
		c.Next()
	}
}
