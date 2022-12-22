package responses

type ResponseModel[T any] struct {
	Status int `json:"code"`
	Data   T   `json:"data"`
}

// func NewFailureResponse(status int, message string) *FailureResponseModel {
// 	return &FailureResponseModel{
// 		Status: status,
// 		Data: FailureMessageResponseModel{
// 			Message: message,
// 		},
// 	}
// }
