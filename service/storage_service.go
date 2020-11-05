package service

import (
	"context"
	"github.com/aulang/site/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
)

type StorageService interface {
	Put(bucket, name string, contentType string, reader io.Reader, size int64) error
	Get(bucket, name string) (*minio.Object, error)
	Remove(bucket, name string) error
}

type storageService struct {
	client *minio.Client
	ctx    context.Context
}

func (s storageService) Put(bucket, name string, contentType string, reader io.Reader, size int64) error {
	_, err := s.client.PutObject(s.ctx, bucket, name, reader, size, minio.PutObjectOptions{ContentType: contentType})
	return err
}

func (s storageService) Get(bucket, name string) (*minio.Object, error) {
	return s.client.GetObject(s.ctx, bucket, name, minio.GetObjectOptions{})
}

func (s storageService) Remove(bucket, name string) error {
	return s.client.RemoveObject(s.ctx, bucket, name, minio.RemoveObjectOptions{})
}

var _ StorageService = (*storageService)(nil)

func NewStorageService() StorageService {
	client, err := minio.New(config.Config.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Config.Minio.AccessKey, config.Config.Minio.SecretKey, ""),
		Secure: config.Config.Minio.UseSSL,
	})

	if err != nil {
		log.Fatalf("连接Minio失败: %v", err)
	}

	return storageService{
		client: client,
		ctx:    ctx,
	}
}
