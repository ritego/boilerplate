package res

type Response struct {
	Status  string      `json:"status"`
	Title   string      `json:"title"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func New(data interface{}, message string, title string, status string) Response {
	return Response{
		Status:  status,
		Title:   title,
		Message: message,
		Data:    data,
	}
}

func Success(data interface{}, message string, title string) Response {
	if title == "" {
		title = "Success"
	}

	if message == "" {
		message = "Request was successful"
	}

	return New(data, message, title, SuccessResponseStatus)
}

func Error(message string, title string, data interface{}) Response {
	if title == "" {
		title = "Error"
	}

	if data == nil {
		data = nil
	}

	return New(data, message, title, ErrorResponseStatus)
}

func S(data interface{}) Response {
	return Success(data, "", "")
}

func E(message string) Response {
	return Error(message, "", nil)
}

func FromError(e error) Response {
	return Error(e.Error(), "", nil)
}
