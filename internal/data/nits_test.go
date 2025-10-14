
package data

import (
	"testing"

	"github.com/Syha-01/national-inservice-training/internal/validator"
)

func TestValidateCourse(t *testing.T) {
	// Create a new Course instance for testing
	course := &Course{
		Title:       "Defensive Driving",
		Description: "A course on defensive driving techniques.",
		Category:    "Mandatory",
		CreditHours: 8,
	}

	t.Run("valid course", func(t *testing.T) {
		v := validator.New()
		ValidateCourse(v, course)
		if !v.IsEmpty() {
			t.Errorf("expected course to be valid, but got errors: %v", v.Errors)
		}
	})

	t.Run("title is empty", func(t *testing.T) {
		c := *course // copy the course
		c.Title = ""
		v := validator.New()
		ValidateCourse(v, &c)
		if v.IsEmpty() {
			t.Error("expected course to be invalid, but it was valid")
		}
		if _, exists := v.Errors["title"]; !exists {
			t.Error("expected error on title field, but it was not found")
		}
	})

	t.Run("category is invalid", func(t *testing.T) {
		c := *course // copy the course
		c.Category = "Optional"
		v := validator.New()
		ValidateCourse(v, &c)
		if v.IsEmpty() {
			t.Error("expected course to be invalid, but it was valid")
		}
		if _, exists := v.Errors["category"]; !exists {
			t.Error("expected error on category field, but it was not found")
		}
	})

	t.Run("credit hours are zero", func(t *testing.T) {
		c := *course // copy the course
		c.CreditHours = 0
		v := validator.New()
		ValidateCourse(v, &c)
		if v.IsEmpty() {
			t.Error("expected course to be invalid, but it was valid")
		}
		if _, exists := v.Errors["credit_hours"]; !exists {
			t.Error("expected error on credit_hours field, but it was not found")
		}
	})
}
