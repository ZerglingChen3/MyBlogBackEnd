package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"myBlog/Response"
	"myBlog/common"
	"myBlog/dto"
	"myBlog/model"
	"myBlog/util"
	"net/http"
)

func Register(ctx *gin.Context) {
	db := common.GetDB()
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	if len(telephone) != 11 {
		Response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "Telephone Must Be 11!")
		return
	}

	if len(password) < 6 {
		Response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "Password Should Not Less Than 6!")
		return
	}

	if len(name) == 0 {
		name = util.RandomString(10)
	}

	if isTelephoneExists(db, telephone) {
		Response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "Telephone Exists!")
		return
	}

	// Create User
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		Response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "Encryption Error!")
		return
	}

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(encryptedPassword),
	}
	db.Create(&newUser)

	Response.Success(ctx, nil, "Register OK!")
}

func Login(ctx *gin.Context) {
	db := common.GetDB()

	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	if len(telephone) != 11 {
		Response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "Telephone Must Be 11!")
		return
	}

	if len(password) < 6 {
		Response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "Password Should Not Less Than 6!")
		return
	}

	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		Response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "User not exists!")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		Response.Response(ctx, http.StatusBadRequest, 400, nil, "Wrong Password!")
		return
	}

	token, err := common.ReleaseToken(user)
	if err != nil {
		Response.Response(ctx, http.StatusInternalServerError, 500, nil, "System Error!")
		log.Printf("token generate error : %v", err)
		return
	}

	Response.Success(ctx, gin.H{"token": token}, "Login Successfully!")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	Response.Success(ctx, gin.H{"user": dto.ToUserDto(user.(model.User))}, "")
}

func isTelephoneExists(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
