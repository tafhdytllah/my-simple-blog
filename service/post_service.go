package service

import (
	"my-simple-blog/dto"
	"my-simple-blog/entity"
	"my-simple-blog/errorhandler"
	"my-simple-blog/repository"
)

type PostService interface {
	CreateArticle(req *dto.PostRequest) error
	FindArticles() ([]entity.Post, error)
	FindArticleById(ID int) (entity.Post, error)
}

type postService struct {
	repository repository.PostRepository
}

func NewPostService(r repository.PostRepository) *postService {
	return &postService{
		repository: r,
	}
}

func (s *postService) FindArticleById(ID int) (entity.Post, error) {
	article, err := s.repository.FindArticleById(ID)

	if article.ID == 0 {
		return article, &errorhandler.NotFoundError{
			Message: "Article not found",
		}
	}

	return article, err
}

func (s *postService) FindArticles() ([]entity.Post, error) {
	articles, err := s.repository.FindArticles()

	return articles, err
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
