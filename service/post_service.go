package service

import (
	"my-simple-blog/dto"
	"my-simple-blog/entity"
	"my-simple-blog/errorhandler"
	"my-simple-blog/repository"
	"strings"

	"gorm.io/gorm"
)

type PostService interface {
	CreateArticle(req *dto.PostRequest) error
	FindArticles() ([]entity.Post, error)
	FindArticleById(ID int) (entity.Post, error)
	FindArticleByTitle(title string) ([]entity.Post, error)
	UpdateArticle(ID int, req *dto.PostRequest) (entity.Post, error)
}

type postService struct {
	repository repository.PostRepository
}

func NewPostService(r repository.PostRepository) *postService {
	return &postService{
		repository: r,
	}
}

func (s *postService) UpdateArticle(ID int, req *dto.PostRequest) (entity.Post, error) {

	article, err := s.FindArticleById(ID)
	if err != nil {
		return entity.Post{}, &errorhandler.BadRequestError{
			Message: err.Error(),
		}
	}

	article = entity.Post{
		ID:      article.ID,
		UserID:  article.UserID,
		Title:   strings.TrimSpace(req.Title),
		Content: strings.TrimSpace(req.Content),
	}

	if req.Picture != nil {
		article.PictureUrl = &req.Picture.Filename
	}

	updatedArticle, err := s.repository.UpdateArticle(article)
	if err != nil {
		return entity.Post{}, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	return updatedArticle, nil
}

func (s *postService) FindArticleByTitle(title string) ([]entity.Post, error) {
	articles, err := s.repository.FindArticleByTitle(title)
	if err != nil {
		return []entity.Post{}, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	return articles, nil
}

func (s *postService) FindArticleById(ID int) (entity.Post, error) {
	article, err := s.repository.FindArticleById(ID)

	if err != nil {
		// article not found
		if err == gorm.ErrRecordNotFound {
			return entity.Post{}, &errorhandler.NotFoundError{
				Message: "Article not found",
			}
		}

		return entity.Post{}, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	return article, nil
}

func (s *postService) FindArticles() ([]entity.Post, error) {
	articles, err := s.repository.FindArticles()
	if err != nil {
		return []entity.Post{}, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	return articles, nil
}

func (s *postService) CreateArticle(req *dto.PostRequest) error {
	// set value post
	post := entity.Post{
		UserID:  req.UserID,
		Title:   req.Title,
		Content: req.Content,
	}

	// set picture full ext if picture exist
	if req.Picture != nil {
		post.PictureUrl = &req.Picture.Filename
	}

	if err := s.repository.CreateArticle(&post); err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	return nil
}
