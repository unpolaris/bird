package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type BirdModel interface {
	Create(
		birdName string,
		birdType int64,
		description string,
		picUrl string) error
	Update(
		birdId int64,
		birdName string,
		birdType int64,
		description string,
		picUrl string) error
	List(
		page int,
		pageSize int) ([]*Bird, error)
	Get(
		birdId int64,
	) (*Bird, error)
	Search(
		birdName string,
	) ([]*Bird, error)
	Delete(
		birdId int64,
	) error
}

type Bird struct {
	Id          int64     `gorm:"column:id" json:"id"`
	BirdName    string    `gorm:"column:birdName" json:"birdName"`
	BirdType    int64     `gorm:"column:birdType" json:"birdType"`
	Description string    `gorm:"column:description" json:"description"`
	PicUrl      string    `gorm:"column:picUrl" json:"picUrl"`
	CreateTime  time.Time `gorm:"column:createTime" json:"createTime"`
	UpdateTime  time.Time `gorm:"column:updateTime" json:"updateTime"`
}

func (t *Bird) TableName() string {
	return "bird"
}

type defaultBirdModel struct {
	db  *gorm.DB
	ctx context.Context
}

func NewBirdModel(db *gorm.DB, ctx context.Context) BirdModel {
	return &defaultBirdModel{
		db:  db,
		ctx: ctx,
	}
}

func (m *defaultBirdModel) TableName() string {
	return "bird"
}

func (m *defaultBirdModel) Create(
	birdName string,
	birdType int64,
	description string,
	picUrl string) error {
	b := map[string]any{
		"birdName":    birdName,
		"birdType":    birdType,
		"description": description,
		"picUrl":      picUrl,
	}
	return m.db.Table(m.TableName()).Create(b).Error
}

func (m *defaultBirdModel) Update(
	birdId int64,
	birdName string,
	birdType int64,
	description string,
	picUrl string) error {
	updates := map[string]any{
		"birdName":    birdName,
		"birdType":    birdType,
		"description": description,
		"picUrl":      picUrl,
	}
	return m.db.Table(m.TableName()).Where("id = ?", birdId).Updates(updates).Error
}

func (m *defaultBirdModel) List(
	page int,
	pageSize int) ([]*Bird, error) {
	var birds []*Bird
	err := m.db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&birds).Error
	return birds, err
}

func (m *defaultBirdModel) Get(
	birdId int64,
) (*Bird, error) {
	var b Bird
	err := m.db.Where("id = ?", birdId).Find(&b).Error
	return &b, err
}

func (m *defaultBirdModel) Search(
	birdName string) ([]*Bird, error) {
	var birds []*Bird
	err := m.db.Where("birdName Like ? ", "%"+birdName+"%").Limit(10).Find(&birds).Error
	return birds, err
}

func (m *defaultBirdModel) Delete(
	birdId int64,
) error {
	err := m.db.Where("id = ?", birdId).Delete(&Bird{}).Error
	return err
}
