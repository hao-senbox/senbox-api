package uploader

import (
	"context"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dest string) (*string, error)
	GetFileUploaded(ctx context.Context, key string) (*string, error)
	DeleteFileUploaded(ctx context.Context, key string) error
}
