package request

type UpdatePasswordRequest struct {
	CurrentPasssword string `json:"current_password" validate:"required"`
	NewPassword      string `json:"new_password" validate:"required,min=8"`
	ConfirmPassword  string `json:"confirm_password" validate:"required,eqfield=NewPassword"`
}
