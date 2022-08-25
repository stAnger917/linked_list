package linkedListImpl

import "errors"

var ErrNoItem = errors.New("list is empty")
var ErrNilHead = errors.New("head element is nil")
var ErrNilTail = errors.New("tail element is nil")
