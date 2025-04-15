package types

type DbError struct {
	Message string
	InternalError error
}

type ValidationError struct {
	Message string
}

func (d *DbError) Error() string {
	return d.Message
}

func (v *ValidationError) Error() string {
	return v.Message
}