package controllers

import (
    "net/http"
    "server/internal/domain/models"
    "server/internal/domain/repository"
    "server/internal/usecase"

    "github.com/gin-gonic/gin"
)

// ArticleController представляет контроллер для работы со статьями.
type ArticleController struct {
    repo    *repository.ArticleRepository
    usecase *usecase.ArticleUsecase
}

// NewArticleController создает новый экземпляр ArticleController.
func NewArticleController(repo *repository.ArticleRepository, usecase *usecase.ArticleUsecase) *ArticleController {
    return &ArticleController{
        repo:    repo,
        usecase: usecase,
    }
}

// Add godoc
// @Summary      Add a new article
// @Description  Add a new article with the provided name and text.
// @Tags         articles
// @Accept       json
// @Produce      json
// @Param        article  body      object true  "Article to add"
// @Success      201      {object}  map[string]string
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Security     BearerAuth
// @Router       /api/add [post]
func (a *ArticleController) Add(ctx *gin.Context) {
    var input struct {
        Name string `json:"name"`
        Text string `json:"article"`
    }

    if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    article := models.Article{Name: input.Name, Article: input.Text}

    if err := a.repo.CreateArticle(&article); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add article"})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{"message": "article added"})
}

func (a *ArticleController) Find(ctx *gin.Context) {
    var input struct {
        Name string `json:"name"`
    }

    if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
        return
    }

    article, err := a.repo.FindArticle(input.Name)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "article not found"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"article": article.Article})
}

func (a *ArticleController) GetAllArticles(ctx *gin.Context) {
    articles, err := a.repo.GetAllArticle()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error receiving articles"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"all articles": articles})
}

// Delete godoc
// @Summary      Delete an article by name
// @Description  Delete an article by its unique name.
// @Tags         articles
// @Accept       json
// @Produce      json
// @Param        name  body      object true  "Article name"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security     BearerAuth
// @Router       /api/delete [delete]
func (a *ArticleController) Delete(ctx *gin.Context) {
    var input struct {
        Name string `json:"name"`
    }
    if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    if err := a.repo.DeleteArticle(input.Name); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while deleting articles"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Article was deleted"})
}

// Update godoc
// @Summary      Update an article
// @Description  Update an existing article with the provided details.
// @Tags         articles
// @Accept       json
// @Produce      json
// @Param        article  body      object  true  "Updated article"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Security     BearerAuth
// @Router       /api/update [put]
func (a *ArticleController) Update(ctx *gin.Context) {
    var input struct {
        Name    string `json:"name"`
        Article string `json:"article"`
    }
    if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    if err := a.repo.UpdateArtecle(input.Name, input.Article); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update data"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Data updated"})
}

// ArticlePagination godoc
// @Summary      Get paginated articles
// @Description  Retrieve a paginated list of articles.
// @Tags         articles
// @Accept       json
// @Produce      json
// @Param        pagination  body      object  true  "Pagination parameters"
// @Success      200         {object}  models.ArticlePagination
// @Failure      400         {object}  map[string]string
// @Failure      500         {object}  map[string]string
// @Router       /article [post]
func (a *ArticleController) ArticlePagination(ctx *gin.Context) {
    var input struct {
        Page  int `json:"page"`
        Limit int `json:"limit"`
    }

    if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
        return
    }

    articles, err := a.usecase.GetArticlePaginated(input.Page, input.Limit)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error receiving articles"})
        return
    }

    ctx.JSON(http.StatusOK, articles)
}