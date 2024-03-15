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

func (h *postHandler) Create(c *gin.Context) {
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

	// userID, _ := c.Get("userID")
	post.UserID = 1

	//pass data to create service
	if err := h.service.Create(&post); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Sucess create article",
	})

	c.JSON(http.StatusCreated, res)
}
