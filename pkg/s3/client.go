package s3

import (
	"context"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type With struct {
	ctx    context.Context
	bucket string
	region string
}

type Client struct {
	minio *minio.Client
	with  With
}

func NewClient(endpoint string, accessKey, secretKey string, secure bool) (*Client, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: secure,
	})

	if err != nil {
		return nil, err
	}

	return &Client{
		minio: client,
	}, nil
}

func (c *Client) WithContext(ctx context.Context) *Client {
	c.with.ctx = ctx
	return c
}

func (c *Client) WithBucket(bucket string) *Client {
	c.with.bucket = bucket
	return c
}

func (c *Client) WithRegion(region string) *Client {
	c.with.region = region
	return c
}

func (c *Client) upload(name string, filePath string, metadata map[string]string) (*FileInfo, error) {

	contentType, err := c.getMimeType(filePath)

	if err != nil {
		return nil, err
	}

	metadata["Content-Type"] = contentType

	objectInfo, err := c.minio.FPutObject(c.with.ctx, c.with.bucket, name, filePath, minio.PutObjectOptions{UserMetadata: metadata})

	if err != nil {
		return nil, err
	}

	fileInfo := NewFileInfo(name, contentType, objectInfo.Size, objectInfo.LastModified, objectInfo.Key, objectInfo.Bucket)
	return fileInfo, nil
}

func (c *Client) Download(name string, filePath string) error {
	return c.minio.FGetObject(c.with.ctx, c.with.bucket, name, filePath, minio.GetObjectOptions{})
}

func (c *Client) getMimeType(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Читаем первые 512 байт
	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	// MIME-тип по содержимому
	contentMIME := http.DetectContentType(buffer[:n])

	// MIME-тип по расширению
	ext := filepath.Ext(filePath)
	extensionMIME := mime.TypeByExtension(ext)

	// Выбираем наиболее подходящий MIME-тип
	finalMIME := contentMIME
	if contentMIME == "application/octet-stream" && extensionMIME != "" {
		finalMIME = extensionMIME
	}

	return finalMIME, nil
}
