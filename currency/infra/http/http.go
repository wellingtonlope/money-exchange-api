package http

type ResponseError struct {
	Message string `json:"message"`
}

func NewResponseError(err error) ResponseError {
	if err != nil {
		return ResponseError{err.Error()}
	}
	return ResponseError{}
}
