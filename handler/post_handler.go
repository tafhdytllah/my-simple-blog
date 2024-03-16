package handler

import (
	"fmt"
	"my-simple-blog/dto"
	"my-simple-blog/errorhandler"
	"my-simple-blog/helper"
	"my-simple-blog/service"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type postHandler struct {
	service service.PostService
}

func NewPostHandler(s service.PostService) *postHandler {
	return &postHandler{
		service: s,
	}
}

// Get Article
func (h *postHandler) GetArticle(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	result, err := h.service.FindArticleById(ID)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Success get article",
		Data:       result,
	})

	c.JSON(http.StatusOK, res)
}

// Get All Article
func (h *postHandler) GetArticles(c *gin.Context) {
	result, err := h.service.FindArticles()
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Success get articles",
		Data:       result,
	})

	c.JSON(http.StatusOK, res)

}

// Create New Article
func (h *postHandler) CreateArticle(c *gin.Context) {
	var post dto.PostRequest

	// bind form data request to type post
	if err := c.ShouldBind(&post); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{
			Message: err.Error(),
		})
		return
	}

	// if form image data is exixst save to local
	if post.Picture != nil {
		// mkdir
		if err := os.MkdirAll("public/picture", 0755); err != nil {
			errorhandler.HandleError(c, &errorhandler.InternalServerError{
				Message: err.Error(),
			})
			return
		}

		// rename file
		ext := filepath.Ext(post.Picture.Filename)
		newFileName := uuid.New().String() + ext

		// save image to local dir
		dst := filepath.Join("public/picture", filepath.Base(newFileName))
		c.SaveUploadedFile(post.Picture, dst)

		post.Picture.Filename = fmt.Sprintf("%s/public/picture/%s", c.Request.Host, newFileName)
	}

	userID, _ := c.Get("userID")
	post.UserID = userID.(int)

	//pass data to create service
	if err := h.service.CreateArticle(&post); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Sucess create article",
	})

	c.JSON(http.StatusCreated, res)
}
