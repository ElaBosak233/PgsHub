package implements

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/elabosak233/pgshub/internal/middlewares"
	"github.com/elabosak233/pgshub/internal/models/misc"
	"github.com/elabosak233/pgshub/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

type AuthMiddlewareImpl struct {
	appService *services.AppService
}

func NewAuthMiddleware(appService *services.AppService) middlewares.AuthMiddleware {
	return &AuthMiddlewareImpl{
		appService: appService,
	}
}

func (m *AuthMiddlewareImpl) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		pgsToken, err := jwt.ParseWithClaims(token, &misc.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("jwt.secret_key")), nil
		})
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  err.Error(),
			})
			ctx.Abort()
			return
		}
		if claims, ok := pgsToken.Claims.(*misc.Claims); ok && pgsToken.Valid {
			userId := claims.UserId
			ctx.Set("UserId", userId)
			user, err := m.appService.UserService.FindById(userId)
			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"code": http.StatusUnauthorized,
					"msg":  "无效 Token",
				})
				ctx.Abort()
				return
			}
			ctx.Set("UserRole", user.Role)
			ctx.Next()
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "无效 Token",
			})
			ctx.Abort()
			return
		}
	}
}
