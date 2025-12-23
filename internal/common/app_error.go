package common

type AppError struct {
	StatusCode int
	Message    string
	Reason     string
}

func (e *AppError) Error() string {
	return e.Reason
}
