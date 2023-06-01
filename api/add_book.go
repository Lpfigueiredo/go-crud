package api

import (
	"log"
	"lpfigueiredo/go-crud/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) HandleAddBook(ctx *gin.Context) {
	var book pkg.Book

	err := ctx.BindJSON(&book)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())

		return
	}

	book.Id = uuid.New()

	r := s.DB.Create(&book)

	if r.Error != nil {
		log.Fatal(r.Error)
	}

	s.DB.Save(&book)

	ctx.JSON(http.StatusOK, &book)
}
