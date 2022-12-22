package responses

type FailureMessageResponseModel struct {
	Message string `json:"msg"`
}

func NewFailureResponse(status int, message string) *ResponseModel[FailureMessageResponseModel] {
	return &ResponseModel[FailureMessageResponseModel]{
		Status: status,
		Data: FailureMessageResponseModel{
			Message: message,
		},
	}
}
