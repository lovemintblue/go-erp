package app

import "go-erp-api/pkg/config"

// IsLocal 是否本地环境
func IsLocal() bool {
	return config.Get("app.env") == "local"
}

// IsProduction 是否生产环境
func IsProduction() bool {
	return config.Get("app.env") == "production"
}

// IsTesting 是否测试环境
func IsTesting() bool {
	return config.Get("app.env") == "testing"
}
