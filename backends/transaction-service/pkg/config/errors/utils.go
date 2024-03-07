package errors

func Is(err error, code *ErrorCode) bool {
	if err == nil {
		return false
	}

	if err, ok := err.(*Error); ok {
		return err.code == code
	}

	return false
}
