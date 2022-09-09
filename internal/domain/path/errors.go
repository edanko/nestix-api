package path

import "errors"

var (
	ErrValidationLength   = errors.New("length must be greater than zero")
	ErrValidationHeight   = errors.New("height must be greater than zero")
	ErrValidationMetafile = errors.New("metafile cannot be empty")
	ErrValidationFilename = errors.New("filename cannot be empty")
)
