package responses

type SignUpResponseModel struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewSignUpResponse(status int, data SignUpResponseModel) *ResponseModel[SignUpResponseModel] {
	return &ResponseModel[SignUpResponseModel]{
		Status: status,
		Data:   data,
	}
}
