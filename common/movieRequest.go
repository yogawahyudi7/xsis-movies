package common

type UpsertMovieRequest struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Rating      float64 `json:"rating" validate:"required,numeric"`
	Image       string  `json:"image"`
}
