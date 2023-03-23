package storage

import (
	"context"
	"log"
	"produx/domain/attribute/service"
	"produx/domain/entity"

	"github.com/mindstand/gogm/v2"
)

type attribute struct {
	driver *gogm.Gogm
}

func NewAttributeService(driver *gogm.Gogm) service.Attribute {
	return &attribute{
		driver: driver,
	}
}

func (a *attribute) Get(id string) *entity.Attribute {
	var attribute entity.Attribute

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

	err = sess.Load(context.Background(), &attribute, &id)
	if err != nil {
		log.Println(err.Error())
	}

	return &attribute
}

func (a *attribute) Add(attribute entity.Attribute) *entity.Attribute {
	sess, err := a.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		log.Fatal(err)
	}

	defer a.commitAndClose(sess)

	err = sess.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = sess.Save(context.Background(), &attribute)
	if err != nil {
		log.Fatal(err)
	}

	var result entity.Attribute
	err = sess.Load(context.Background(), &result, attribute.UUID)
	if err != nil {
		log.Fatal(err)
	}

	return &result
}

func (a *attribute) Update(item entity.Attribute) *entity.Attribute {
	return &item
}

func (a *attribute) Delete(item entity.Attribute) bool {
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

func (a *attribute) List() []*entity.Attribute {
	var items []*entity.Attribute

	sess, err := a.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		log.Println(err.Error())
		return items
	}

	err = sess.Begin(context.Background())
	if err != nil {
		log.Println(err.Error())
		return items
	}

	defer a.commitAndClose(sess)

	err = sess.LoadAll(context.Background(), &items)
	if err != nil {
		log.Println(err.Error())
	}

	return items
}

func (a *attribute) commitAndClose(sess gogm.SessionV2) {
	err := sess.Commit(context.Background())
	if err != nil {
		log.Fatal(sess.RollbackWithError(context.Background(), err))
	}

	_ = sess.Close()
}
