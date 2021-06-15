package setting

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitSetting(t *testing.T) {
	t.Run("InitSetting", func(t *testing.T) {
		InitSetting()
		assert.Equal(t, 30, Setting.Redis.MaxActive)
		fmt.Println(Setting)
	})
}

func TestInitRedisPool(t *testing.T) {
	t.Run("InitRedisPool", func(t *testing.T) {
		InitSetting()
		InitRedisPool()

	})
}

func TestInitDatabase(t *testing.T) {
	t.Run("InitDatabase", func(t *testing.T) {
		InitSetting()
		InitDatabase()
	})
}
