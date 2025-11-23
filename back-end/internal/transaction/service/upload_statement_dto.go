package service

import (
	"io"
)

type UploadRequest struct {
	File io.Reader
}

type ParseError struct {
	Fields map[string]any
}

func (ve *ParseError) Error() string {
	return "validation failed"
}
