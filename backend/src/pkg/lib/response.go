package lib

type StatusCode int

const (
	StatusCodeOK                  StatusCode = 200
	StatusCodeBadRequest          StatusCode = 400
	StatusCodeInternalServerError StatusCode = 500
)

type ErrorResponseBody struct {
	Message string         `json:"message"`
	Errors  []ErrorContent `json:"errors,omitempty"`
}

type ErrorContentType string

const (
	ErrorTypeBadRequest ErrorContentType = "bad request"
	ErrorTypeNotFound   ErrorContentType = "not found"
	ErrorInvalidValue   ErrorContentType = "invalid value"
)

type ErrorContent struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewValidateErrorResponseBody(errors []ErrorContent) ErrorResponseBody {
	return ErrorResponseBody{
		Message: "invalid request",
		Errors:  errors,
	}
}
