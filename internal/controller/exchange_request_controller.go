package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Maksatus123/go-final-project/internal/models"
	"github.com/Maksatus123/go-final-project/internal/service"
)

type ExchangeRequestController struct {
	svc *service.ExchangeRequestService
}

func NewExchangeRequestController(svc *service.ExchangeRequestService) *ExchangeRequestController {
	return &ExchangeRequestController{svc: svc}
}

func (ctrl *ExchangeRequestController) CreateExchangeRequest(c *gin.Context) {
	userID, _ := c.Get("userID")
	var req models.ExchangeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.svc.CreateExchangeRequest(&req, userID.(int)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, req)
}

func (ctrl *ExchangeRequestController) GetExchangeRequest(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	req, err := ctrl.svc.GetExchangeRequestByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

func (ctrl *ExchangeRequestController) GetExchangeRequestsByRequester(c *gin.Context) {
	userID, _ := c.Get("userID")
	reqs, err := ctrl.svc.GetExchangeRequestsByRequester(userID.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reqs)
}

func (ctrl *ExchangeRequestController) UpdateExchangeRequestStatus(c *gin.Context) {
	userID, _ := c.Get("userID")
	id, _ := strconv.Atoi(c.Param("id"))
	var input struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.svc.UpdateExchangeRequestStatus(id, input.Status, userID.(int)); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}