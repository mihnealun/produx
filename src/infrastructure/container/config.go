package container

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config contains all the variables from the .route file
type Config struct {
	Env   string `required:"true" envconfig:"env"`
	Debug bool   `required:"true" envconfig:"debug"`

	LogLevel string `required:"true" envconfig:"log_level"`
	LogFile  string `required:"true" envconfig:"log_file"`

	Interface string `required:"true" envconfig:"interface"`
	Port      int    `required:"true" envconfig:"port"`

	NeoHost          string `required:"true" envconfig:"neo4j_host"`
	NeoPort          int    `required:"true" envconfig:"neo4j_port"`
	NeoUser          string `required:"true" envconfig:"neo4j_user"`
	NeoPass          string `required:"true" envconfig:"neo4j_password"`
	NeoLogLevel      string `required:"true" envconfig:"neo4j_log_level"`
	NeoPoolSize      int    `required:"true" envconfig:"neo4j_pool_size"`
	NeoIndexStrategy string `required:"true" envconfig:"neo4j_index_strategy"`
}

var instanceConfig *Config
var onceConfig sync.Once

// GetConfigInstance method reads the .env file by default, validate the fields
func getConfigInstance() (*Config, error) {
	var err error
	onceConfig.Do(func() {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		fmt.Println(exPath)

		err = godotenv.Load(path.Join(exPath, ".env"))
		if err != nil {
			return
		}

		instanceConfig = &Config{}

		err = envconfig.Process("", instanceConfig)
		if err != nil {
			return
		}
	})

	if err != nil {
		return nil, err
	}

	return instanceConfig, nil
}
