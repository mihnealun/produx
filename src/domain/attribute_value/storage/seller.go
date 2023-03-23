package storage

import (
	"context"
	"log"
	"produx/domain/attribute_value/service"
	"produx/domain/entity"

	"github.com/mindstand/gogm/v2"
)

type attributeValue struct {
	driver *gogm.Gogm
}

func NewAttributeValueService(driver *gogm.Gogm) service.AttributeValue {
	return &attributeValue{
		driver: driver,
	}
}

func (a *attributeValue) Get(id string) *entity.AttributeValue {
	var attributeValue entity.AttributeValue

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

	err = sess.Load(context.Background(), &attributeValue, &id)
	if err != nil {
		log.Println(err.Error())
	}

	return &attributeValue
}

func (a *attributeValue) Add(attributeValue entity.AttributeValue) *entity.AttributeValue {
	sess, err := a.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		log.Fatal(err)
	}

	defer a.commitAndClose(sess)

	err = sess.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = sess.Save(context.Background(), &attributeValue)
	if err != nil {
		log.Fatal(err)
	}

	var result entity.AttributeValue
	err = sess.Load(context.Background(), &result, attributeValue.UUID)
	if err != nil {
		log.Fatal(err)
	}

	return &result
}

func (a *attributeValue) Update(item entity.AttributeValue) *entity.AttributeValue {
	return &item
}

func (a *attributeValue) Delete(item entity.AttributeValue) bool {
	sess, err := a.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		log.Println(err.Error())
		return false
	}

	err = sess.Begin(context.Background())
	if err != nil {
		log.Println(err.Error())
		return false
	}

	defer a.commitAndClose(sess)

	err = sess.Delete(context.Background(), item)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	return true
}

func (a *attributeValue) List() []*entity.AttributeValue {
	var allProds []*entity.AttributeValue

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

func (a *attributeValue) commitAndClose(sess gogm.SessionV2) {
	err := sess.Commit(context.Background())
	if err != nil {
		log.Fatal(sess.RollbackWithError(context.Background(), err))
	}

	_ = sess.Close()
}
