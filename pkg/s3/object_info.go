package s3

import "time"

type FileInfo struct {
	name        string
	contentType string
	size        int64
	s3Key       string
	s3Bucket    string
	createdAt   time.Time
}

func NewFileInfo(
	name string,
	contentType string,
	size int64,
	createdAt time.Time,
	s3Key string,
	s3Bucket string,
) *FileInfo {
	return &FileInfo{
		name:        name,
		contentType: contentType,
		size:        size,
		createdAt:   createdAt,
		s3Key:       s3Key,
		s3Bucket:    s3Bucket,
	}
}

func (f *FileInfo) Name() string {
	return f.name
}

func (f *FileInfo) Size() int64 {
	return f.size
}

func (f *FileInfo) ContentType() string {
	return f.contentType
}

func (f *FileInfo) S3Key() string {
	return f.s3Key
}

func (f *FileInfo) S3Bucket() string {
	return f.s3Bucket
}

func (f *FileInfo) CreatedAt() time.Time {
	return f.createdAt
}
