package exceptions

type InternalServerError struct {
	Detail string `json:"detail,omitempty"`
}

func (e *InternalServerError) Error() string {
	return e.Detail
}

type ConflictError struct {
	Detail string `json:"detail,omitempty"`
}

func (e *ConflictError) Error() string {
	return e.Detail
}

type NotFoundError struct {
	Detail string `json:"detail,omitempty"`
}

func (e *NotFoundError) Error() string {
	return e.Detail
}
