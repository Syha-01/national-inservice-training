package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Syha-01/national-inservice-training/internal/data"
)

func (a *application) createNitHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		CourseID  int64     `json:"course_id"`
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		Location  string    `json:"location"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		a.errorResponseJSON(w, r, http.StatusBadRequest, err.Error())
		return
	}

	nit := &data.Nit{
		CourseID:  input.CourseID,
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
		Location:  input.Location,
		CreatedAt: time.Now(),
		Version:   1,
	}

	fmt.Fprintf(w, "%+v\n", nit)
}
