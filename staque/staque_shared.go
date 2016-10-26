package staque

type empty struct {
	s string
}
func (e *empty) Error() string {
	return e.s
}