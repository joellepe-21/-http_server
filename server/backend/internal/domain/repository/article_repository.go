package repository

import (
	"server/internal/domain/models"

	"gorm.io/gorm"
)

type ArticleRepository struct{
	Db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository{
	return &ArticleRepository{db}
}

func (a *ArticleRepository) CreateArticle(article *models.Article) error{
	return a.Db.Create(article).Error
}

func (a *ArticleRepository) FindArticle(nameArticle string) (*models.Article, error){
	var article models.Article
	if err := a.Db.Where("name=?", nameArticle).First(&article).Error; err != nil{
		return nil, err
	}
	return &article, nil
}

func (a *ArticleRepository) GetAllArticle() (*[]models.ArticleSelect, error){
	var articles []models.ArticleSelect
	if err := a.Db.Model(&models.Article{}).
		Select("name", "article").
		Find(&articles).Error; err != nil{
		return nil, err
	}
	return &articles, nil
}

func (a *ArticleRepository) DeleteArticle(nameArticle string) error{
	var article models.Article
	return a.Db.Where("name=?",nameArticle).Delete(&article).Error
}

func (a *ArticleRepository) UpdateArtecle(nameArticle, article_text string) error{
	return a.Db.Model(&models.Article{}).Where("name=?", nameArticle).Update("article", article_text).Error
}