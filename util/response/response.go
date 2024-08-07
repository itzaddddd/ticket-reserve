package response

type ErrorData struct {
	Error string `json:"error"`
}

func ErrorResponse(err error) *ErrorData {
	return &ErrorData{
		Error: err.Error(),
	}
}
