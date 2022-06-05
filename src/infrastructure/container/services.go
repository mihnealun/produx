package container

import "C"
import (
	"context"
	"github.com/mindstand/gogm/v2"
	"log"
	service2 "produx/domain/category/service"
	storage2 "produx/domain/category/storage"
	"produx/domain/entity"
	service3 "produx/domain/product/service"
	storage3 "produx/domain/product/storage"
	service4 "produx/domain/seller/service"
	storage4 "produx/domain/seller/storage"
	"produx/domain/service"
	"produx/infrastructure/storage"
	"sync"
)

// Container interface that described what services it holds
type Container interface {
	GetConfig() *Config
	GetLogger(ctx context.Context) (Logger, error)
	GetCommentService() service.Comment
	GetUserService() service.User
	GetTargetService() service.Target
	GetAppService() service.App
	GetCategoryService() service2.Category
	GetProductService() service3.Product
	GetSellerService() service4.Seller
}

type container struct {
	config    *Config
	ogmConfig gogm.Config
	gogm      *gogm.Gogm
}

var instance *container
var once sync.Once

// GetInstance return the container as a singleton instance
func GetInstance() (c Container, err error) {
	once.Do(func() {
		instance = &container{}
		instance.config, err = getConfigInstance()
		if err != nil {
			return
		}
		log.Printf("%v+", instance.config)

		instance.ogmConfig = gogm.Config{
			Host:     instance.config.NeoHost,
			Port:     instance.config.NeoPort,
			Username: instance.config.NeoUser,
			LogLevel: instance.config.NeoLogLevel,
			Password: instance.config.NeoPass,
			PoolSize: instance.config.NeoPoolSize,
			// Encrypted:     false,
			IndexStrategy: gogm.IGNORE_INDEX,
		}
		err = instance.InitStorageDriver()
		if err != nil {
			panic(err)
		}
	})

	return instance, err
}

// GetConfig is returning the Config instance
func (c *container) GetConfig() *Config {
	return c.config
}

func (c *container) GetLogger(ctx context.Context) (Logger, error) {
	return newStdLogger(ctx, c.config)
}

func (c *container) InitStorageDriver() error {
	var err error

	c.gogm, err = gogm.New(
		&c.ogmConfig,
		gogm.UUIDPrimaryKeyStrategy,
		&entity.Seller{},
		&entity.Product{},
		&entity.Category{},
		&entity.Comment{},
		&entity.Target{},
		&entity.User{},
		&entity.App{},
	)
	if err != nil {
		panic(err)
	}

	gogm.SetGlobalGogm(c.gogm)

	return nil
}

func (c *container) GetCommentService() service.Comment {
	return storage.NewCommentService(c.gogm)
}

func (c *container) GetUserService() service.User {
	return storage.NewUserService(c.gogm)
}

func (c *container) GetTargetService() service.Target {
	return storage.NewTargetService(c.gogm)
}

func (c *container) GetAppService() service.App {
	return storage.NewAppService(c.gogm)
}

func (c *container) GetCategoryService() service2.Category {
	return storage2.NewCategoryService(c.gogm)
}

func (c *container) GetProductService() service3.Product {
	return storage3.NewProductService(c.gogm)
}

func (c *container) GetSellerService() service4.Seller {
	return storage4.NewSellerService(c.gogm)
}
