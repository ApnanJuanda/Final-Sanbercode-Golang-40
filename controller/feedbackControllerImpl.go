package controller

import (
	"fmt"
	"net/http"
	"project/model"
	"project/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FeedbackControllerImpl struct {
	feedbackService service.FeedbackService
}

func NewFeedbackController(feedbackService service.FeedbackService) FeedbackController{
	return &FeedbackControllerImpl{
		feedbackService: feedbackService,
	}
}

func (fc *FeedbackControllerImpl) AddFeedback(ctx *gin.Context) {
	var addFeedback model.AddFeedback
	if err := ctx.ShouldBindJSON(&addFeedback); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	CustomerId := ctx.Request.Context().Value("customer_id")
	feedback, err := fc.feedbackService.AddFeedback(&addFeedback, CustomerId.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, feedback)	
}

func (fc *FeedbackControllerImpl) GetFeedbacks(ctx *gin.Context) {
	feedbacks, err := fc.feedbackService.GetFeedbacks()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, feedbacks)
}

func (fc *FeedbackControllerImpl) GetFeedbackById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	feedback, err := fc.feedbackService.GetFeedbackById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, feedback)	
}

func (fc *FeedbackControllerImpl) UpdateFeedbackById(ctx *gin.Context) {
	var addFeedback model.AddFeedback
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	if err := ctx.ShouldBindJSON(&addFeedback); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	count, err := fc.feedbackService.UpdateFeedbackById(id, &addFeedback)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	message := fmt.Sprintf("Updated data amount %d", count)
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func (fc *FeedbackControllerImpl) DeleteFeedbackById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	count, err := fc.feedbackService.DeleteFeedbackById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	message := fmt.Sprintf("Deleted data amount %d", count)
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}