package request

type CategoryReqest struct {
	Title string `json:"title" validate:"required,min=3,max=100"`
}
