package repository

import (
	"my-simple-blog/entity"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *entity.Post) error
	FindArticles() ([]entity.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) FindArticles() ([]entity.Post, error) {
	var articles []entity.Post

	err := r.db.Find(&articles).Error

	return articles, err
}

func (r *postRepository) Create(post *entity.Post) error {
	result := r.db.Create(&post)
	return result.Error
}
