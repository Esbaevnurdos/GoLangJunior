package models

type Review struct {
	ID         int    `json:"id"`
	MovieID    int    `json:"movie_id"`
	UserID     int    `json:"user_id"`
	Rating     int    `json:"rating"`
	ReviewText string `json:"review_text"`
}