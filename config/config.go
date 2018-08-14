package config

import (
	"log"
	"strings"
	"sync"

	"fmt"

	"github.com/spf13/viper"
)

var (
	// 單例模式實例
	instance *viper.Viper
	once     sync.Once
)

// MockConfig : 由外部注入Config實例，供測試使用，一般正常流程不應呼叫
func MockConfig(v *viper.Viper) {
	instance = v
}

// Forge : 取得實例
func Forge() *viper.Viper {
	once.Do(func() {
		instance = New()
	})

	return instance
}

// New : 產生Config物件
func New() *viper.Viper {
	instance = viper.New()
	instance.SetConfigName("config")
	instance.SetConfigType("json")
	instance.AddConfigPath(".")
	instance.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	instance.AutomaticEnv()

	if err := instance.ReadInConfig(); err != nil {
		log.Println(
			fmt.Errorf("No Default configuration file loaded, because %v", err.Error()),
		)
	}

	return instance
}
