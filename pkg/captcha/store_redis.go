package captcha

import (
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/redis"
	"time"
)

// RedisStore 实现 base64Captcha.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

// Set 实现 base64Captcha.Store interface 的 Set 方法
func (s *RedisStore) Set(id string, value string) {

	ExpireTime := time.Minute * time.Duration(config.GetInt64("captcha.expire_time"))
	// 方便本地开发调试
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("captcha.debug_expire_time"))
	}

	s.RedisClient.Set(s.KeyPrefix+id, value, ExpireTime)
}

// Get 实现 base64Captcha.Store interface 的 Get 方法
func (s *RedisStore) Get(key string, clear bool) string {
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val
}

// Verify 实现 base64Captcha.Store interface 的 Verify 方法
func (s *RedisStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
