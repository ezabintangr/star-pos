package response

type MapHelloResponse struct {
	Data string `json:"data,omitempty"`
}

func WebJSONHelloResponse(data string) MapHelloResponse {
	return MapHelloResponse{
		Data: data,
	}
}

type MapResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func WebJSONResponse(msg string, data interface{}) MapResponse {
	return MapResponse{
		Message: msg,
		Data:    data,
	}
}
