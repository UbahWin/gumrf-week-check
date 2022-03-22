package week

import (
	"encoding/json"
	"fmt"
	"github.com/ubahwin/week-of-learn/internal/models"
	"net/http"
	"time"
)

func WhatWeek(w http.ResponseWriter, r *http.Request) {
	var dto models.WhatWeekRequestDTO

	json.NewDecoder(r.Body).Decode(&dto)

	currentDate, err := time.Parse("2/1/2006", fmt.Sprintf("%d/%d/%d", dto.Day, dto.Month, dto.Year))
	if err != nil {
		panic(err)
	}

	var startYear int

	if dto.Month < 9 {
		startYear = dto.Year - 1
	} else {
		startYear = dto.Year
	}

	startDate, err := time.Parse("2/1/2006", fmt.Sprintf("%d/%d/%d", 1, 9, startYear))

	var days int = int(currentDate.Sub(startDate).Hours()) / 24

	var daysLeft int = days - (int)(7-startDate.Weekday())

	var isEven bool = false
	for i := 0; daysLeft > 0; i += 1 {
		daysLeft -= 7
		isEven = !isEven
	}

	var result models.WhatWeekResponseDTO

	if isEven {
		result.WeekFraction = "denominator"
	} else {
		result.WeekFraction = "numerator"
	}

	json.NewEncoder(w).Encode(result)
}
