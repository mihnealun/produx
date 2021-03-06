package storage

import (
	"context"
	"github.com/mindstand/gogm/v2"
	"log"
	"produx/domain/entity"
	"produx/domain/product/service"
)

type product struct {
	driver *gogm.Gogm
}

func NewProductService(driver *gogm.Gogm) service.Product {
	return &product{
		driver: driver,
	}
}

func (a *product) Get(id string) *entity.Product {
	sess, err := a.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	err = sess.Begin(context.Background())
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	defer a.commitAndClose(sess)

	var result entity.Product

	err = sess.Load(context.Background(), &result, id)
	if err != nil {
		log.Println(err.Error())
	}

	return &result
}

func (a *product) Add(prod entity.Product) *entity.Product {
	sess, err := a.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		log.Fatal(err)
	}

	defer a.commitAndClose(sess)

	err = sess.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = sess.SaveDepth(context.Background(), &prod, 1)
	if err != nil {
		log.Fatal(err)
	}

	var result entity.Product
	err = sess.Load(context.Background(), &result, prod.UUID)
	if err != nil {
		log.Fatal(err)
	}

	return &result
}

func (a *product) Update(prod entity.Product) *entity.Product {
	return &prod
}

func (a *product) Delete(prod entity.Product) bool {
	return true
}

func (a *product) List() []*entity.Product {
	var allProds []*entity.Product

	sess, err := a.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		log.Println(err.Error())
		return allProds
	}

	err = sess.Begin(context.Background())
	if err != nil {
		log.Println(err.Error())
		return allProds
	}

	defer a.commitAndClose(sess)

	err = sess.LoadAll(context.Background(), &allProds)
	if err != nil {
		log.Println(err.Error())
	}

	return allProds
}

func (a *product) commitAndClose(sess gogm.SessionV2) {
	err := sess.Commit(context.Background())
	if err != nil {
		log.Fatal(sess.RollbackWithError(context.Background(), err))
	}

	_ = sess.Close()
}
