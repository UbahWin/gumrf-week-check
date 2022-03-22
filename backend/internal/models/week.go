package models

type WhatWeekRequestDTO struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type WhatWeekResponseDTO struct {
	WeekFraction string `json:"result"`
}
