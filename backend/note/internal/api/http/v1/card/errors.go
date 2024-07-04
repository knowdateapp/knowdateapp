package card

import "fmt"

func NewDefaultErrorResponse(code string, format string, a ...any) DefaultErrorResponse {
	return DefaultErrorResponse{
		Code:    code,
		Message: fmt.Sprintf(format, a...),
	}
}
