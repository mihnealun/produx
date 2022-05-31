package storage

import (
	"context"
	"github.com/mindstand/gogm/v2"
	"log"
	"produx/domain/entity"
	"produx/domain/seller/service"
)

type seller struct {
	driver *gogm.Gogm
}

func NewSellerService(driver *gogm.Gogm) service.Seller {
	return &seller{
		driver: driver,
	}
}

func (a *seller) Get(id string) *entity.Seller {
	var seller entity.Seller

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

	err = sess.Load(context.Background(), &seller, &id)
	if err != nil {
		log.Println(err.Error())
	}

	return &seller
}

func (a *seller) Add(prod entity.Seller) *entity.Seller {
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

	var result entity.Seller
	err = sess.Load(context.Background(), &result, prod.UUID)
	if err != nil {
		log.Fatal(err)
	}

	return &result
}

func (a *seller) Update(prod entity.Seller) *entity.Seller {
	return &prod
}

func (a *seller) Delete(prod entity.Seller) bool {
	return true
}

func (a *seller) List() []*entity.Seller {
	var allProds []*entity.Seller

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

func (a *seller) commitAndClose(sess gogm.SessionV2) {
	err := sess.Commit(context.Background())
	if err != nil {
		log.Fatal(sess.RollbackWithError(context.Background(), err))
	}

	_ = sess.Close()
}
