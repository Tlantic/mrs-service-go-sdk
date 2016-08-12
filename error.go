package mrs_service_go_sdk

type ApiError struct {
	Status 		int 		`json:"status"`
	Error 		string 		`json:"error"`
	Message 	interface{} 	`json:"message"`
}