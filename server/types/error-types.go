package types

type ServerError struct {
	Message string
	InternalError error
}

type ValidationError struct {
	Message string
}

func (s *ServerError) Error() string {
	return s.Message
}

func (v *ValidationError) Error() string {
	return v.Message
}