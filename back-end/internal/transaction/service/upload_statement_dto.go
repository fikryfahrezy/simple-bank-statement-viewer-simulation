package service

import (
	"io"
)

type UploadRequest struct {
	File io.Reader
}
