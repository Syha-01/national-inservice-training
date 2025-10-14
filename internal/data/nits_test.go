package data

import (
	"testing"
	"time"

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

func TestValidateOfficer(t *testing.T) {
	officer := &Officer{
		RegulationNumber: "12345",
		FirstName:        "John",
		LastName:         "Doe",
		Sex:              "Male",
	}

	t.Run("valid officer", func(t *testing.T) {
		v := validator.New()
		ValidateOfficer(v, officer)
		if !v.IsEmpty() {
			t.Errorf("expected officer to be valid, but got errors: %v", v.Errors)
		}
	})

	t.Run("regulation number is empty", func(t *testing.T) {
		o := *officer
		o.RegulationNumber = ""
		v := validator.New()
		ValidateOfficer(v, &o)
		if v.IsEmpty() {
			t.Error("expected officer to be invalid, but it was valid")
		}
		if _, exists := v.Errors["regulation_number"]; !exists {
			t.Error("expected error on regulation_number field, but it was not found")
		}
	})

	t.Run("sex is invalid", func(t *testing.T) {
		o := *officer
		o.Sex = "Other"
		v := validator.New()
		ValidateOfficer(v, &o)
		if v.IsEmpty() {
			t.Error("expected officer to be invalid, but it was valid")
		}
		if _, exists := v.Errors["sex"]; !exists {
			t.Error("expected error on sex field, but it was not found")
		}
	})
}

func TestValidateNit(t *testing.T) {
	nit := &Nit{
		CourseID:  1,
		StartDate: time.Now(),
		EndDate:   time.Now().Add(24 * time.Hour),
		Location:  "Training Center",
	}

	t.Run("valid nit", func(t *testing.T) {
		v := validator.New()
		ValidateNit(v, nit)
		if !v.IsEmpty() {
			t.Errorf("expected nit to be valid, but got errors: %v", v.Errors)
		}
	})

	t.Run("end date is before start date", func(t *testing.T) {
		n := *nit
		n.StartDate = time.Now().Add(24 * time.Hour)
		n.EndDate = time.Now()
		v := validator.New()
		ValidateNit(v, &n)
		if v.IsEmpty() {
			t.Error("expected nit to be invalid, but it was valid")
		}
		if _, exists := v.Errors["end_date"]; !exists {
			t.Error("expected error on end_date field, but it was not found")
		}
	})

	t.Run("location is empty", func(t *testing.T) {
		n := *nit
		n.Location = ""
		v := validator.New()
		ValidateNit(v, &n)
		if v.IsEmpty() {
			t.Error("expected nit to be invalid, but it was valid")
		}
		if _, exists := v.Errors["location"]; !exists {
			t.Error("expected error on location field, but it was not found")
		}
	})
}