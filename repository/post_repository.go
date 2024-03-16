package repository

import (
	"my-simple-blog/entity"

	"gorm.io/gorm"
)

type PostRepository interface {
	CreateArticle(post *entity.Post) error
	FindArticles() ([]entity.Post, error)
	FindArticleById(ID int) (entity.Post, error)
	FindArticleByTitle(title string) ([]entity.Post, error)
	UpdateArticle(post entity.Post) (entity.Post, error)
	DeleteArticle(post entity.Post) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{
		db: db,
	}
}

// delete record by id
func (r *postRepository) DeleteArticle(post entity.Post) error {

	err := r.db.Delete(&post).Error

	return err
}

// update record by id
func (r *postRepository) UpdateArticle(post entity.Post) (entity.Post, error) {

	err := r.db.Save(&post).Error

	return post, err
}

// get record with string condition
func (r *postRepository) FindArticleByTitle(title string) ([]entity.Post, error) {
	var articles []entity.Post

	// SELECT * FROM posts WHERE title LIKE 'title_value';
	err := r.db.Where("title = ?", title).Find(&articles).Error

	return articles, err
}

// get record with primary key
func (r *postRepository) FindArticleById(ID int) (entity.Post, error) {
	var article entity.Post

	// SELECT * FROM posts WHERE id = 10;
	err := r.db.First(&article, ID).Error

	return article, err
}

// get all record
func (r *postRepository) FindArticles() ([]entity.Post, error) {
	var articles []entity.Post

	// SELECT * FROM posts;
	err := r.db.Find(&articles).Error

	return articles, err
}

// insert record
func (r *postRepository) CreateArticle(post *entity.Post) error {

	// INSERT INTO posts (column_name1, column_name2) VALUES (value1, value2);
	result := r.db.Create(&post)

	return result.Error
}
