package service

import "project/model"

type FeedbackService interface {
	AddFeedback(addFeedback *model.AddFeedback, customerId int) (*model.Feedback, error)
	GetFeedbacks() (*[]model.Feedback, error)
	GetFeedbackById(id int) (*model.Feedback, error)
	UpdateFeedbackById(id int, addFeedback *model.AddFeedback) (int, error)
	DeleteFeedbackById(id int) (int, error)
}
