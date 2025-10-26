package handler

import "github.com/pkg/errors"

var (
	errBadRequest     = errors.New("bad request")
	errInternalServer = errors.New("internal server error")
)
