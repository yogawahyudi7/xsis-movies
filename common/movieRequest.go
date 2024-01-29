package common

type AddMovieRequest struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Rating      float64 `json:"rating" validate:"required,numeric"`
	Image       string  `json:"image"`
}

type UpdateMovieRequest struct {
	Id          int     `json:"id" validate:"required"`
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Rating      float64 `json:"rating" validate:"required"`
	Image       string  `json:"image"`
}
