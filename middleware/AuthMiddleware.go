package middleware

import (
	"github.com/gin-gonic/gin"
	"myBlog/Response"
	"myBlog/common"
	"myBlog/model"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			Response.Response(ctx, http.StatusUnauthorized, 401, nil, "Permission Error!")
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)

		if err != nil || !token.Valid {
			Response.Response(ctx, http.StatusUnauthorized, 401, nil, "Permission Error!")
			ctx.Abort()
			return
		}

		userId := claims.UserId
		db := common.GetDB()
		var user model.User
		db.First(&user, userId)

		if user.ID == 0 {
			Response.Response(ctx, http.StatusUnauthorized, 401, nil, "Permission Error!")
			ctx.Abort()
			return
		}

		ctx.Set("user", user)

		ctx.Next()
	}
}
