package storage

import (
	"context"
	"github.com/mindstand/gogm/v2"
	"log"
	"produx/domain/category/service"
	"produx/domain/entity"
)

type category struct {
	driver *gogm.Gogm
}

func NewCategoryService(driver *gogm.Gogm) service.Category {
	return &category{
		driver: driver,
	}
}

func (a *category) Add(categ entity.Category) *entity.Category {
	sess, err := a.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		log.Fatal(err)
	}

	defer a.commitAndClose(sess)

	err = sess.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = sess.SaveDepth(context.Background(), &categ, 1)
	if err != nil {
		log.Fatal(err)
	}

	var result entity.Category
	err = sess.Load(context.Background(), &result, categ.UUID)
	if err != nil {
		log.Fatal(err)
	}

	return &result
}

func (a *category) Update(categ entity.Category) *entity.Category {
	return &categ
}

func (a *category) Delete(categ entity.Category) bool {
	return true
}

func (a *category) Get(id string) *entity.Category {
	var category entity.Category

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

	err = sess.Load(context.Background(), &category, &id)
	if err != nil {
		log.Println(err.Error())
	}

	return &category
}

func (a *category) List() []*entity.Category {
	var allCategs []*entity.Category

	sess, err := a.driver.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		log.Println(err.Error())
		return allCategs
	}

	err = sess.Begin(context.Background())
	if err != nil {
		log.Println(err.Error())
		return allCategs
	}

	defer a.commitAndClose(sess)

	err = sess.LoadAll(context.Background(), &allCategs)
	if err != nil {
		log.Println(err.Error())
	}

	return allCategs
}

func (a *category) commitAndClose(sess gogm.SessionV2) {
	err := sess.Commit(context.Background())
	if err != nil {
		log.Fatal(sess.RollbackWithError(context.Background(), err))
	}

	_ = sess.Close()
}
