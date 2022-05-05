package entity

type (
	CustomError struct {
		Err      error
		HTTPCode int
		Code     string
	}
	ResponseError struct {
		Error string `json:"error"`
		Code  string `json:"code"`
	}
)

func (ce CustomError) Error() string {
	return ce.Err.Error()
}

func (ce CustomError) ToResponseError() ResponseError {
	return ResponseError{
		Error: ce.Error(),
		Code:  ce.Code,
	}
}
