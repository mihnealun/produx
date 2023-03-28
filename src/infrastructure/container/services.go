package container

import "C"
import (
	"context"
	"fmt"
	"log"
	"os"
	attributeService "produx/domain/attribute/service"
	attributeStorage "produx/domain/attribute/storage"
	attributeGroupService "produx/domain/attribute_group/service"
	attributeGroupStorage "produx/domain/attribute_group/storage"
	categoryService "produx/domain/category/service"
	categoryStorage "produx/domain/category/storage"
	"produx/domain/entity"
	productService "produx/domain/product/service"
	productStorage "produx/domain/product/storage"
	sellerService "produx/domain/seller/service"
	sellerStorage "produx/domain/seller/storage"
	"produx/domain/service"
	"produx/infrastructure/storage"
	"sync"

	"github.com/mindstand/gogm/v2"
)

// Container interface that described what services it holds
type Container interface {
	GetConfig() *Config
	GetLogger(ctx context.Context) (Logger, error)
	GetCommentService() service.Comment
	GetUserService() service.User
	GetTargetService() service.Target
	GetAppService() service.App
	GetCategoryService() categoryService.Category
	GetProductService() productService.Product
	GetSellerService() sellerService.Seller
	GetAttributeService() attributeService.Attribute
	GetAttributeGroupService() attributeGroupService.AttributeGroup
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
		&entity.Attribute{},
		&entity.AttributeGroup{},
		&entity.AttributeValue{},
		&entity.Comment{},
		&entity.Target{},
		&entity.User{},
		&entity.App{},
	)
	if err != nil {
		log.Println(fmt.Errorf("cannot connect to Neo4J server: %w", err))
		os.Exit(0)
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

func (c *container) GetCategoryService() categoryService.Category {
	return categoryStorage.NewCategoryService(c.gogm)
}

func (c *container) GetProductService() productService.Product {
	return productStorage.NewProductService(c.gogm)
}

func (c *container) GetSellerService() sellerService.Seller {
	return sellerStorage.NewSellerService(c.gogm)
}

func (c *container) GetAttributeService() attributeService.Attribute {
	return attributeStorage.NewAttributeService(c.gogm)
}

func (c *container) GetAttributeGroupService() attributeGroupService.AttributeGroup {
	return attributeGroupStorage.NewAttributeGroupService(c.gogm)
}
