package cos

import "time"

type Config struct {
	// 通用配置
	Region string // 地域/区域
	Bucket string // 存储桶名称

	// 认证配置
	AccessKey string // 访问密钥ID（AWS S3: AccessKeyId, 腾讯云: SecretID, 阿里云: AccessKeyId）
	SecretKey string // 访问密钥Key（AWS S3: SecretAccessKey, 腾讯云: SecretKey, 阿里云: AccessKeySecret）
	Token     string // 临时令牌（可选）

	// 高级配置
	Timeout    time.Duration // 请求超时时间
	RetryCount int           // 重试次数
	EnableSSL  bool          // 是否启用SSL
	PathStyle  bool          // 是否使用路径样式URL（主要用于S3兼容服务）
}
