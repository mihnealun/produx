package storage

import (
	"context"
	"log"
	"produx/domain/attribute_group/service"
	"produx/domain/entity"

	"github.com/mindstand/gogm/v2"
)

type attributeGroup struct {
	driver *gogm.Gogm
}

func NewAttributeGroupService(driver *gogm.Gogm) service.AttributeGroup {
	return &attributeGroup{
		driver: driver,
	}
}

func (a *attributeGroup) GetByName(name string) *entity.AttributeGroup {
	var atGroup entity.AttributeGroup

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

	query := "match c=(:AttributeGroup{name:$name}) return c"
	err = sess.Query(context.Background(), query, map[string]interface{}{
		"name": name,
	}, &atGroup)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return &atGroup
}

func (a *attributeGroup) Get(id string) *entity.AttributeGroup {
	var attribute entity.AttributeGroup

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

func (a *attributeGroup) Add(attribute entity.AttributeGroup) *entity.AttributeGroup {
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

	var result entity.AttributeGroup
	err = sess.Load(context.Background(), &result, attribute.UUID)
	if err != nil {
		log.Fatal(err)
	}

	return &result
}

func (a *attributeGroup) Update(item entity.AttributeGroup) *entity.AttributeGroup {
	return &item
}

func (a *attributeGroup) Delete(item entity.AttributeGroup) bool {
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

func (a *attributeGroup) List() []*entity.AttributeGroup {
	var items []*entity.AttributeGroup

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

func (a *attributeGroup) commitAndClose(sess gogm.SessionV2) {
	err := sess.Commit(context.Background())
	if err != nil {
		log.Fatal(sess.RollbackWithError(context.Background(), err))
	}

	_ = sess.Close()
}
