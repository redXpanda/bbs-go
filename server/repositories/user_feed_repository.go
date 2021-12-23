package repositories

import (
	"bbs-go/model"
	"github.com/mlogclub/simple"
	"gorm.io/gorm"
)

var UserFeedRepository = newUserFeedRepository()

func newUserFeedRepository() *userFeedRepository {
	return &userFeedRepository{}
}

type userFeedRepository struct {
}

func (r *userFeedRepository) Get(db *gorm.DB, id int64) *model.UserFeed {
	ret := &model.UserFeed{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userFeedRepository) Take(db *gorm.DB, where ...interface{}) *model.UserFeed {
	ret := &model.UserFeed{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userFeedRepository) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.UserFeed) {
	cnd.Find(db, &list)
	return
}

func (r *userFeedRepository) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.UserFeed {
	ret := &model.UserFeed{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *userFeedRepository) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.UserFeed, paging *simple.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *userFeedRepository) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.UserFeed, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.UserFeed{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *userFeedRepository) Count(db *gorm.DB, cnd *simple.SqlCnd) int64 {
	return cnd.Count(db, &model.UserFeed{})
}

func (r *userFeedRepository) Create(db *gorm.DB, t *model.UserFeed) (err error) {
	err = db.Create(t).Error
	return
}

func (r *userFeedRepository) Update(db *gorm.DB, t *model.UserFeed) (err error) {
	err = db.Save(t).Error
	return
}

func (r *userFeedRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.UserFeed{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *userFeedRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.UserFeed{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *userFeedRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.UserFeed{}, "id = ?", id)
}
