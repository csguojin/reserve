package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"

	"github.com/csguojin/reserve/util"
	"github.com/csguojin/reserve/util/logger"
)

func AuthUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := checkAuth(c, viper.GetString("jwt.userkey"))
		if err != nil || claims == nil {
			logger.L.Errorln(err)
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		userid, ok := claims["useid"].(int)
		if !ok {
			logger.L.Errorln("missing user id in token")
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		username, ok := claims["username"].(string)
		if !ok {
			logger.L.Errorln("missing user name in token")
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Set("userid", userid)
		c.Set("username", username)
		c.Next()
	}
}

func AuthAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := checkAuth(c, viper.GetString("jwt.adminkey"))
		if err != nil || claims == nil {
			logger.L.Errorln(err)
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		userid, ok := claims["adminid"].(int)
		if !ok {
			logger.L.Errorln("missing admin id in token")
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		username, ok := claims["adminname"].(string)
		if !ok {
			logger.L.Errorln("missing admin name in token")
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Set("adminid", userid)
		c.Set("adminname", username)
		c.Next()
	}
}

func checkAuth(c *gin.Context, key string) (jwt.MapClaims, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, errors.New("missing Authorization header")
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		return nil, util.ErrTokenInvalid
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, util.ErrTokenInvalid
	}

	expTime, ok := claims["exp"].(time.Time)
	if !ok {
		return nil, errors.New("missing expired time in token")
	}
	if expTime.Before(time.Now()) {
		return nil, errors.New("token is expired")
	}

	return claims, nil
}
