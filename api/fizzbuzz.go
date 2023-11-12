package api

import (
	"errors"
	db "github.com/Wintersunner/xplor/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FizzBuzzResponse struct {
	ID        int64     `json:"id"`
	Message   string    `json:"message"`
	UserAgent string    `json:"useragent"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateFizzBuzzRequest struct {
	Message string `json:"message" binding:"required"`
}

func (server *Server) listFizzBuzzes(context *gin.Context) {
	list, err := server.store.ListFizzBuzzes(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	context.JSON(http.StatusOK, list)
}

func (server *Server) getFizzBuzz(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, errorResponse(errors.New("invalid id")))
		return
	}
	fizzBuzz, err := server.store.GetFizzBuzz(context, int64(id))

	if err != nil {
		context.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	context.JSON(http.StatusOK, FizzBuzzResponse{
		ID:        fizzBuzz.ID,
		Message:   fizzBuzz.Message,
		UserAgent: fizzBuzz.Useragent,
		CreatedAt: fizzBuzz.CreatedAt,
	})
}

func (server *Server) createFizzBuzz(context *gin.Context) {
	var req CreateFizzBuzzRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusUnprocessableEntity, errorResponse(err))
		return
	}
	result, err := server.store.CreateFizzBuzz(context, db.CreateFizzBuzzParams{
		Useragent: context.Request.UserAgent(),
		Message:   req.Message,
		CreatedAt: time.Now().UTC(),
	})

	if err != nil {
		context.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	id, err := result.LastInsertId()

	if err != nil {
		context.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	context.JSON(http.StatusOK, FizzBuzzResponse{
		ID:        id,
		Message:   req.Message,
		UserAgent: context.Request.UserAgent(),
		CreatedAt: time.Now().UTC(),
	})
}
