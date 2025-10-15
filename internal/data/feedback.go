package data

import (
	"database/sql"
	"time"
)

type FacilitatorFeedback struct {
	ID                  int64     `json:"id"`
	FacilitatorID       int64     `json:"facilitator_id"`
	SessionEnrollmentID int64     `json:"session_enrollment_id"`
	Score               int       `json:"score"`
	Comment             string    `json:"comment"`
	CreatedAt           time.Time `json:"created_at"`
}

type CourseFeedback struct {
	ID                  int64     `json:"id"`
	CourseID            int64     `json:"course_id"`
	SessionEnrollmentID int64     `json:"session_enrollment_id"`
	Score               int       `json:"score"`
	Comment             string    `json:"comment"`
	CreatedAt           time.Time `json:"created_at"`
}

type FeedbackModel struct {
	DB *sql.DB
}

func (m FeedbackModel) InsertFacilitatorFeedback(feedback *FacilitatorFeedback) error {
	query := `
		INSERT INTO facilitator_ratings (facilitator_id, session_enrollment_id, score, comment)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at`

	args := []interface{}{feedback.FacilitatorID, feedback.SessionEnrollmentID, feedback.Score, feedback.Comment}

	return m.DB.QueryRow(query, args...).Scan(&feedback.ID, &feedback.CreatedAt)
}

func (m FeedbackModel) GetAllForFacilitator(facilitatorID int64) ([]*FacilitatorFeedback, error) {
	query := `
		SELECT id, facilitator_id, session_enrollment_id, score, comment, created_at
		FROM facilitator_ratings
		WHERE facilitator_id = $1
		ORDER BY id`

	rows, err := m.DB.Query(query, facilitatorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	feedbacks := []*FacilitatorFeedback{}

	for rows.Next() {
		var feedback FacilitatorFeedback
		err := rows.Scan(
			&feedback.ID,
			&feedback.FacilitatorID,
			&feedback.SessionEnrollmentID,
			&feedback.Score,
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
		INSERT INTO course_ratings (session_enrollment_id, score, comment)
		VALUES ($1, $2, $3)
		RETURNING id, created_at`

	args := []interface{}{feedback.SessionEnrollmentID, feedback.Score, feedback.Comment}

	err := m.DB.QueryRow(query, args...).Scan(&feedback.ID, &feedback.CreatedAt)
	if err != nil {
		return err
	}

	// Get the course_id from the training_sessions table
	query = `
		SELECT ts.course_id
		FROM training_sessions ts
		JOIN session_enrollment se ON ts.id = se.session_id
		WHERE se.id = $1`

	return m.DB.QueryRow(query, feedback.SessionEnrollmentID).Scan(&feedback.CourseID)
}

func (m FeedbackModel) GetAllForCourse(courseID int64) ([]*CourseFeedback, error) {
	query := `
		SELECT cr.id, ts.course_id, cr.session_enrollment_id, cr.score, cr.comment, cr.created_at
		FROM course_ratings cr
		JOIN session_enrollment se ON cr.session_enrollment_id = se.id
		JOIN training_sessions ts ON se.session_id = ts.id
		WHERE ts.course_id = $1
		ORDER BY cr.id`

	rows, err := m.DB.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	feedbacks := []*CourseFeedback{}

	for rows.Next() {
		var feedback CourseFeedback
		err := rows.Scan(
			&feedback.ID,
			&feedback.CourseID,
			&feedback.SessionEnrollmentID,
			&feedback.Score,
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
