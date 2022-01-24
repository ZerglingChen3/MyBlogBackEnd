package middleware

import (
	"github.com/gin-gonic/gin"
	"myBlog/common"
	"myBlog/model"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Permission Error!"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Permission Error!"})
			ctx.Abort()
			return
		}

		userId := claims.UserId
		db := common.GetDB()
		var user model.User
		db.First(&user, userId)

		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "Permission Error!"})
			ctx.Abort()
			return
		}

		ctx.Set("user", user)

		ctx.Next()
	}
}
