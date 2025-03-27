package controllers

import (
	"net/http"
	"server/internal/domain/models"
	"server/internal/domain/repository"
	"server/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ArticleController struct{
	repo *repository.ArticleRepository
	usecase *usecase.ArticleUsecase
}

func NewArticleController(repo *repository.ArticleRepository, usecase *usecase.ArticleUsecase) *ArticleController{
	return &ArticleController{
		repo: repo,
		usecase: usecase,
	}
}

func (a *ArticleController) Add(ctx *gin.Context){
	var input struct{
		Name string `json:"name"`
		Text string `json:"article"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error":"Invalid input"})
		return
	}

	article := models.Article{Name: input.Name, Article: input.Text}

	if err :=  a.repo.CreateArticle(&article); err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Could not add article"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message":"article added"})
} 


func (a *ArticleController) Find(ctx *gin.Context){
	var input struct{
		Name string `json:"name"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error":"Invalid Input"})
		return
	}

	article, err := a.repo.FindArticle(input.Name)
	if err != nil{
		ctx.JSON(http.StatusNotFound, gin.H{"error":"article not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"article": article.Article})
}

func (a *ArticleController) GetAllArticles(ctx *gin.Context){
	articles, err := a.repo.GetAllArticle()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":"error receiving articles"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"all articles": articles})
}

func (a *ArticleController) Delete(ctx *gin.Context){
	var input struct{
		Name string `json:"name"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error":"Invalid input"})
	}

	if err := a.repo.DeleteArticle(input.Name); err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Error while deleting articals"})
	}

	ctx.JSON(http.StatusOK, gin.H{"message":"Article was deleted"})
}

func (a *ArticleController) Update(ctx *gin.Context){
	var input struct{
		Name string `json:"name"`
		Article string `json:"article"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error":"Invalid input"})
		return
	}

	if err := a.repo.UpdateArtecle(input.Name, input.Article); err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to update data"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message":"Data updated"})
}

func (a *ArticleController) ArticlePagination(ctx *gin.Context){
	var input struct{
		Page int `json:"page"`
		Limit int `json:"limit"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error":"Invalid Input"})
		return
	}
	
	articles, err := a.usecase.GetArticlePaginated(input.Page, input.Limit)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":"error receiving articles"})
		return
	}

	ctx.JSON(http.StatusOK, articles)
}	