package erroring

type ErrorCode string

const (
	// Пример для Swagger, должен быть первым
	ErrorCodeExample ErrorCode = "error_code_example"

	RequestValidationCode ErrorCode = "request_validation"
	InternalServerCode    ErrorCode = "internal_error"
)

type HTTPError[T any] struct {
	Message string    `json:"message"`
	Code    ErrorCode `json:"code"`

	Data T `json:"data,omitempty"`
}

type ValidationErrorFieldDTO struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type HTTPRequestValidationError HTTPError[[]ValidationErrorFieldDTO]

type HTTPInternalServerError HTTPError[any]
