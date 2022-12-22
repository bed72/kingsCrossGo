package request

type SignUpRequestModel struct {
	Name     string `json:"name" validate:"required,lte=16"`
	Password string `json:"password" validate:"required,lte=64"`
	Email    string `json:"email" validate:"required,email,lte=128"`
}
