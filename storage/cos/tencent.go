package cos

import (
	"bytes"
	"context"
	"errors"
	"github.com/A-Drop-Water/xutil/storage"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
)

// cosStorage 腾讯的cos storage
// 他这个就是基于bucket的
type cosStorage struct {
	clients map[string]*cos.Client // 存储多个bucket的客户端
	config  *Config                // 存储配置
}

func (c *cosStorage) setNewClient(bucket string) error {
	// 1. 生成baseURL
	bucketURL, err := cos.NewBucketURL(bucket, c.config.Region, c.config.EnableSSL)
	if err != nil {
		return err
	}
	baseURL := &cos.BaseURL{BucketURL: bucketURL}
	// 2. 创建HTTP客户端
	httpClient := &http.Client{
		Timeout: c.config.Timeout,
		Transport: &cos.AuthorizationTransport{
			SecretID:     c.config.AccessKey,
			SecretKey:    c.config.SecretKey,
			SessionToken: c.config.Token,
		},
	}
	// 3. 创建腾讯云COS客户端
	c.clients[bucket] = cos.NewClient(baseURL, httpClient)
	return nil
}

func NewCosStorage(config *Config) storage.ObjectStorage {
	sto := &cosStorage{
		config:  config,
		clients: map[string]*cos.Client{},
	}
	if config.Bucket != "" {
		if err := sto.setNewClient(config.Bucket); err != nil {
		}
	}
	// 4. 返回对象存储实例
	return sto
}

func (c *cosStorage) getBucketClient(bucket string) (*cos.Client, error) {
	if client, ok := c.clients[bucket]; ok {
		return client, nil
	}
	if err := c.setNewClient(bucket); err != nil {
		return nil, errors.New("没有对应的存储桶")
	}
	return c.clients[bucket], nil

}

func (c *cosStorage) Upload(ctx context.Context, bucket, key string, data []byte) error {
	client, err := c.getBucketClient(bucket)
	if err != nil {
		return err
	}
	_, err = client.Object.Put(ctx, key, bytes.NewReader(data), nil)
	return err
}

func (c *cosStorage) Get(ctx context.Context, bucket, key string) ([]byte, error) {
	client, err := c.getBucketClient(bucket)
	if err != nil {
		return nil, err
	}
	resp, err := client.Object.Get(ctx, key, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func (c *cosStorage) Delete(ctx context.Context, bucket, key string) error {
	client, err := c.getBucketClient(bucket)
	if err != nil {
		return err
	}
	_, err = client.Object.Delete(ctx, key, nil)
	return err
}
