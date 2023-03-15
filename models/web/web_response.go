package model_web

const (
	SuccessStatus             = "SUCCESS"
	BadRequestStatus          = "BAD_REQUEST"
	UnauthorizedStatus        = "UNAUTHORIZED"
	ForbiddenStatus           = "FORBIDDEN"
	NotFoundStatus            = "NOT_FOUND"
	InternalServerErrorStatus = "INTERNAL_SERVER_ERROR"
)

type WebResponse[T any] struct {
	Status  string `json:"status"`
	Data    T      `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

var NotFound = &WebResponse[any]{
	Status: NotFoundStatus,
}

var InternalServerError = &WebResponse[any]{
	Status: InternalServerErrorStatus,
}
