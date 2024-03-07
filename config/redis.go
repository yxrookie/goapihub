package config

import (
	"goapihub/pkg/config"
)

func init() {
	config.Add("redis", func() map[string]interface{}{
		return map[string]interface{} {
			"host": config.Env("REDIS_HOST", "127.0.0.1"),
			"port": config.Env("REDIS_PORT", "6379"),
			"password": config.Env("REDIS_PASSWORD",""),
			// picture verification code, message verification code, and session use the database 1
			"database": config.Env("REDIS_MIN_DB", 1),
		}
	})
}