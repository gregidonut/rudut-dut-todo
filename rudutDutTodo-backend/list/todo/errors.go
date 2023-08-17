package todo

import "errors"

var (
	MoreThanOneStateErr  = errors.New("there is more than one progress state")
	TodoInstantiationErr = errors.New("having trouble instantiating todo struct")
)
