package staque

type Empty string
func (empty Empty) Error() string {
	return string(empty)
}