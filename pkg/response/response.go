package response

func Success(message string, data interface{}) *response {
	return &response{
		Message:     message,
		MessageType: "SUCCESS",
		Data:        data,
	}
}

func Err(message string, data interface{}) *response {
	return &response{
		Message:     message,
		MessageType: "ERROR",
		Data:        data,
	}
}
