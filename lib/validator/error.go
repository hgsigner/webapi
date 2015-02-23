package validator

type ErrorMessage struct {
	Field   string
	Message string
}

type ErrorsList []ErrorMessage

func (e *ErrorsList) Any() bool {
	if len(*e) > 0 {
		return true
	}
	return false
}

func (e *ErrorsList) AppendError(err ErrorMessage) ErrorsList {
	*e = append(*e, err)
	return *e
}
