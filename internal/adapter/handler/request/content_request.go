package request

type ContentRequest struct {
	Title       string   `json:"title" validate:"required,min=3,max=255"`
	Excerpt     string   `json:"excerpt" validate:"required,min=10,max=255"`
	Description string   `json:"description,omitempty" validate:"omitempty,min=10,max=1000"`
	Image       string   `json:"image" validate:"required"`
	Tags        []string `json:"tags"`
	CategoryID  int64    `json:"category_id" validate:"required"`
	Status      string   `json:"status" validate:"required"`
}
