package helper

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseBuilder(status int, msg string, data interface{}) Response {
	return Response{Status: status, Message: msg, Data: data}
}
