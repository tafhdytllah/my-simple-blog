package repository

import (
	"fmt"
	"my-simple-blog/entity"

	"gorm.io/gorm"
)

type PostRepository interface {
	CreateArticle(post *entity.Post) error
	FindArticles() ([]entity.Post, error)
	FindArticleById(ID int) (entity.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) FindArticleById(ID int) (entity.Post, error) {
	var article entity.Post

	err := r.db.Find(&article, "id = ?", ID).Error

	fmt.Println(article)

	return article, err
}

func (r *postRepository) FindArticles() ([]entity.Post, error) {
	var articles []entity.Post

	err := r.db.Find(&articles).Error

	return articles, err
}

func (r *postRepository) CreateArticle(post *entity.Post) error {
	result := r.db.Create(&post)
	return result.Error
}
