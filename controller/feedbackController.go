package controller

import "github.com/gin-gonic/gin"

type FeedbackController interface {
	AddFeedback(ctx *gin.Context)
	GetFeedbacks(ctx *gin.Context)
	GetFeedbackById(ctx *gin.Context)
	UpdateFeedbackById(ctx *gin.Context)
	DeleteFeedbackById(ctx *gin.Context)
}