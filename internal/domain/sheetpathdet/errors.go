package sheetpathdet

import "errors"

var (
	ErrValidationSequenceNo       = errors.New("sequence number lower than 1")
	ErrValidationDetailType       = errors.New("unknown detail type")
	ErrValidationPartOrderID      = errors.New("part order id not found")
	ErrValidationPartPartID       = errors.New("part part id not found")
	ErrValidationRemnantProductID = errors.New("remnant product id not found")
)
