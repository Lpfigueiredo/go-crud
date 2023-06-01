package api

import (
	"log"
	"lpfigueiredo/go-crud/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) HandleGetByISBN(ctx *gin.Context) {
	var book pkg.Book

	isbn := ctx.Param("isbn")

	ret := s.DB.First(&book, "isbn = ?", isbn)

	if ret.RowsAffected == 0 {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	if ret.Error != nil {
		err := ctx.AbortWithError(http.StatusBadRequest, ret.Error)

		if err != nil {
			log.Fatal(err)
		}

		return
	}

	ctx.JSON(http.StatusOK, book)
}
