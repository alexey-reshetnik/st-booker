package server

import (
	"net/http"
	"strconv"

	"st-booker/internal/model"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateBooking(ctx *gin.Context) {
	var b model.Booking
	if err := ctx.ShouldBindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := s.service.CreateBooking(ctx, b)
	if err != nil {
		s.log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, id)
}

func (s *Server) GetBookings(ctx *gin.Context) {
	var (
		l = ctx.Query("limit")
		o = ctx.Query("offset")
	)

	if l == "" || o == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "limit/offset params should not be empty"})
		return
	}

	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse limit"})
		return
	}

	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse offset"})
		return
	}

	bk, err := s.service.GetBookings(ctx, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, bk)
}

func (s *Server) DeleteBooking(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "\"id\" param can't be empty"})
		return
	}

	err := s.service.DeleteBooking(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
