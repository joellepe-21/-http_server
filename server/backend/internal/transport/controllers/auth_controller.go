package controllers

import (
	"net/http"
	"server/internal/domain/models"
	"server/internal/domain/repository"
	"server/pkg"

	"github.com/gin-gonic/gin"
)

type AuthController struct{
	repo *repository.UserRepository
}

func NewAuthController(repo *repository.UserRepository) *AuthController{
	return &AuthController{repo}
}

func (c *AuthController) Register(ctx *gin.Context){
	var input struct{
		Login string `json:"login"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error":"Invalid input"})
		return
	}

	hashPassword, err := pkg.HashPassword(input.Password)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":"could not hash password"})
		return
	}

	user := models.User{Login: input.Login, Password: hashPassword}
	if err := c.repo.CreateUser(&user); err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Could not create user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message":"User created"})
}

func (c *AuthController) Authorization(ctx *gin.Context){
	var input struct{
		Login string `json:"login"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error":"Invalid input"})
		return
	}

	user, err := c.repo.GetUserByLogin(input.Login)
	if err != nil{
		ctx.JSON(http.StatusUnauthorized, gin.H{"error":"User not found"})
		return
	}

	if err := pkg.CheckPassword(user.Password, input.Password); err != nil{
		ctx.JSON(http.StatusUnauthorized, gin.H{"error":"Invalid password"})
		return
	}

	token, err := pkg.GenerateJWT(user.Login)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Could not generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}