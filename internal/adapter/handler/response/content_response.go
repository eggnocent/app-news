package response

type ContentResponse struct {
	ID           int64    `json:"id"`
	Title        string   `json:"title"`
	Excerpt      string   `json:"excerpt"`
	Description  string   `json:"description,omitempty"`
	Image        string   `json:"image"`
	Tags         []string `json:"tags"`
	Status       string   `json:"status"`
	CategoryID   int64    `json:"category_id,omitempty"`
	CreatedByID  int64    `json:"created_by_id,omitempty"`
	CreatedAt    string   `json:"created_at"`
	CategoryName string   `json:"category_name"`
	Author       string   `json:"author"`
}
