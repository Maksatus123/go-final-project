package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Maksatus123/go-final-project/internal/models"
	"github.com/Maksatus123/go-final-project/internal/service"
)

type BookController struct {
	svc *service.BookService
}

func NewBookController(svc *service.BookService) *BookController {
	return &BookController{svc: svc}
}

func (ctrl *BookController) CreateBook(c *gin.Context) {
	userID, _ := c.Get("userID")
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.svc.CreateBook(&book, userID.(int)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, book)
}

func (ctrl *BookController) GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := ctrl.svc.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (ctrl *BookController) GetAllBooks(c *gin.Context) {
	books, err := ctrl.svc.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (ctrl *BookController) UpdateBook(c *gin.Context) {
	userID, _ := c.Get("userID")
	id, _ := strconv.Atoi(c.Param("id"))
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.ID = id
	if err := ctrl.svc.UpdateBook(&book, userID.(int)); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (ctrl *BookController) DeleteBook(c *gin.Context) {
	userID, _ := c.Get("userID")
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.svc.DeleteBook(id, userID.(int)); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (ctrl *BookController) GetBooksByOwner(c *gin.Context) {
	ownerID, _ := strconv.Atoi(c.Query("owner_id"))
	books, err := ctrl.svc.GetBooksByOwner(ownerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}