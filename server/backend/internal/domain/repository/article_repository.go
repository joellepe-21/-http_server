package repository

import (
	"server/internal/domain/models"
	//"time"

	"math"

	"gorm.io/gorm"
)

type ArticleRepository struct{
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository{
	return &ArticleRepository{db}
}

func (a *ArticleRepository) CreateArticle(article *models.Article) error{
	return a.db.Create(article).Error
}

func (a *ArticleRepository) FindArticle(nameArticle string) (*models.Article, error){
	var article models.Article
	if err := a.db.Where("name=?", nameArticle).First(&article).Error; err != nil{
		return nil, err
	}
	return &article, nil
}

func (a *ArticleRepository) GetAllArticle() (*[]models.ArticleSelect, error){
	var articles []models.ArticleSelect
	if err := a.db.Model(&models.Article{}).
		Select("name", "article").
		Find(&articles).Error; err != nil{
		return nil, err
	}
	return &articles, nil
}

func (a *ArticleRepository) GetArticlePaginated(page, limit int) (*models.ArticlePagination, error){
	//time.Sleep(4 * time.Second)
	var articles []models.ArticleSelect
	var totalRows int64

	offset := (page - 1) * limit

	if err := a.db.Model(&models.Article{}).Count(&totalRows).Error; err != nil{
		return nil, err
	}

	if err := a.db.Table("articles").Offset(offset).Limit(limit).Select("name", "article").Find(&articles).Error; err != nil{
		return nil, err
	}

	totalPages := int(math.Ceil(float64(totalRows) / float64(limit)))

	return &models.ArticlePagination{
		Data: articles,
		TotalRows: totalRows,
		TotalPage: totalPages,
		Page: page,
		Limit: limit,
	}, nil
}

func (a *ArticleRepository) GetRandThreeArticles() (*[]models.ArticleSelect, error){
	var articles []models.ArticleSelect
	if err := a.db.Model(&models.Article{}).
		Order("RANDOM()").
		Limit(3).
		Select("name", "article").
		Find(&articles).Error; err != nil{
		return nil, err
	}
	return &articles, nil
}

func (a *ArticleRepository) DeleteArticle(nameArticle string) error{
	var article models.Article
	return a.db.Where("name=?",nameArticle).Delete(&article).Error
}

func (a *ArticleRepository) UpdateArtecle(nameArticle, article_text string) error{
	return a.db.Model(&models.Article{}).Where("name=?", nameArticle).Update("article", article_text).Error
}