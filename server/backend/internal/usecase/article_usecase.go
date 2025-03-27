package usecase

import (
	"server/internal/domain/repository"
	"server/internal/domain/models"

	"math"
)

type ArticleUsecase struct{
	ArticleRepo repository.ArticleRepository
}

func NewArticleUsecase(ArticleRepository repository.ArticleRepository) *ArticleUsecase{
	return &ArticleUsecase{ArticleRepo: ArticleRepository}
}

func (a *ArticleUsecase) GetArticlePaginated(page, limit int) (*models.ArticlePagination, error){
	var articles []models.ArticleSelect
	var totalRows int64

	offset := (page - 1) * limit

	if err := a.ArticleRepo.Db.Model(&models.Article{}).Count(&totalRows).Error; err != nil{
		return nil, err
	}

	if err := a.ArticleRepo.Db.Table("articles").Offset(offset).Limit(limit).Select("name", "article").Find(&articles).Error; err != nil{
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