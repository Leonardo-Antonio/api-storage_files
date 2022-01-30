package response

type response struct {
	Message     string      `json:"message"`
	MessageType string      `json:"message_type"`
	Data        interface{} `json:"data"`
}
