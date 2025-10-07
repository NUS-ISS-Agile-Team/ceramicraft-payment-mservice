package biz_error

// BizError represents a business error with a code, message, and an optional underlying error.
type BizError struct {
	Code    int
	Message string
	Err     error
}

func (e *BizError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}
