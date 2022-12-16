package service

import (
	"database/sql"
	"errors"
	"fmt"
	"project/helper"
	"project/model"
)

type FeedbackServiceImpl struct {
	DB *sql.DB
}

func NewFeedbackService(DB *sql.DB) FeedbackService {
	return &FeedbackServiceImpl{
		DB: DB,
	}
}

func (fs *FeedbackServiceImpl) AddFeedback(addFeedback *model.AddFeedback, customerId int) (*model.Feedback, error) {
	//validasi product_id
	productService := NewProductService(fs.DB)
	product, _ := productService.GetProductById(addFeedback.ProductId)
	if product == nil {
		return nil, errors.New("product_id is not valid")
	}

	//populator
	feedback, err := helper.FeedbackPopulator(addFeedback, customerId)
	if err != nil {
		return nil, err
	}

	//save query
	var newFeedback = model.Feedback{}
	sqlStatement = `INSERT INTO feedback (customer_id, product_id, review)
	VALUES ($1, $2, $3)
	Returning *
	`
	err = fs.DB.QueryRow(sqlStatement, feedback.CustomerId, feedback.ProductId, feedback.Review).
		Scan(&newFeedback.ID, &newFeedback.CustomerId, &newFeedback.ProductId, &newFeedback.Review, &newFeedback.CreatedAt, &newFeedback.UpdatedAt)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return &newFeedback, nil

}

func (fs *FeedbackServiceImpl) GetFeedbacks() (*[]model.Feedback, error) {
	var feedbacks = []model.Feedback{}

	sqlStatement = `SELECT * FROM feedback`

	rows, err := fs.DB.Query(sqlStatement)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var feedback = model.Feedback{}

		err = rows.Scan(&feedback.ID, &feedback.CustomerId, &feedback.ProductId, &feedback.Review, &feedback.CreatedAt, &feedback.UpdatedAt)


		if err != nil {
			return nil, err
		}

		feedbacks = append(feedbacks, feedback)

	}
	return &feedbacks, nil
}

func (fs *FeedbackServiceImpl) GetFeedbackById(id int) (*model.Feedback, error) {
	var feedback = model.Feedback{}
	sqlStatement := `SELECT * FROM feedback WHERE id=($1)`
	err = fs.DB.QueryRow(sqlStatement, id).Scan(&feedback.ID, &feedback.CustomerId, &feedback.ProductId, &feedback.Review, &feedback.CreatedAt, &feedback.UpdatedAt)
	if err != nil {
		return nil, err		
	}
	return &feedback, err
}

func (fs *FeedbackServiceImpl) UpdateFeedbackById(id int, addFeedback *model.AddFeedback) (int, error) {
	if addFeedback.Review == "" {
		return 0, errors.New("invalid request body")
	}

	sqlStatement = `UPDATE feedback SET review=$2 WHERE id=$1;`
	
	result, err := fs.DB.Exec(sqlStatement, id, addFeedback.Review)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while updating FEEDBACK record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := result.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after update data: %v", err)
		return 0, errors.New(e) 
	}

	if count == 0 {
		e := "could not update the feedback, please try again after sometime"
		return 0, errors.New(e) 
	}
	return int(count), nil
}

func (fs *FeedbackServiceImpl) DeleteFeedbackById(id int) (int, error) {
	sqlStatement = `DELETE FROM feedback WHERE id=$1;`
	res, err := fs.DB.Exec(sqlStatement, id)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while delete feedback record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := res.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after delete data: %v", err)
		return 0, errors.New(e)		
	}

	if count == 0 {
		e := fmt.Sprintf("could not delete the feedback, there might be no data for ID %d", id)
		return 0, errors.New(e) 
	}
	return int(count), nil
}
