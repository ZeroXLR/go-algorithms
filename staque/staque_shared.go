package staque

import "errors"

var (
	emptypeek = errors.New("Cannot Peek() on empty staque")
	emptypop = errors.New("Cannot Pop() on empty staque")
)