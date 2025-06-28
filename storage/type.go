package storage

import "context"

// ObjectStorage 对象存储服务的抽象
// 最重要的功能: 上传 + 获取
type ObjectStorage interface {
	// Upload 上传数据到对象存储
	// bucket: 存储桶名称
	// key: 存储对象的键
	// data: 要上传的数据
	// 返回错误信息
	Upload(ctx context.Context, bucket, key string, data []byte) error
	// Get 获取对象存储中的数据
	Get(ctx context.Context, bucket, key string) ([]byte, error)
	// Delete 删除对象存储中的数据
	Delete(ctx context.Context, bucket, key string) error
}
