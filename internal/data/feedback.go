package data

import (
	"database/sql"
	"time"
)

type FacilitatorFeedback struct {
	ID            int64     `json:"id"`
	FacilitatorID int64     `json:"facilitator_id"`
	UserID        int64     `json:"user_id"`
	Rating        int       `json:"rating"`
	Comment       string    `json:"comment"`
	CreatedAt     time.Time `json:"created_at"`
}

type CourseFeedback struct {
	ID        int64     `json:"id"`
	CourseID  int64     `json:"course_id"`
	UserID    int64     `json:"user_id"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

type FeedbackModel struct {
	DB *sql.DB
}

func (m FeedbackModel) InsertFacilitatorFeedback(feedback *FacilitatorFeedback) error {
	query := `
		INSERT INTO facilitator_ratings (facilitator_id, user_id, rating, comment)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at`

	args := []interface{}{feedback.FacilitatorID, feedback.UserID, feedback.Rating, feedback.Comment}

	return m.DB.QueryRow(query, args...).Scan(&feedback.ID, &feedback.CreatedAt)
}

func (m FeedbackModel) GetAllForFacilitator(facilitatorID int64) ([]*FacilitatorFeedback, error) {
	query := `
		SELECT id, facilitator_id, user_id, rating, comment, created_at
		FROM facilitator_ratings
		WHERE facilitator_id = $1
		ORDER BY id`

	rows, err := m.DB.Query(query, facilitatorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var feedbacks []*FacilitatorFeedback

	for rows.Next() {
		var feedback FacilitatorFeedback
		err := rows.Scan(
			&feedback.ID,
			&feedback.FacilitatorID,
			&feedback.UserID,
			&feedback.Rating,
			&feedback.Comment,
			&feedback.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		feedbacks = append(feedbacks, &feedback)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return feedbacks, nil
}

func (m FeedbackModel) InsertCourseFeedback(feedback *CourseFeedback) error {
	query := `
		INSERT INTO course_ratings (course_id, user_id, rating, comment)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at`

	args := []interface{}{feedback.CourseID, feedback.UserID, feedback.Rating, feedback.Comment}

	return m.DB.QueryRow(query, args...).Scan(&feedback.ID, &feedback.CreatedAt)
}

func (m FeedbackModel) GetAllForCourse(courseID int64) ([]*CourseFeedback, error) {
	query := `
		SELECT id, course_id, user_id, rating, comment, created_at
		FROM course_ratings
		WHERE course_id = $1
		ORDER BY id`

	rows, err := m.DB.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var feedbacks []*CourseFeedback

	for rows.Next() {
		var feedback CourseFeedback
		err := rows.Scan(
			&feedback.ID,
			&feedback.CourseID,
			&feedback.UserID,
			&feedback.Rating,
			&feedback.Comment,
			&feedback.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		feedbacks = append(feedbacks, &feedback)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return feedbacks, nil
}
