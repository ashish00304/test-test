package controllers

import (
	"net/http"
	"strconv"

	"github.com/ashish00304/test-test/model"
	"github.com/ashish00304/test-test/services"
	"github.com/gin-gonic/gin"
)

func PrintHello(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, gin.H{
		"massage": 4 + 7,
	})

}

func AddUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	services.AddUserDetails(user)
	ctx.JSON(http.StatusCreated, gin.H{
		"user":    user,
		"massage": "user created successfully",
	})

}

func GetAllUser(ctx *gin.Context) {
	users := services.GetAllUserDetails()
	ctx.JSON(http.StatusAccepted, gin.H{
		"users": users,
	})
}

func DeleteUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	if services.DeleteUserDetails(id) {
		ctx.JSON(http.StatusAccepted, gin.H{
			"massage": "User deleted successfully",
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"massage": "User not found",
		})
	}
}

func UpdateUserDetails(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"massage": "Invalide user",
		})
		return
	}

	var updatedUser model.User
	if err := ctx.ShouldBindJSON(&updatedUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"massage": err.Error(),
		})
		return
	}

	user, ok := services.UpdateUserDetails(id, updatedUser)

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"massage": "user not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"massage":     "user upadated successfully",
		"updatedUser": user,
	})

}

func Login(ctx *gin.Context) {
	userName := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if userName == "admin" && password == "1234" {
		ctx.JSON(http.StatusOK, gin.H{
			"massage": "You are welcome in Admin Portal",
		})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"massage": "Invalid credentials",
	})

}

func SecureProfile(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"massage": "Welcome ! You are authorized",
	})
}
